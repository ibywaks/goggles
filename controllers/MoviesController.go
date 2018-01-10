package controllers

import (
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
	AppId: "your_app_id",
	Key: "your_app_key",
	Secret: "your_app_secret",
	Cluster: "your_app_cluster",
}

//Get - get a list of all available movies
func (m MoviesController) Get() {
	movie := models.Movies{}
	movies := movie.Get()

	m.saveEndpointCall("Movies List")

	m.Cntx.JSON(iris.Map{"status": "success", "data": movies})
}

//GetByID - Get movie by ID
func (m MoviesController) GetByID(ID int64) {
	movie := models.Movies{}
	movie = movie.GetByID(ID)

	m.saveEndpointCall("Single Movie Retrieval")
	
	m.Cntx.JSON(iris.Map{"status": "success", "data": movie})
}

//Add - Add new movie
func (m MoviesController) Add() {
	movie := models.Movies{}
	m.Cntx.ReadForm(&movie)

	m.saveEndpointCall("Add Movie")

	movie.Create()
	m.Cntx.JSON(iris.Map{"status":"success", "data": movie})
}

//Edit - Edit a movie
func (m MoviesController) Edit(ID int64) {
	movie := models.Movies{}
	movie = movie.GetByID(ID)
	m.Cntx.ReadForm(&movie)

	m.saveEndpointCall("Single Movie Edit")

	movie.Edit()
	m.Cntx.JSON(iris.Map{"status":"success", "data": movie})	
}

//Delete - delete a specific movie
func (m MoviesController) Delete(ID int64) {
	movie := models.Movies{}
	movie = movie.GetByID(ID)

	m.saveEndpointCall("Single Movie Delete")

	movie.Delete()
	m.Cntx.JSON(iris.Map{"status":"success", "message": "Movie with ID: "})
}

func (m MoviesController) saveEndpointCall(name string) {

	endpoint := models.EndPoints{
		Name: name,
		URL: m.Cntx.Path(),
		Type: m.Cntx.Request().Method,
	}

	endpoint.SaveCount()
}
