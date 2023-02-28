package cache

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
)

type CacheRepository interface {
	Get(cache cache.Cache, field string) (string, error)
	GetAll(cache cache.Cache) (map[string]string, error)
	Set(cache cache.Cache) error
}
