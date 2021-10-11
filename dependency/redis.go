package dependency

import (
	"time"
	"webapi/conf"

	"github.com/gomodule/redigo/redis"
)

func InitRedis(cfg *conf.Conf) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     5,
		MaxActive:   100,
		IdleTimeout: time.Duration(30 * time.Second),
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(cfg.Redis)
		},
	}
}
