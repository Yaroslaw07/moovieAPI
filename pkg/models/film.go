package models

import (
	"github.com/Yaroslaw07/go_in_practice/pkg/config"
	"github.com/jinzhu/gorm"
	"time"
)

var db *gorm.DB

type Film struct {
	gorm.Model
	Name     string        `gorm:""json:"name"`
	Producer string        `json:"producer"`
	Duration time.Duration `json:"duration"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Film{})
}

func (f *Film) CreateFilm() *Film {
	db.NewRecord(f)
	db.Create(&f)
	return f
}

func GetAllFilms() []Film {
	var films []Film
	db.Find(&films)
	return films
}

func GetFilmById(Id int64) (*Film, *gorm.DB) {
	var film Film
	db := db.Where("ID=?", Id).Find(&film)
	return &film, db
}

func DeleteFilm(ID int64) Film {
	var film Film
	db.Where("ID=?", ID).Delete(film)
	return film
}
