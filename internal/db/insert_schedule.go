package db

import "github.com/hawkkiller/timely/internal/model"

func (db scheduleDB) InsertSchedule(schedule *model.Schedule) error {
	err := db.conn.Debug().Create(schedule).Error
	return err
}
