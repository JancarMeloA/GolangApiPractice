package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllReviews(w http.ResponseWriter, r *http.Request) {
	var reviewdetails []struct {
		Id_Review     uint64 `json:"id"`
		Title         string `json:"title_book"`
		Gender        string `json:"gender_book"`
		Content       string `json:"comment"`
		Date_creation string `json:"date_creation"`
	}
	DB.DB.Table("book b").Select("id_review,title,gender,content,rw.date_creation").Joins("JOIN review rw ON rw.id_book = b.id_book").Joins("JOIN users us ON us.id_user = rw.id_user").Scan(&reviewdetails)

	json.NewEncoder(w).Encode(reviewdetails)

}

func CreateReview(w http.ResponseWriter, r *http.Request) {
	var review models.Review

	if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error: Datos inválidos"))
		return
	}
	defer r.Body.Close()
	createReview := DB.DB.Create(&review)

	if err := createReview.Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&review)

}

func UpdateReview(w http.ResponseWriter, r *http.Request) {

	var Newreview models.Review
	param := r.URL.Path
	fmt.Println(param)
	id := param[14:15]

	if r.Method == http.MethodPut {
		if err := json.NewDecoder(r.Body).Decode(&Newreview); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: Datos inválidos"))
			return
		}
		defer r.Body.Close()
		err := DB.DB.Table("review").Where("id_review = ?", id).Update("content", Newreview.Content).Error
		if err != nil {
			w.Write([]byte("Error DataBase"))
		}
		json.NewEncoder(w).Encode(&Newreview)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: esta ruta no existe"))
	}

}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	var review []models.Review
	param := r.URL.Path
	id := param[14:15]

	if err := DB.DB.Where("id_review = ?", id).Delete(&review).Error; err != nil {
		w.Write([]byte("Error Delete review"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("success"))
}
