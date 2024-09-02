package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllLibrary(w http.ResponseWriter, r *http.Request) {
	var Library []models.Library
	DB.DB.Find(&Library)
	json.NewEncoder(w).Encode(&Library)
}

func AllLibrarydetails(w http.ResponseWriter, r *http.Request) {

	var libraryDetails []struct {
		IDLibrary   uint64    `json:"id_library"`
		IDUser      uint64    `json:"id_user"`
		LibraryDate time.Time `json:"library_date"`
		IDBook      uint64    `json:"id_book"`
		Title       string    `json:"title_book"`
		Gender      string    `json:"gender_book"`
		FrontPage   string    `json:"front_page"`
		BookDate    time.Time `json:"book_date"`
	}
	result := DB.DB.Table("library l").
		Select("l.id_library, l.id_user, l.date_creation, b.id_book, b.title, b.gender, a.name_author, a.last_name_author, b.front_page, b.date_creation AS book_date").
		Joins("JOIN library_book lb ON l.id_library = lb.id_library").
		Joins("JOIN book b ON b.id_book = lb.id_book").
		Joins("JOIN author_book ab ON ab.id_book = b.id_book").
		Joins("JOIN author a ON ab.id_author = a.id_author").
		Scan(&libraryDetails)

	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(""))
	}
	json.NewEncoder(w).Encode(&libraryDetails)
}

func CreateLibrary(w http.ResponseWriter, r *http.Request) {
	var Library models.Library
	if err := json.NewDecoder(r.Body).Decode(&Library); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Datos inválidos"))
	}
	defer r.Body.Close()
	CreateLibrary := DB.DB.Create(&Library)

	if err := CreateLibrary.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Library)
}
