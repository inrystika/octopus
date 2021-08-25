package session

import (
	"context"
	"server/common/constant"
	"server/common/log"
	ss "server/common/session"
	"server/openai-server/internal/conf"
)

type SessionClient interface {
	ss.SessionStore
}

func NewSessionClient(config *conf.Data, logger log.Logger) SessionClient {
	logHelper := log.NewHelper("Session", logger)

	storeConfig := ss.SessionStoreConfig{
		RedisAddr:     config.Redis.Addr,
		RedisPassword: config.Redis.Password,
		RedisUsername: config.Redis.Username,
		RedisDBIndex:  "0",
	}
	return &SessionClientImpl{
		config: config,
		store:  ss.NewSessionStore(constant.SESSION_KEY, storeConfig, logger),
		logger: logHelper,
	}
}

type SessionClientImpl struct {
	config *conf.Data
	store  ss.SessionStore
	logger *log.Helper
}

func (s *SessionClientImpl) Create(ctx context.Context, session *ss.Session) error {
	return s.store.Create(ctx, session)
}

func (s *SessionClientImpl) Get(ctx context.Context, sessionId string) (*ss.Session, error) {
	return s.store.Get(ctx, sessionId)
}

func (s *SessionClientImpl) Update(ctx context.Context, session *ss.Session) error {
	return s.store.Update(ctx, session)
}

func (s *SessionClientImpl) Delete(ctx context.Context, sessionId string) error {
	return s.store.Delete(ctx, sessionId)
}
