package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	var Books []models.Book
	DB.DB.Find(&Books)
	json.NewEncoder(w).Encode(&Books)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var Book models.Book
	if err := json.NewDecoder(r.Body).Decode(&Book); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Datos inválidos"))
	}
	defer r.Body.Close()
	newBook := DB.DB.Create(&Book)

	err := newBook.Error

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&Book)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var libraryBook models.LibraryBook

	if err := json.NewDecoder(r.Body).Decode(&libraryBook); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Datos inválidos"))
		return
	}
	defer r.Body.Close()
	AgglibraryBook := DB.DB.Create(&libraryBook)

	err := AgglibraryBook.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&libraryBook)
}
