package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Jan/GolangApiPractice/DB"
	models "github.com/Jan/GolangApiPractice/Models"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.Users
	DB.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var user models.Users
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Error: Datos inválidos"))
		}
		defer r.Body.Close()
		createUser := DB.DB.Create(&user)
		err := createUser.Error

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		json.NewEncoder(w).Encode(&user)

	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	param := r.URL.Path
	Id := param[12:13]
	fmt.Println(Id)
}
