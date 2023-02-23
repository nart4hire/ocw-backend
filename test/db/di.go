package db

import (
	"github.com/google/wire"
	"gitlab.informatika.org/ocw/ocw-backend/provider/db"
)

var DbTestSet = wire.NewSet(
	New,
	wire.Bind(new(db.Database), new(*MockDatabase)),
)
