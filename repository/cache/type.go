package cache

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
)

type CacheRepository interface {
	Get(key cache.Key) (string, error)
	Set(str cache.String) error
	Delete(key string) error
	HGet(cache cache.Hash, field string) (string, error)
	HGetAll(cache cache.Hash) (map[string]string, error)
	HSet(cache cache.Hash) error
}
