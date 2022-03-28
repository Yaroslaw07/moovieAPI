package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/Yaroslaw07/go_in_practice/pkg/models"
	"github.com/Yaroslaw07/go_in_practice/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var newFilm models.Film

func GetFilms(w http.ResponseWriter, r *http.Request) {
	newFilms := models.GetAllFilms()
	res, _ := json.Marshal(newFilms)
	w.Header().Set("Content-Type", "pkqlication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetFilmById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["filmId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	filmDetails, _ := models.GetFilmById(ID)
	res, _ := json.Marshal(filmDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateFilm(w http.ResponseWriter, r *http.Request) {
	CreateFilm := &models.Film{}
	utils.ParseBody(r, CreateFilm)
	b := CreateFilm.CreateFilm()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteFilm(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filmId := vars["filmId"]
	ID, err := strconv.ParseInt(filmId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	film := models.DeleteFilm(ID)
	res, _ := json.Marshal(film)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateFilm(w http.ResponseWriter, r *http.Request) {
	var updateFilm = &models.Film{}
	utils.ParseBody(r, updateFilm)
	vars := mux.Vars(r)
	bookId := vars["filmId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	booksDetails, db := models.GetFilmById(ID)
	if updateFilm.Name != "" {
		booksDetails.Name = updateFilm.Name
	}
	if updateFilm.Producer != "" {
		booksDetails.Producer = updateFilm.Producer
	}
	if updateFilm.Duration != 0 {
		booksDetails.Duration = updateFilm.Duration
	}
	db.Save(&booksDetails)
	res, _ := json.Marshal(booksDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
