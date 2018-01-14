package models

import (
	"github.com/kataras/iris"
	"github.com/jinzhu/gorm"
)

//EndPointsCall - Object for storing endpoints call details
type EndPointsCall struct {
	gorm.Model
	EndPointID uint
	RequestIP  string
	ResponseCode int
}

//SaveCall - Save the call details of an endpoint
func (ep EndPoints) SaveCall(context iris.Context) EndPointsCall {
	epCall := EndPointsCall{
		EndPointID: ep.ID,
		RequestIP: context.RemoteAddr(),
		ResponseCode: context.GetStatusCode(),
	}

	db.Create(&epCall)
	return epCall
}