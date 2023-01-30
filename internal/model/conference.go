package model

import (
	"database/sql"

	"gorm.io/gorm"
)

type Conference struct {
	gorm.Model
	// identifier is the unique identifier of the conference
	Identifier sql.NullString
	// link is the link to the conference
	Link      sql.NullString
	Password  sql.NullString
	TeacherID uint
}

type ConferenceInsertDTO struct {
	Identifier string `json:"identifier"`
	Link       string `json:"link"`
	Password   string `json:"password"`
}
