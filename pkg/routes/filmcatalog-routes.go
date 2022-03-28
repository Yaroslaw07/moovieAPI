package routes

import (
	"github.com/Yaroslaw07/go_in_practice/pkg/controllers"
	_ "github.com/Yaroslaw07/go_in_practice/pkg/controllers"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
)

var RegisterMoovieRoutes = func(router *mux.Router) {
	router.HandleFunc("/film/", controllers.CreateFilm).Methods("POST")
	router.HandleFunc("/film/", controllers.GetFilms).Methods("GET")
	router.HandleFunc("/film/{filmId}", controllers.GetFilmById).Methods("GET")
	router.HandleFunc("/film/{filmId}", controllers.UpdateFilm).Methods("PUT")
	router.HandleFunc("/film/{filmId}", controllers.DeleteFilm).Methods("DELETE")
}
