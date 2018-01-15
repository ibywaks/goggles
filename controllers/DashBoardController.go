package controllers

import (
	"goggles/models"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

//DashBoardControllers - Controller object for Endpoints dashboard
type DashBoardControllers struct {
	mvc.BaseController
	Cntx iris.Context
}

//ShowEndpoints - show list of endpoints
func (d DashBoardControllers) ShowEndpoints() {
	endpoints := (models.EndPoints{}).GetWithLastCall()

	d.Cntx.ViewData("endpoints", endpoints)
	d.Cntx.View("endpoints.html")
}
