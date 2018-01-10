package models

import (
	"github.com/jinzhu/gorm"
)

//EndPoints - endpoint model
type EndPoints struct {
	gorm.Model
	Name, URL string
	Type      string `gorm:"DEFAULT:'GET'"`
	Payload   string `gorm:"type:text"`
	Calls 	  int64
}

var db, _ = gorm.Open("sqlite3", "./db/gorm.db")

//Get - get all endpoints
func (ep EndPoints) Get() EndPoints {
	db.Order("calls DESC").Find(&ep)

	return ep
}

//SaveCount - save endpoint call counts
func (ep EndPoints) SaveCount() {
	db.FirstOrCreate(&ep)
	ep.Calls = (ep.Calls + 1)
	db.Save(&ep)
}
