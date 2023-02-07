package user

import (
	"database/sql/driver"
	"errors"
)

type UserRole string

const (
	Admin       UserRole = "admin"
	Student     UserRole = "student"
	Contributor UserRole = "contributor"
)

func (ur *UserRole) Scan(value interface{}) error {
	val := UserRole(value.([]byte))

	if val != Contributor && val != Student && val != Admin {
		return errors.New("invalid role")
	}

	*ur = val
	return nil
}

func (ur UserRole) Value() (driver.Value, error) {
	return string(ur), nil
}
