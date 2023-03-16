package repository

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
	"gitlab.informatika.org/ocw/ocw-backend/repository/course"
	"gitlab.informatika.org/ocw/ocw-backend/repository/cache"
	"gitlab.informatika.org/ocw/ocw-backend/repository/content"
	"gitlab.informatika.org/ocw/ocw-backend/repository/material"
	"gitlab.informatika.org/ocw/ocw-backend/repository/transaction"
)

var RepositoryBasicSet = wire.NewSet(
	// User Repository
	user.New,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),

	// Course Repository
	course.New,
	wire.Bind(new(course.CourseRepository), new(*course.CourseRepositoryImpl)),

	// Cache Repository
	cache.New,
	wire.Bind(new(cache.CacheRepository), new(*cache.CacheRepositoryImpl)),

	material.NewMaterial,
	material.NewMaterialContent,

	wire.Struct(new(content.ContentRepositoryImpl), "*"),

	wire.Bind(new(material.MaterialContentRepository), new(*material.MaterialContentRepositoryImpl)),
	wire.Bind(new(material.MaterialRepository), new(*material.MaterialRepositoryImpl)),

	wire.Bind(new(content.ContentRepository), new(*content.ContentRepositoryImpl)),

	transaction.New,
	transaction.NewBuilder,
	wire.Bind(new(transaction.Transaction), new(*transaction.TransactionRepositoryImpl)),
	wire.Bind(new(transaction.TransactionBuilder), new(*transaction.TransactionBuilderImpl)),
)

var RepositorySet = wire.NewSet(
	RepositoryBasicSet,
)
