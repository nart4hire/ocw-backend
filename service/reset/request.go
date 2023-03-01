package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/request"
)

func (rs ResetServiceImpl) Request(payload request.RequestRequestPayload) error {
	c := cache.NewHash(*cache.NewKey("Test", "123"), *cache.NewValue("Test", "123"), 30)
	c.AppendValue(*cache.NewValue("Hello", "World"))

	err := rs.CacheRepository.HSet(*c)
	if err != nil {
		panic(err)
	}

	_, err = rs.CacheRepository.HGet(*c, "Test")
	if err != nil {
		panic(err)
	}
	
	return nil
}