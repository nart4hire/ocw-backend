package repository

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/cache"
)

var RepositoryBasicSet = wire.NewSet(
	// User Repository
	user.New,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),

	// Cache Repository
	cache.New,
	wire.Bind(new(cache.CacheRepository), new(*cache.CacheRepositoryImpl)),
)

var RepositorySet = wire.NewSet(
	RepositoryBasicSet,
)
