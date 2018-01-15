package controllers

import (
	"fmt"
	"goggles/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/pusher/pusher-http-go"
)

// MoviesController - serve movie data
type MoviesController struct {
	mvc.BaseController
	Cntx iris.Context
}

var client = pusher.Client{
	AppId:   "your_app_id",
	Key:     "your_app_key",
	Secret:  "your_app_secret",
	Cluster: "your_app_cluster",
}

//Get - get a list of all available movies
func (m MoviesController) Get() {
	movie := models.Movies{}
	movies := movie.Get()

	go m.saveEndpointCall("Movies List")

	m.Cntx.JSON(iris.Map{"status": "success", "data": movies})
}

//GetByID - Get movie by ID
func (m MoviesController) GetByID(ID int64) {
	movie := models.Movies{}
	movie = movie.GetByID(ID)

	if !movie.Validate() {
		msg := fmt.Sprintf("Movie with ID: %v not found", ID)
		m.Cntx.StatusCode(iris.StatusNotFound)
		m.Cntx.JSON(iris.Map{"status": "error", "message": msg})
	} else {
		m.Cntx.JSON(iris.Map{"status": "success", "data": movie})
	}

	name := fmt.Sprintf("Single Movie with ID: %v Retrieval", ID)
	go m.saveEndpointCall(name)
}

//Add - Add new movie
func (m MoviesController) Add() {
	movie := models.Movies{}
	m.Cntx.ReadForm(&movie)

	if !movie.Validate() {
		m.Cntx.StatusCode(iris.StatusBadRequest)
		m.Cntx.JSON(iris.Map{"status": "error", "message": "Movie not added"})
	} else {
		movie.Create()
		m.Cntx.JSON(iris.Map{"status": "success", "data": movie})
	}

	go m.saveEndpointCall("Add Movie")
}

//Edit - Edit a movie
func (m MoviesController) Edit(ID int64) {
	movie := models.Movies{}
	movie = movie.GetByID(ID)
	m.Cntx.ReadForm(&movie)

	if !movie.Validate() {
		msg := fmt.Sprintf("Movie with ID: %v not found", ID)
		m.Cntx.StatusCode(iris.StatusNotFound)
		m.Cntx.JSON(iris.Map{"status": "error", "message": msg})
	} else {
		movie.Edit()
		m.Cntx.JSON(iris.Map{"status": "success", "data": movie})
	}

	name := fmt.Sprintf("Single Movie with ID: %v Edit", ID)
	go m.saveEndpointCall(name)

}

//Delete - delete a specific movie
func (m MoviesController) Delete(ID int64) {
	movie := models.Movies{}
	movie = movie.GetByID(ID)

	if !movie.Validate() {
		msg := fmt.Sprintf("Movie with ID: %v not found", ID)
		m.Cntx.StatusCode(iris.StatusNotFound)
		m.Cntx.JSON(iris.Map{"status": "error", "message": msg})
	} else {
		movie.Delete()
		m.Cntx.JSON(iris.Map{"status": "success", "message": "Movie with ID: "})
	}

	name := fmt.Sprintf("Single Movie with ID: %v Delete", ID)
	go m.saveEndpointCall(name)
}

func (m MoviesController) saveEndpointCall(name string) {

	endpoint := models.EndPoints{
		Name: name,
		URL:  m.Cntx.Path(),
		Type: m.Cntx.Request().Method,
	}

	endpoint = endpoint.SaveOrCreate()

	endpoint.SaveCall(m.Cntx)
}
