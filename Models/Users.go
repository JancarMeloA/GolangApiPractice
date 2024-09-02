package models

import (
	"time"
)

type Users struct {
	IdUser       uint64 `gorm:"primaryKey;autoIncrement"`
	Name         string `gorm:"type:varchar(70);"`
	LastName     string `gorm:"type:varchar(110)"`
	Email        string `gorm:"type:varchar(120);unique_index;not null"`
	Password     string `gorm:"type:varchar(120);not null"`
	DateCreation time.Time
}
