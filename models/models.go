package models

import (
	"github.com/jinzhu/gorm"
)

type Sighting struct {
	gorm.Model
	Registration       string `json:"registration"`
	PlaneConfiguration string `json:"plane_configuration"`
	PlaneStatus        string `json:"plane_status"`
	Venue              Venue
	Photo              string `json:"photo"`
}

type Venue struct {
	gorm.Model
	Name     string  `json:"name"`
	IATACode string  `json:"iata_code"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}
