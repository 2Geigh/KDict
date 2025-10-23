package routes

import (
	"KDict/src/controllers"
	"net/http"
)

func RegisterDictRoutes() {

	http.HandleFunc("/", controllers.RootHandler)

	http.HandleFunc("/results", controllers.ResultsHandler)
}
