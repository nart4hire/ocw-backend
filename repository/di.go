package repository

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/repository/cache"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
)

var RepositoryBasicSet = wire.NewSet(
	// User Repository
	user.New,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),

	// Cache Repository
	cache.New,
	wire.Bind(new(cache.CacheRepository), new(*cache.CacheRepositoryImpl)),

	material.NewMaterial,
	material.NewMaterialContent,

	wire.Bind(new(material.MaterialContentRepository), new(*material.MaterialContentRepositoryImpl)),
	wire.Bind(new(material.MaterialRepository), new(*material.MaterialRepositoryImpl)),
)

var RepositorySet = wire.NewSet(
	RepositoryBasicSet,
)
