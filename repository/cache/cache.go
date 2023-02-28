package cache

import (
	"github.com/gomodule/redigo/redis"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	rd "gitlab.informatika.org/ocw/ocw-backend/provider/redis"
)

type CacheRepositoryImpl struct {
	pool *redis.Pool
}

func New(
	cache rd.Redis,
) *CacheRepositoryImpl {
	return &CacheRepositoryImpl{cache.Pool()}
}

func (c CacheRepositoryImpl) Get(cache cache.Cache, field string) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("HGET", cache.Key.String(), field))

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c CacheRepositoryImpl) GetAll(cache cache.Cache) (map[string]string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	value, err := redis.StringMap(conn.Do("HGETALL", cache.Key.String()))
	
	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c CacheRepositoryImpl) Set(cache cache.Cache) error {
	conn := c.pool.Get()
	defer conn.Close()

	slice := cache.Slice()
	_, err := conn.Do("HSET", slice...)

	if err != nil {
		return err
	}

	return nil
}