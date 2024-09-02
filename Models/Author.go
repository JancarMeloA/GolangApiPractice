package models

import (
	"time"
)

type Author struct {
	IdAuthor       uint64 `gorm:"primaryKey;autoIncrement;"`
	NameAuthor     string `gorm:"type:varchar(70)"`
	LastNameAuthor string `gorm:"type:varchar(70)"`
	DateCreation   time.Time
}
