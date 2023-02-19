package test

import (
	"gitlab.informatika.org/ocw/ocw-backend/test/db"
	"gitlab.informatika.org/ocw/ocw-backend/utils/app"
)

type ApiTestPack struct {
	*db.MockDatabase
	app.Server
}
