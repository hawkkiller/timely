package model

import (
	"time"

	"gorm.io/gorm"
)

type Subject struct {
	gorm.Model
	Title      string
	Teachers   []*Teacher `gorm:"many2many:teacher_subjects;"`
	StartTime  time.Time
	EndTime    time.Time
	ScheduleID uint
}

type SubjectInsertDTO struct {
	Title     string             `json:"title"`
	StartTime string             `json:"start_time"`
	EndTime   string             `json:"end_time"`
	Teachers  []TeacherInsertDTO `json:"teachers"`
}
