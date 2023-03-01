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

func (c CacheRepositoryImpl) Get(key cache.Key) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("GET", key))

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c CacheRepositoryImpl) Set(str cache.String) error {
	conn := c.pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", str.Key, str.Value)

	if err != nil {
		return err
	}

	if str.ExpiryInMinutes > 0 {
		_, err = conn.Do("EXPIRE", str.Key, str.ExpiryInMinutes * 60)

		if err != nil {
			return err
		}
	}

	return nil
}

func (c CacheRepositoryImpl) HGet(cache cache.Hash, field string) (string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	value, err := redis.String(conn.Do("HGET", cache.Key.String(), field))

	if err != nil {
		return "", err
	}

	return value, nil
}

func (c CacheRepositoryImpl) HGetAll(cache cache.Hash) (map[string]string, error) {
	conn := c.pool.Get()
	defer conn.Close()

	value, err := redis.StringMap(conn.Do("HGETALL", cache.Key.String()))

	if err != nil {
		return nil, err
	}

	return value, nil
}

func (c CacheRepositoryImpl) HSet(cache cache.Hash) error {
	conn := c.pool.Get()
	defer conn.Close()

	slice := cache.Slice()
	_, err := conn.Do("HSET", slice...)

	if err != nil {
		return err
	}

	if cache.ExpiryInMinutes > 0 {
		_, err = conn.Do("EXPIRE", cache.Key, cache.ExpiryInMinutes * 60)

		if err != nil {
			return err
		}
	}

	return nil
}
