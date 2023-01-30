package model

import (
	"database/sql"

	"gorm.io/gorm"
)

// Schedule represents a timetable for the day
// Each day can have a schedule
type Schedule struct {
	gorm.Model
	Title string
	// DayOfWeek is the day of the week
	DayOfWeek uint8
	// CreatedBy is the id of the user who created the schedule
	CreatedBy sql.NullInt64
	Subjects  []*Subject
}

type ScheduleInsertDTO struct {
	Title     string `json:"title"`
	CreatedBy int64  `json:"created_by,omitempty"`
	DayOfWeek uint8  `json:"day_of_week"`
	Subjects  []SubjectInsertDTO
}
