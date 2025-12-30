package models

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type AddDriverRequest struct {
	FirstName string  `json:"firstName" validate:"required,min=2,max=50"`
	LastName  string  `json:"lastName" validate:"required,min=2,max=50"`
	Plate     string  `json:"plate" validate:"required,min=6,max=9"`
	TaxiType  string  `json:"taxiType" validate:"required,oneof=sari korsan uber tag"`
	CarModel  string  `json:"carModel" validate:"required"`
	CarBrand  string  `json:"carBrand" validate:"required"`
	Latitude  float64 `json:"lat" validate:"required,latitude"`
	Longitude float64 `json:"lon" validate:"required,longitude"`
}

func (req *AddDriverRequest) ToDriver() Driver {
	return Driver{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

type ListDriversRequest struct {
	Page     int `query:"page" validate:"omitempty,gte=0"`
	PageSize int `query:"pageSize" validate:"omitempty,gte=0,max=100"`
}

type ListNearbyDriversRequest struct {
	Latitude  float64 `query:"lat" validate:"required,latitude"`
	Longitude float64 `query:"lon" validate:"required,longitude"`
	TaxiType  string  `query:"taxiType" validate:"required,oneof=sari korsan uber tag"`
}

type UpdateDriverRequest struct {
	FirstName *string   `json:"firstName,omitempty" bson:"firstName,omitempty"`
	LastName  *string   `json:"lastName,omitempty" bson:"lastName,omitempty"`
	Plate     *string   `json:"plate,omitempty" bson:"plate,omitempty"`
	TaxiType  *string   `json:"taxiType,omitempty" bson:"taxiType,omitempty"`
	CarModel  *string   `json:"carModel,omitempty" bson:"carModel,omitempty"`
	CarBrand  *string   `json:"carBrand,omitempty" bson:"carBrand,omitempty"`
	Latitude  *float64  `json:"lat,omitempty" bson:"latitude,omitempty"`
	Longitude *float64  `json:"lon,omitempty" bson:"longitude,omitempty"`
	UpdatedAt time.Time `json:"-" bson:"updatedAt"`
}

func (req *UpdateDriverRequest) ToUpdateDoc() (bson.M, error) {
	req.UpdatedAt = time.Now()

	data, err := bson.Marshal(req)
	if err != nil {
		return nil, err
	}

	update := bson.M{}
	if err := bson.Unmarshal(data, &update); err != nil {
		return nil, err
	}

	if len(update) == 0 {
		return nil, errors.New("no fields to update")
	}

	return update, nil
}
