package routes

import (
	"net/http"

	controllers "github.com/Jan/GolangApiPractice/Controllers"
)

func BookRoutes() {
	http.HandleFunc("/createBook", controllers.CreateBook)
	http.HandleFunc("/AddBook", controllers.AddBook)
	http.HandleFunc("/allBooks", controllers.AllBooks)
}
