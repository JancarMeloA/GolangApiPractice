package models

type LibraryBook struct {
	IdLibrary uint64 `gorm:"primaryKey;autoIncrement:false"`
	IdBook    uint64 `gorm:"primaryKey;autoIncrement:false"`
}
