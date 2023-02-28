package reset

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/model/web/reset/request"
)

func (rs ResetServiceImpl) Request(payload request.RequestRequestPayload) error {
	c := cache.NewCache(*cache.NewKey("Test", "123"), *cache.NewValue("Test", "123"), 30)
	c.AppendValue(*cache.NewValue("Hello", "World"))

	err := rs.CacheRepository.Set(*c)
	if err != nil {
		panic(err)
	}

	_, err = rs.CacheRepository.Get(*c, "Test")
	if err != nil {
		panic(err)
	}
	
	return nil
}