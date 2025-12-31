package models

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Driver struct {
	Id        bson.ObjectID `bson:"_id,omitempty"`
	FirstName string        `bson:"firstName"`
	LastName  string        `bson:"lastName"`
	Plate     string        `bson:"plate"`
	TaxiType  string        `bson:"taxiType"`
	CarBrand  string        `bson:"carBrand"`
	CarModel  string        `bson:"carModel"`
	Latitude  float64       `bson:"lat"`
	Longitude float64       `bson:"lon"`
	CreatedAt time.Time     `bson:"createdAt"`
	UpdatedAt time.Time     `bson:"updatedAt"`
}

type PaginationMetadata struct {
	Page       int32 `json:"page"`
	PageSize   int32 `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int64 `json:"totalPages"`
}

type NearbyDriver struct {
	FirstName  string  `json:"firstName" bson:"firstName"`
	LastName   string  `json:"lastName" bson:"lastName"`
	Plate      string  `json:"plate" bson:"plate"`
	DistanceKm float64 `json:"distanceKm" bson:"distanceKm"`
}
