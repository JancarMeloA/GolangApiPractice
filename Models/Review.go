package models

import (
	"time"
)

type Review struct {
	IdReview     uint64    `gorm:"primaryKey;autoIncrement;"`
	IdBook       uint64    `gorm:"primaryKey;autoIncrement:false;"`
	IdUser       uint64    `gorm:"primaryKey;autoIncrement:false;"`
	Content      string    `gorm:"type:varchar(700)"`
	DateCreation time.Time `gorm:"type:date"`
}
