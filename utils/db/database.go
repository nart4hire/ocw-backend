package db

import (
	"database/sql"
	"fmt"
	"os"
	"runtime/debug"
	"strings"

	"gitlab.informatika.org/ocw/ocw-backend/service/logger"
	"gitlab.informatika.org/ocw/ocw-backend/utils/env"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseImpl struct {
	DB *gorm.DB
}

func resolver(log logger.Logger) {
	if rec := recover(); rec != nil {
		log.Error("Some panic occured when processing request:")
		log.Error(fmt.Sprint(rec))
		log.Error("")

		log.Error("Stack Trace:")
		stacks := strings.Split(string(debug.Stack()), "\n")

		for _, val := range stacks {
			log.Error(val)
		}

		os.Exit(-1)
	}
}

func NewPostgresConn(
	conn *sql.Conn,
	log logger.Logger,
) (*DatabaseImpl, error) {
	defer resolver(log)

	res, err := gorm.Open(postgres.New(postgres.Config{
		Conn: conn,
	}), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DatabaseImpl{res}, nil
}

func NewPostgresEnv(
	env *env.Environment,
	log logger.Logger,
) (*DatabaseImpl, error) {
	defer resolver(log)

	res, err := gorm.Open(postgres.Open(env.DatabaseConnection), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &DatabaseImpl{res}, nil
}

func (p DatabaseImpl) Connect() *gorm.DB {
	return p.DB
}
