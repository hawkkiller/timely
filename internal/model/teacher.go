package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Name       string
	Email      sql.NullString `gorm:"unique"`
	Phone      sql.NullString `gorm:"unique"`
	Conference Conference
	Subjects   []*Subject `gorm:"many2many:teacher_subjects;"`
}

type TeacherInsertDTO struct {
	Name       string              `json:"name"`
	Email      string              `json:"email"`
	Phone      string              `json:"phone"`
	Conference ConferenceInsertDTO `json:"conference"`
}
