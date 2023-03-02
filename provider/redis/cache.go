package redis

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gomodule/redigo/redis"
	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
)

type RedisImpl struct {
	pool *redis.Pool
}

func resolver(log logger.Logger) {
	if rec := recover(); rec != nil {
		log.Error("Some panic occured when processing request:")
		log.Error(fmt.Sprint(rec))
		log.Error("")

		log.Error("Stack Trace:")
		stacks := strings.Split(string(debug.Stack()), "\n")

		for _, val := range stacks {
			log.Error(val)
		}

		os.Exit(-1)
	}
}

func NewRedisEnv(
	env *env.Environment,
	log logger.Logger,
) (*RedisImpl, error) {
	return &RedisImpl{
		&redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				defer resolver(log)
				conn, err := redis.Dial("tcp", env.RedisConnection+":"+env.RedisPort)

				if err != nil {
					log.Warning("failed connect to redis server: tcp " + env.RedisConnection + ":" + env.RedisPort)
					log.Warning(err.Error())

					return nil, err
				}

				if env.RedisUseAuth {
					if _, err := conn.Do("AUTH", env.RedisUsername, env.RedisPassword); err != nil {
						conn.Close()

						log.Warning("failed connect to redis server: authentication failed")

						return nil, err
					}
				}

				return conn, err
			},
			TestOnBorrow: func(c redis.Conn, t time.Time) error {
				if time.Since(t) < time.Minute {
					return nil
				}
				_, err := c.Do("PING")

				if err != nil {
					log.Warning("redis server is down")
				}

				return err
			},
		}}, nil
}

func (r RedisImpl) Pool() *redis.Pool {
	return r.pool
}
