package controllers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// MoviesController - serve movie data
type MoviesController struct {
	mvc.BaseController
	Cntx iris.Context
}

//Get - get a list of all available movies
func (m MoviesController) Get() {
	m.Cntx.JSON(iris.Map{"status": "success"})
}

//GetByID - Get movie by ID
func (m MoviesController) GetByID(ID string) {
	m.Cntx.JSON(iris.Map{"status": "success"})
}

//Add - Add new movie
func (m MoviesController) Add() {

}

//Edit - Edit a movie
func (m MoviesController) Edit(ID string) {

}

//Delete - delete a specific movie
func (m MoviesController) Delete(ID string) {

}
