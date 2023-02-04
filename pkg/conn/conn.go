package conn

import (
	"github.com/hawkkiller/timely/internal/model"
	"github.com/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	conn, err := gorm.Open(sqlite.Open("timelydb/timely.db"))
	if err != nil {
		return nil, errors.Wrap(err, "failed to open database")
	}
	err = conn.AutoMigrate(&model.Schedule{}, &model.Subject{}, &model.Teacher{}, &model.Conference{})
	if err != nil {
		return nil, errors.Wrap(err, "failed to migrate database")
	}
	return conn, nil
}
