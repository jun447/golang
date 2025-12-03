package router

import (
	"github.com/gorilla/mux"
	"github.com/jun447/mongoApi/controller"
)

func Router() *mux.Router {
	router:=mux.NewRouter()

	router.HandleFunc("/api/addMovie",controller.InsertMovie).Methods("POST")
	router.HandleFunc("/api/getMovies",controller.GetAllMovies).Methods("GET")
	router.HandleFunc("/api/getMovie/{id}",controller.GetMovieById).Methods("GET")
	router.HandleFunc("/api/deleteMovie/{id}",controller.DeleteMovieByID).Methods("DELETE")
	router.HandleFunc("/api/updateMovie/{id}",controller.UpdateMovieByID).Methods("PUT")

	return router
}