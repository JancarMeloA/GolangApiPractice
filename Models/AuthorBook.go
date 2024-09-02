package models

type AuthorBook struct {
	IdAuthor uint64 `gorm:"primaryKey;autoIncrement:false"`
	IdBook   uint64 `gorm:"primaryKey;autoIncrement:false"`
}
