package dto

import (
	"time"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
)

type DriverResponse struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Plate     string    `json:"plate"`
	TaxiType  string    `json:"taxiType"`
	CarBrand  string    `json:"carBrand"`
	CarModel  string    `json:"carModel"`
	Latitude  float64   `json:"lat"`
	Longitude float64   `json:"lon"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func ProtoToDriverResponse(driver *pb.Driver) *DriverResponse {
	return &DriverResponse{
		Id:        driver.Id,
		FirstName: driver.FirstName,
		LastName:  driver.LastName,
		Plate:     driver.Plate,
		TaxiType:  driver.TaxiType,
		CarBrand:  driver.CarBrand,
		CarModel:  driver.CarModel,
		Latitude:  driver.Latitude,
		Longitude: driver.Longitude,
		CreatedAt: driver.CreatedAt.AsTime(),
		UpdatedAt: driver.UpdatedAt.AsTime(),
	}
}
