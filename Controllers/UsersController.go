package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var users []models.Users
		if err := DB.DB.Find(&users).Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&users)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var user models.Users
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		defer r.Body.Close()
		createUser := DB.DB.Create(&user)
		err := createUser.Error

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(&user)

	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {

		var user models.Users
		param := r.URL.Path

		limit := len(param)

		Id, err := strconv.ParseInt(param[12:limit], 10, 0)

		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		defer r.Body.Close()
		DB.DB.Model(&models.Users{}).Where("id_user = ?", Id).Updates(models.Users{Name: user.Name, LastName: user.LastName, Email: user.Email, Password: user.Password})
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("se ha actualizado el registro "))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodDelete {

		var user []models.Users
		param := r.URL.Path
		limit := len(param)

		Id, err := strconv.ParseInt(param[12:limit], 10, 0)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}

		if err := DB.DB.Where("id_user=?", Id).Delete(&user).Error; err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("se ha borrado exitosamente el registro"))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}
