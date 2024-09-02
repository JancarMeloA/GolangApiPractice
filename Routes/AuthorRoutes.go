package routes

import (
	"net/http"

	controllers "github.com/Jan/GolangApiPractice/Controllers"
)

func AuthorRoutes() {
	http.HandleFunc("/CreateAuthor", controllers.CreateAuthor)
	http.HandleFunc("/AddAuthor", controllers.AddAuthor)
}
