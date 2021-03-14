package db

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewDB(driver, args string) (*gorm.DB, error) {
	var db *gorm.DB
	var err error
	switch driver {
	case "mysql":
		db, err = gorm.Open(mysql.Open(args), nil)
		if err != nil {
			return nil, errors.Wrap(err, "open db error")
		}
	case "sqlite":
		db, err = gorm.Open(sqlite.Open(args), nil)
		if err != nil {
			return nil, errors.Wrap(err, "open db error")
		}
	case "":
		return nil, errors.New("driver cannot be empty")
	default:
		return nil, errors.Errorf("unknown driver: %s", driver)
	}

	return db, nil
}
