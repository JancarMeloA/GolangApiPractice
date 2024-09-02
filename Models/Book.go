package models

import (
	"time"
)

type Book struct {
	IdBook       uint64 `gorm:"primaryKey;autoIncrement;not null"`
	Title        string `gorm:"type:varchar(120); not null"`
	Gender       string `gorm:"type:varchar(50); not null"`
	FrontPage    string
	DateCreation time.Time
}
