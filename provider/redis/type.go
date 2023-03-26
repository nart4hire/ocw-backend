package redis

import "github.com/gomodule/redigo/redis"

type Redis interface {
	Pool() *redis.Pool
}