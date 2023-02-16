package db

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type MockDatabase struct {
	db   *gorm.DB
	Mock sqlmock.Sqlmock
}

func New() (*MockDatabase, error) {
	db, mock, err := sqlmock.New()

	if err != nil {
		return nil, err
	}

	g, err := gorm.Open(postgres.New(
		postgres.Config{
			Conn: db,
		},
	), &gorm.Config{})

	return &MockDatabase{g, mock}, err
}

func (m MockDatabase) Connect() *gorm.DB {
	return m.db
}
