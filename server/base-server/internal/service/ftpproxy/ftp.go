package ftpproxy

import (
	"context"
	"io/ioutil"
	"net/http"
	"server/common/errors"
	"sync"
	"time"

	pb "server/base-server/api/v1"
	"server/base-server/internal/conf"
	"server/base-server/internal/data"
	"server/common/log"
	sftpgov2 "server/common/sftpgo/v2/openapi"
	commUtils "server/common/utils"
)

const (
	UNLIMITED = 0
)

type FtpProxyService struct {
	pb.UnimplementedFtpProxyServiceServer
	conf   *conf.Bootstrap
	data   *data.Data
	client *sftpgov2.APIClient
	token  *ftpAuthToken
}

type ftpAuthToken struct {
	mu sync.Mutex

	username string
	password string
	token    *sftpgov2.Token
	client   *sftpgov2.APIClient
}

func (t *ftpAuthToken) getToken(ctx context.Context) (string, error) {
	if !t.isExpired() {
		return *t.token.AccessToken, nil
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.isExpired() {
		return *t.token.AccessToken, nil
	}
	tk, resp, err := t.client.TokenApi.GetToken(context.WithValue(ctx, sftpgov2.ContextBasicAuth, sftpgov2.BasicAuth{
		UserName: t.username,
		Password: t.password,
	})).Execute()
	if err != nil {
		log.Errorf(ctx, "FtpGO GetToken failed, err: %v", err)
		return "", errors.Errorf(err, errors.ErrorSFtpGOAPIRequestFailed)
	}
	if resp.StatusCode == http.StatusOK {
		t.token.AccessToken = tk.AccessToken
		t.token.ExpiresAt = tk.ExpiresAt
		return *t.token.AccessToken, nil
	}
	log.Errorf(ctx, "FtpGO GetToken failed, statusCode: %v", resp.StatusCode)
	return "", errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
}

func (t *ftpAuthToken) isExpired() bool {
	if t.token.AccessToken == nil || *t.token.AccessToken == "" {
		return true
	}
	// 30 seconds in advance
	return !time.Now().Add(30 * time.Second).Before(*t.token.ExpiresAt)
}

func NewFtpProxyService(conf *conf.Bootstrap, data *data.Data) pb.FtpProxyServiceServer {
	config := sftpgov2.NewConfiguration()
	config.Scheme = "http"
	config.Host = conf.Data.Sftpgo.BaseUrl

	client := sftpgov2.NewAPIClient(config)
	token := ftpAuthToken{
		client:   client,
		token:    &sftpgov2.Token{},
		username: conf.Data.Sftpgo.Username,
		password: conf.Data.Sftpgo.Password,
	}
	ctx := context.WithValue(context.Background(), sftpgov2.ContextServerVariables, map[string]string{
		"basePath": "v2",
	})
	token.getToken(ctx)

	return &FtpProxyService{
		token:  &token,
		client: client,
		conf:   conf,
		data:   data,
	}
}

func (s *FtpProxyService) CreateOrUpdateFtpAccount(ctx context.Context, req *pb.CreateOrUpdateFtpAccountRequest) (*pb.CreateOrUpdateFtpAccountReply, error) {
	var err error
	password := ""
	if req.Password != "" {
		password, err = commUtils.EncryptPassword(req.Password)
		if err != nil {
			return nil, err
		}
	}

	fuser, err := s.getFtpUser(ctx, req.Username)
	if err != nil && !errors.IsError(errors.ErrorSFtpGOUserNotExist, err) {
		return nil, err
	}

	if fuser == nil {
		fuser = sftpgov2.NewUser()
		//去掉minio中转
		//fileSystemConfig := s.newFileSystemConfig(req.HomeS3Bucket, req.HomeS3Object)
		permissions := map[string][]sftpgov2.Permission{
			"/": {sftpgov2.PERMISSION_STAR},
		}

		fuser.SetStatus(1)
		fuser.SetUid(0)
		fuser.SetGid(0)
		fuser.SetMaxSessions(UNLIMITED)
		fuser.SetQuotaSize(UNLIMITED)
		fuser.SetQuotaFiles(UNLIMITED)
		fuser.SetExpirationDate(UNLIMITED)
		fuser.SetPermissions(permissions)
		//fuser.SetFilesystem(*fileSystemConfig)
		fuser.SetHomeDir(req.HomeDir)
		fuser.SetUploadBandwidth(UNLIMITED)
		fuser.SetDownloadBandwidth(UNLIMITED)

		fuser.SetUsername(req.Username)
		fuser.SetEmail(req.Email)
		fuser.SetPassword(password)

		_, err = s.createFtpUser(ctx, fuser)
		if err != nil {
			return nil, err
		}
	} else {
		if req.Username != "" {
			fuser.SetUsername(req.Username)
		}
		if req.Email != "" {
			fuser.SetEmail(req.Email)
		}
		if req.Password != "" {
			fuser.SetPassword(password)
		}
		if req.HomeDir != "" {
			fileSystemConfig := sftpgov2.NewFilesystemConfig()
			fileSystemConfig.SetProvider(sftpgov2.FSPROVIDERS__0)
			fuser.SetFilesystem(*fileSystemConfig)
			fuser.SetHomeDir(req.HomeDir)
		}
		err := s.updateFtpUser(ctx, *fuser, 1)
		if err != nil {
			return nil, err
		}
	}

	return &pb.CreateOrUpdateFtpAccountReply{}, nil
}

func (s *FtpProxyService) getAuthCtx(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, sftpgov2.ContextServerVariables, map[string]string{
		"basePath": "v2",
	})
	token, err := s.token.getToken(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, sftpgov2.ContextAccessToken, token), nil
}

func (s *FtpProxyService) getFtpUser(ctx context.Context, username string) (*sftpgov2.User, error) {
	ctx, err := s.getAuthCtx(ctx)
	if err != nil {
		return nil, err
	}
	user, resp, err := s.client.UsersApi.GetUserByUsername(ctx, username).Execute()
	httpStatusCode := resp.StatusCode
	if err != nil && httpStatusCode != http.StatusNotFound {
		return nil, err
	}

	switch httpStatusCode {
	case http.StatusOK:
		return user, nil
	case http.StatusNotFound:
		return nil, errors.Errorf(err, errors.ErrorSFtpGOUserNotExist)
	default:
		log.Errorf(ctx, "FtpGO GetUserByUsername, username: %s, statusCode: %v", username, httpStatusCode)
		return nil, errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}

func (s *FtpProxyService) createFtpUser(ctx context.Context, fuser *sftpgov2.User) (*sftpgov2.User, error) {
	ctx, err := s.getAuthCtx(ctx)
	if err != nil {
		return nil, err
	}
	u, resp, err := s.client.UsersApi.AddUser(ctx).User(*fuser).Execute()
	if err != nil {
		s.printResponse(ctx, resp)
		return nil, errors.Errorf(err, errors.ErrorSFtpGOAPIRequestFailed)
	}
	httpStatusCode := resp.StatusCode
	switch httpStatusCode {
	case http.StatusCreated:
		return u, nil
	default:
		log.Errorf(ctx, "FtpGO createFtpUser, username: %v, statusCode: %v", fuser.Username, httpStatusCode)
		return nil, errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}

func (s *FtpProxyService) newFileSystemConfig(bucket string, object string) *sftpgov2.FilesystemConfig {
	accessSecret := sftpgov2.NewSecret()
	accessSecret.SetPayload(s.conf.Data.Minio.Base.SecretAccessKey)
	accessSecret.SetStatus("Plain")

	s3EndPoint := "://" + s.conf.Data.Minio.Base.EndPoint
	if s.conf.Data.Minio.Base.UseSSL {
		s3EndPoint = "https" + s3EndPoint
	} else {
		s3EndPoint = "http" + s3EndPoint
	}

	s3Config := sftpgov2.NewS3Config()
	s3Config.SetBucket(bucket)
	s3Config.SetRegion("us-east-1")
	s3Config.SetAccessKey(s.conf.Data.Minio.Base.AccessKeyID)
	s3Config.SetAccessSecret(*accessSecret)
	s3Config.SetEndpoint(s3EndPoint)
	s3Config.SetForcePathStyle(true)
	s3Config.SetKeyPrefix(object)

	fileSystemConfig := sftpgov2.NewFilesystemConfig()
	fileSystemConfig.SetProvider(sftpgov2.FSPROVIDERS__1)
	fileSystemConfig.SetS3config(*s3Config)
	return fileSystemConfig
}

func (s *FtpProxyService) updateFtpUser(ctx context.Context, user sftpgov2.User, disconnect int32) error {
	for i := 0; i < len(user.VirtualFolders); i++ {
		user.VirtualFolders[i].Filesystem = nil
	}
	ctx, err := s.getAuthCtx(ctx)
	if err != nil {
		return err
	}
	_, resp, err := s.client.UsersApi.UpdateUser(ctx, *user.Username).User(user).Disconnect(disconnect).Execute()
	if err != nil {
		s.printResponse(ctx, resp)
		return errors.Errorf(err, errors.ErrorSFtpGOAPIRequestFailed)
	}
	httpStatusCode := resp.StatusCode
	switch httpStatusCode {
	case http.StatusOK:
		return nil
	default:
		log.Errorf(ctx, "FtpGO updateFtpUser, username: %v, statusCode: %v", user.Username, httpStatusCode)
		return errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}

func (s *FtpProxyService) CreateVirtualFolder(ctx context.Context, req *pb.CreateVirtualFolderRequest) (*pb.CreateVirtualFolderReply, error) {
	folder := sftpgov2.NewBaseVirtualFolder()
	folder.SetName(req.Name)
	folder.SetFilesystem(*s.newFileSystemConfig(req.S3Bucket, req.S3Object))
	err := s.addVirtualFolder(ctx, *folder)
	if err != nil {
		return nil, err
	}

	user, err := s.getFtpUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	vf := sftpgov2.NewVirtualFolder(req.VirtualPath)
	vf.SetName(req.Name)
	vf.SetVirtualPath(req.VirtualPath)
	user.VirtualFolders = append(user.VirtualFolders, *vf)

	err = s.updateFtpUser(ctx, *user, 1)
	if err != nil {
		return nil, err
	}

	return &pb.CreateVirtualFolderReply{}, nil
}

func (s *FtpProxyService) addVirtualFolder(ctx context.Context, bvf sftpgov2.BaseVirtualFolder) error {
	ctx, err := s.getAuthCtx(ctx)
	if err != nil {
		return err
	}
	_, resp, err := s.client.FoldersApi.AddFolder(ctx).BaseVirtualFolder(bvf).Execute()
	if err != nil {
		s.printResponse(ctx, resp)
		return errors.Errorf(err, errors.ErrorSFtpGOAPIRequestFailed)
	}

	httpStatusCode := resp.StatusCode
	switch httpStatusCode {
	case http.StatusCreated:
		return nil
	default:
		return errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}

func (s *FtpProxyService) printResponse(ctx context.Context, r *http.Response) {
	if r != nil {
		bytes, _ := ioutil.ReadAll(r.Body)
		log.Infof(ctx, "resp: %v", string(bytes))
	}
}

func (s *FtpProxyService) DeleteVirtualFolder(ctx context.Context, req *pb.DeleteVirtualFolderRequest) (*pb.DeleteVirtualFolderReply, error) {
	user, err := s.getFtpUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	hasVf := false
	for _, v := range user.VirtualFolders {
		if *v.Name == req.Name {
			hasVf = true
		}
	}

	if !hasVf {
		return nil, errors.Errorf(nil, errors.ErrorSFtpGOUserNotOwnVirtualDir)
	}

	err = s.deleteVirtualFolder(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	user, err = s.getFtpUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	err = s.updateFtpUser(ctx, *user, 1)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteVirtualFolderReply{}, nil
}

func (s *FtpProxyService) deleteVirtualFolder(ctx context.Context, name string) error {
	ctx, err := s.getAuthCtx(ctx)
	if err != nil {
		return err
	}
	_, resp, err := s.client.FoldersApi.DeleteFolder(ctx, name).Execute()
	if err != nil {
		return err
	}
	httpStatusCode := resp.StatusCode
	switch httpStatusCode {
	case http.StatusOK:
		return nil
	default:
		return errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}
