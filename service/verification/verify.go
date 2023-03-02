package verification

import (
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/cache"
	"gitlab.informatika.org/ocw/ocw-backend/model/domain/user"
	"gitlab.informatika.org/ocw/ocw-backend/model/web"
)

func (v VerificationServiceImpl) DoVerification(id string) error {
	// TODO
	email, err := v.CacheRepository.Get(cache.Key{
		Id: v.RedisPrefixKey + "verify:id:" + id,
	})

	if err != nil {
		return err
	}

	if email == "" {
		return web.NewResponseErrorf("VERIFY", "id '%s' is not valid", id)
	}

	v.UserRepository.Update(user.User{
		Email:       email,
		IsActivated: true,
	})

	return nil
}
