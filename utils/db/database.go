package db

import (
	"database/sql"

	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseImpl struct {
	DB *gorm.DB
}

func NewPostgresConn(
	conn *sql.Conn,
) (*DatabaseImpl, error) {
	res, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DatabaseImpl{res}, nil
}

func NewPostgresEnv(
	env env.Environment,
) (*DatabaseImpl, error) {
	res, err := gorm.Open(postgres.Open(env.DatabaseConnection), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DatabaseImpl{res}, nil
}

func (p DatabaseImpl) Connect() *gorm.DB {
	return p.DB
}
