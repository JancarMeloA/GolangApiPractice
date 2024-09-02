package routes

import (
	"net/http"

	controllers "github.com/Jan/GolangApiPractice/Controllers"
)

func ReviewRoutes() {
	http.HandleFunc("/AllReviews", controllers.AllReviews)
	// create review of a book
	http.HandleFunc("/createReview", controllers.CreateReview)
	//test put review
	http.HandleFunc("/UpdateReview/{id}", controllers.UpdateReview)
	http.HandleFunc("/DeleteReview/{id}", controllers.DeleteReview)

}
