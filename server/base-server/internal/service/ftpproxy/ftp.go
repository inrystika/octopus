package ftpproxy

import (
	"context"
	"net/http"
	"server/base-server/internal/common"
	"server/base-server/internal/data/dao/model"
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
	conf         *conf.Bootstrap
	log          *log.Helper
	data         *data.Data
	client       *sftpgov2.APIClient
	token        *ftpAuthToken
}

type ftpAuthToken struct {
	mu          sync.Mutex

	username    string
	password    string
	token       *sftpgov2.Token
	client      *sftpgov2.APIClient
	log         *log.Helper
}

func (t *ftpAuthToken) getToken(ctx context.Context) string {
	if !t.isExpired() {
		return *t.token.AccessToken
	}

	t.mu.Lock()
	defer t.mu.Unlock()
	if !t.isExpired() {
		return *t.token.AccessToken
	}
	apiGetTokenRequest := t.client.TokenApi.GetToken(context.WithValue(ctx, sftpgov2.ContextBasicAuth, sftpgov2.BasicAuth{
		UserName: t.username,
		Password: t.password,
	}))
	tk, resp , err := apiGetTokenRequest.Execute()
	if err != nil {
		t.log.Errorf(ctx, "FtpGO GetToken failed, err: %v", err)
		panic(err)
	}
	if resp.StatusCode == http.StatusOK {
		t.token.AccessToken = tk.AccessToken
		t.token.ExpiresAt   = tk.ExpiresAt
		return *t.token.AccessToken
	}
	t.log.Errorf(ctx, "FtpGO GetToken failed, statusCode: %v", resp.StatusCode)
	return ""
}

func (t *ftpAuthToken) isExpired() bool {
	if t.token.AccessToken == nil || *t.token.AccessToken == "" {
		return true
	}
	// 30 seconds in advance
	return !time.Now().Add(30 * time.Second).Before(*t.token.ExpiresAt)
}

func NewFtpProxyService(conf *conf.Bootstrap, logger log.Logger, data *data.Data) pb.FtpProxyServiceServer {
	config := sftpgov2.NewConfiguration()
	config.Scheme = "http"
	config.Host = conf.Data.Sftpgo.BaseUrl

	client := sftpgov2.NewAPIClient(config)
	token := ftpAuthToken{
		client:    client,
		token:     &sftpgov2.Token{},
		username:  conf.Data.Sftpgo.Username,
		password:  conf.Data.Sftpgo.Password,
		log:       log.NewHelper("FtpProxyService.ftpAccessToken", logger),
	}
	ctx := context.WithValue(context.Background(), sftpgov2.ContextServerVariables, map[string]string{
		"basePath": "v2",
	})
	token.getToken(ctx)

	return &FtpProxyService{
		token:        &token,
		client:       client,
		conf:         conf,
		data:         data,
		log:          log.NewHelper("FtpProxyService", logger),
	}
}

func (s *FtpProxyService) CreateOrUpdateAccount(ctx context.Context, req *pb.CreateOrUpdateAccountRequest) (*pb.CreateOrUpdateAccountReply, error) {
	fuser, err := s.getFtpUser(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	password, err := commUtils.EncryptPassword(req.Password)
	if err != nil {
		return nil, err
	}
	if fuser == nil {
		user, err := s.data.UserDao.Find(ctx, &model.UserQuery{Id: req.Id})
		if err != nil {
			return nil, err
		}
		if user == nil {
			return nil, errors.Errorf(nil, errors.ErrorUserAccountNotExisted)
		}

		fuser = sftpgov2.NewUser()
		fuser.SetUsername(req.Username)
		fuser.SetEmail(req.Email)
		fuser.SetPassword(password)
		fuser.SetHomeDir(req.HomeDir)

		_, err = s.createFtpUser(ctx, fuser, user)
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
			fuser.SetHomeDir(req.HomeDir)
		}
		err := s.updateFtpUser(ctx, *fuser, 1)
		if err != nil {
			return nil, err
		}
	}

	return &pb.CreateOrUpdateAccountReply{
	}, nil
}

func (s *FtpProxyService) getAuthCtx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, sftpgov2.ContextServerVariables, map[string]string{
		"basePath": "v2",
	})
	return context.WithValue(ctx, sftpgov2.ContextAccessToken, s.token.getToken(ctx))
}

func (s *FtpProxyService) getFtpUser (ctx context.Context, username string) (*sftpgov2.User, error){
	user, resp, err := s.client.UsersApi.GetUserByUsername(s.getAuthCtx(ctx), username).Execute()
	httpStatusCode := resp.StatusCode
	if err != nil && httpStatusCode != http.StatusNotFound {
		return nil, err
	}

	switch httpStatusCode {
	case http.StatusOK:
		return user, nil
	case http.StatusNotFound:
		return nil, nil
	default:
		s.log.Errorf(ctx, "FtpGO GetUserByUsername, username: %s, statusCode: %v", username, httpStatusCode)
		return nil, errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}

func (s *FtpProxyService) createFtpUser (ctx context.Context, fuser *sftpgov2.User, user *model.User) (*sftpgov2.User, error){
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
	s3Config.SetBucket(common.GetUserHomeBucket(user.Id))
	s3Config.SetRegion("us-east-1")
	s3Config.SetAccessKey(s.conf.Data.Minio.Base.AccessKeyID)
	s3Config.SetAccessSecret(*accessSecret)
	s3Config.SetEndpoint(s3EndPoint)
	s3Config.SetForcePathStyle(true)

	fileSystemConfig := sftpgov2.NewFilesystemConfig()
	fileSystemConfig.SetProvider(sftpgov2.FSPROVIDERS__1)
	fileSystemConfig.SetS3config(*s3Config)

	permissions := map[string][]sftpgov2.Permission{}
	permissions["/"] = []sftpgov2.Permission{sftpgov2.PERMISSION_STAR}

	fuser.SetStatus(1)
	fuser.SetUid(0)
	fuser.SetGid(0)
	fuser.SetMaxSessions(UNLIMITED)
	fuser.SetQuotaSize(UNLIMITED)
	fuser.SetQuotaFiles(UNLIMITED)
	fuser.SetExpirationDate(UNLIMITED)
	fuser.SetPermissions(permissions)
	fuser.SetFilesystem(*fileSystemConfig)
	fuser.SetUploadBandwidth(UNLIMITED)
	fuser.SetDownloadBandwidth(UNLIMITED)

	u, resp, err := s.client.UsersApi.AddUser(s.getAuthCtx(ctx)).User(*fuser).Execute()
	if err != nil {
		return nil, err
	}
	httpStatusCode := resp.StatusCode
	switch httpStatusCode {
	case http.StatusCreated:
		return u, nil
	default:
		s.log.Errorf(ctx, "FtpGO createFtpUser, username: %s, statusCode: %v", fuser.Username, httpStatusCode)
		return nil, errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}

func (s *FtpProxyService) updateFtpUser (ctx context.Context, user sftpgov2.User, disconnect int32) error {
	_, resp, err := s.client.UsersApi.UpdateUser(s.getAuthCtx(ctx),*user.Username).User(user).Disconnect(disconnect).Execute()
	if err != nil {
		return err
	}
	httpStatusCode := resp.StatusCode
	switch httpStatusCode {
	case http.StatusOK:
		return nil
	default:
		s.log.Errorf(ctx, "FtpGO updateFtpUser, username: %s, statusCode: %v", user.Username, httpStatusCode)
		return errors.Errorf(nil, errors.ErrorSFtpGOAPIRequestFailed)
	}
}