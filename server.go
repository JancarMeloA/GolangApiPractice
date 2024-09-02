package main

import (
	"log"
	"net/http"
	"time"

	DB "github.com/Jan/GolangApiPractice/DB"
	migrations "github.com/Jan/GolangApiPractice/DB/Migrations"
	routes "github.com/Jan/GolangApiPractice/Routes"
)

func main() {

	DB.DbConnection()

	migrations.Auto()

	routes.Run()

	s := &http.Server{
		Addr:         ":8000",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
