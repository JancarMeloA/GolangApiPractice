package migrations

import (
	DB "github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func Auto() {
	DB.DB.AutoMigrate(&models.Users{}, &models.Book{}, &models.LibraryBook{})
	DB.DB.AutoMigrate(&models.Library{}, &models.Review{})
	DB.DB.AutoMigrate(&models.AuthorBook{}, &models.Author{})
	
}
