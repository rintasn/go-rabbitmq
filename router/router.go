package router

import (
	"github.com/gorilla/mux"
	"main/controller"
)

func Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/api/register", controller.RegisterAccount).Methods("POST", "OPTIONS")

	return router
}
