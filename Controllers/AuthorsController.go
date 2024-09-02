package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllAuthors(w http.ResponseWriter, r *http.Request) {
	var Authors []models.Author
	DB.DB.Find(&Authors)
	json.NewEncoder(w).Encode(&Authors)
}

func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var Author models.Author
	if err := json.NewDecoder(r.Body).Decode(&Author); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Datos inválidos"))
		return
	}
	defer r.Body.Close()
	NewAuthor := DB.DB.Create(&Author)

	if err := NewAuthor.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&Author)
}

func AddAuthor(w http.ResponseWriter, r *http.Request) {
	var authorBook models.AuthorBook
	fmt.Println(authorBook.IdAuthor)
	if err := json.NewDecoder(r.Body).Decode(&authorBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Datos inválidos"))
		return
	}
	defer r.Body.Close()
	var author models.Author

	result := DB.DB.Find(&author, "id_author = ?", authorBook.IdAuthor)
	if result.Error != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: El autor no existe"))
		return
	}

	if err := DB.DB.Create(&authorBook).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&authorBook)
}
