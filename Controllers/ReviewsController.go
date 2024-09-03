package controllers

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllReviews(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		var reviewdetails []struct {
			Id_Review     uint64 `json:"id"`
			Title         string `json:"title_book"`
			Gender        string `json:"gender_book"`
			Content       string `json:"comment"`
			Date_creation string `json:"date_creation"`
		}
		DB.DB.Table("book b").Select("id_review,title,gender,content,rw.date_creation").Joins("JOIN review rw ON rw.id_book = b.id_book").Joins("JOIN users us ON us.id_user = rw.id_user").Scan(&reviewdetails)

		json.NewEncoder(w).Encode(reviewdetails)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func CreateReview(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var review models.Review
		if err := json.NewDecoder(r.Body).Decode(&review); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()
		createReview := DB.DB.Create(&review)

		if err := createReview.Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&review)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}

}

func UpdateReview(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPut {
		var Newreview models.Review
		param := r.URL.Path
		limit := len(param)

		Id, err := strconv.ParseInt(param[14:limit], 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		if err := json.NewDecoder(r.Body).Decode(&Newreview); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		defer r.Body.Close()

		if err := DB.DB.Where("id_review = ?", Id).First(models.Review{}).Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		if err := DB.DB.Table("review").Where("id_review = ?", Id).Update("content", Newreview.Content).Error; err != nil {
			w.Write([]byte(err.Error()))
		}
		json.NewEncoder(w).Encode(&Newreview)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Error: esta ruta no existe."))
	}

}

func DeleteReview(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {

		var review []models.Review
		param := r.URL.Path
		limit := len(param)
		Id, err := strconv.ParseInt(param[12:limit], 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Error: como 'Id' se permiten solo numeros."))
			return
		}

		if validId := DB.DB.First(review).Where("id_review = ?", Id).Error; validId != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: Este 'Id' no existe."))
			return
		}

		if err := DB.DB.Where("id_review = ?", Id).Delete(&review).Error; err != nil {
			w.Write([]byte("Error: no se pudo eliminar la review."))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("se ha borrado con exito."))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
