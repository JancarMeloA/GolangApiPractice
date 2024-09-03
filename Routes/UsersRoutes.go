package routes

import (
	"net/http"

	controllers "github.com/Jan/GolangApiPractice/Controllers"
)

func UserRoutes() {
	http.HandleFunc("/registerUser", controllers.RegisterUser)
	http.HandleFunc("/allUsers", controllers.AllUsers)
	http.HandleFunc("/UpdateUser/{id}", controllers.UpdateUser)
	http.HandleFunc("/DeleteUser/{id}", controllers.DeleteUser)
}
