package models

import (
	"time"
)

type Library struct {
	IdLibrary    uint64 `gorm:"primaryKey;autoIncrement;"`
	IdUser       uint64 `gorm:"primaryKey; autoIncrement:false"`
	DateCreation time.Time
}
