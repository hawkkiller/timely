package db

import (
	"github.com/hawkkiller/timely/internal/model"
	"github.com/mymmrac/telego"
	"gorm.io/gorm"
)

type ScheduleDB interface {
	InsertSchedule(schedule *model.ScheduleInsertDTO)
	GetSchedule(chatID telego.ChatID) *model.Schedule
}

type scheduleDB struct {
	conn *gorm.DB
}

func NewScheduleDB(conn *gorm.DB) ScheduleDB {
	return &scheduleDB{
		conn: conn,
	}
}
