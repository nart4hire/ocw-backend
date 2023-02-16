package repository

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/repository/user"
)

var RepositoryBasicSet = wire.NewSet(
	// User Repository
	user.New,
	wire.Bind(new(user.UserRepository), new(*user.UserRepositoryImpl)),
)

var RepositorySet = wire.NewSet(
	RepositoryBasicSet,
)
