package models

import (
	"github.com/jinzhu/gorm"
)

//EndPoints - endpoint model
type EndPoints struct {
	gorm.Model
	Name, URL string
	Type      string          `gorm:"DEFAULT:'GET'"`
	Calls     []EndPointCalls `gorm:"ForeignKey:EndPointID"`
}

//EndPointsWithLastCall - Endpoints with last call
type EndPointsWithLastCall struct {
	Name, URL   string
	Type        string
	LastStatus  int
	NumRequests int
}

var db, _ = gorm.Open("sqlite3", "./db/gorm.db")

//GetWithLastCall - get all endpoints with last call details
func (ep EndPoints) GetWithLastCall() []EndPointsWithLastCall {
	var eps []EndPoints
	var epsWithDets []EndPointsWithLastCall

	db.Preload("Calls").Find(&eps)

	for _, elem := range eps {
		calls := elem.Calls
		lastCall := calls[len(calls)-1:][0]

		newElem := EndPointsWithLastCall{
			elem.Name,
			elem.URL,
			elem.Type,
			lastCall.ResponseCode,
			len(calls),
		}

		epsWithDets = append(epsWithDets, newElem)
	}

	return epsWithDets
}

//SaveOrCreate - save endpoint called
func (ep EndPoints) SaveOrCreate() EndPoints {
	db.FirstOrCreate(&ep, ep)
	return ep
}
