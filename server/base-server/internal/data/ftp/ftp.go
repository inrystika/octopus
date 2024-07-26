package ftp

import (
	"context"
	"io/ioutil"
	"net/http"
	"server/base-server/internal/conf"
	"server/common/errors"
	"server/common/log"
	sftpgov2 "server/common/sftpgo/v2/openapi"
	"sync"
	"time"
)

type ftpAuthToken struct {
	mu sync.Mutex

	username string
	password string
	token    *sftpgov2.Token
	client   *sftpgov2.APIClient
}

type Ftp struct {
	client *sftpgov2.APIClient
	token  *ftpAuthToken
}

func NewFtp(conf *conf.Bootstrap) *Ftp {
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

	return &Ftp{
		token:  &token,
		client: client,
	}
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

func (s *Ftp) getAuthCtx(ctx context.Context) (context.Context, error) {
	ctx = context.WithValue(ctx, sftpgov2.ContextServerVariables, map[string]string{
		"basePath": "v2",
	})
	token, err := s.token.getToken(ctx)
	if err != nil {
		return nil, err
	}
	return context.WithValue(ctx, sftpgov2.ContextAccessToken, token), nil
}

func (s *Ftp) GetFtpUser(ctx context.Context, username string) (*sftpgov2.User, error) {
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

func (s *Ftp) printResponse(ctx context.Context, r *http.Response) {
	if r != nil {
		bytes, _ := ioutil.ReadAll(r.Body)
		log.Infof(ctx, "resp: %v", string(bytes))
	}
}

func (s *Ftp) CreateFtpUser(ctx context.Context, fuser *sftpgov2.User) (*sftpgov2.User, error) {
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

func (s *Ftp) UpdateFtpUser(ctx context.Context, user sftpgov2.User, disconnect int32) error {
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
