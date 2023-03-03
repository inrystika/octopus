package redis

import (
	"context"
	"net/url"
	"server/base-server/internal/conf"
	"server/common/log"
	commredis "server/common/redis"
	"time"
)

type Redis interface {
	LockAndCall(ctx context.Context, key string, ttl time.Duration, f func() error) error
	SMembersMinioRemovingObject() ([]string, error)
	SAddMinioRemovingObject(object string) error
	SRemMinioRemovingObject(object string) error
	Close()
}

type redis struct {
	log      *log.Helper
	instance *commredis.RedisInstance
}

func NewRedis(conf *conf.Data, logger log.Logger) (Redis, error) {
	redisUrl := url.URL{
		Scheme: "redis",
		Host:   conf.Redis.Addr,
		Path:   "0",
		User:   url.UserPassword(conf.Redis.Username, conf.Redis.Password),
	}
	rdb, err := commredis.GetRedisInstance(redisUrl.String())
	if err != nil {
		return nil, err
	}

	return &redis{
		log:      log.NewHelper("Redis", logger),
		instance: rdb,
	}, nil
}

func (r *redis) LockAndCall(ctx context.Context, key string, ttl time.Duration, f func() error) error {
	return r.instance.LockAndCall(ctx, key, ttl, f)
}

func (r *redis) SMembersMinioRemovingObject() ([]string, error) {
	return r.instance.SMembersMinioRemovingObject()
}

func (r *redis) SAddMinioRemovingObject(object string) error {
	return r.instance.SAddMinioRemovingObject(object)
}

func (r *redis) SRemMinioRemovingObject(object string) error {
	return r.instance.SRemMinioRemovingObject(object)
}

func (r *redis) Close() {
	ctx := context.Background()
	err := r.instance.Redis.Close()
	if err != nil {
		r.log.Warn(ctx, err)
	}
	r.log.Info(ctx, "close redis")
}
