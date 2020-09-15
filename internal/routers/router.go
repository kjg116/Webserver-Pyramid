package routers

import (
	"github.com/gorilla/mux"
	"github.com/kjg116/webserver-pyramid/internal/controllers"
)

// InitRoutes ..
func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.Path("/pyramid").HandlerFunc(controllers.AreTheyPyramids).Methods("POST")
	return router
}
