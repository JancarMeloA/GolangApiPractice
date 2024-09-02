package routes

import (
	"net/http"

	controllers "github.com/Jan/GolangApiPractice/Controllers"
)

func LibraryRoutes() {

	http.HandleFunc("/AllLibrary", controllers.AllLibrary)
	http.HandleFunc("/AllLibrarydetails", controllers.AllLibrarydetails)
	http.HandleFunc("/CreateLibrary", controllers.CreateLibrary)
}
