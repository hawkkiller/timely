package service

import (
	"github.com/hawkkiller/timely/internal/db"
	"github.com/hawkkiller/timely/internal/model"
	"net/http"
)

type ScheduleService interface {
	InsertSchedule(link string) error
	GetSchedule() *model.Schedule
}

type scheduleService struct {
	dataProvider db.ScheduleDB
	client       *http.Client
}

func NewScheduleService(provider db.ScheduleDB) ScheduleService {
	return &scheduleService{
		dataProvider: provider,
		client:       &http.Client{},
	}
}
