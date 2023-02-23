package test

import (
	"gitlab.informatika.org/ocw/ocw-backend/app"
	"gitlab.informatika.org/ocw/ocw-backend/test/db"
)

type ApiTestPack struct {
	*db.MockDatabase
	app.Server
}
