package dto

import (
	"time"
)

type PaginationMetadataResponse struct {
	Page       int32 `json:"page"`
	PageSize   int32 `json:"pageSize"`
	TotalCount int64 `json:"totalCount"`
	TotalPages int64 `json:"totalPages"`
}

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

type NearbyDriverResponse struct {
	FirstName  string  `json:"firstName" bson:"firstName"`
	LastName   string  `json:"lastName" bson:"lastName"`
	Plate      string  `json:"plate" bson:"plate"`
	DistanceKm float64 `json:"distanceKm" bson:"distanceKm"`
}

type GetDriversResponse struct {
	Drivers            []*DriverResponse           `json:"drivers"`
	PaginationMetadata *PaginationMetadataResponse `json:"paginationMetadata,omitempty"`
}

type GetNearbyDriversResponse struct {
	Drivers []*NearbyDriverResponse `json:"drivers"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
	Time    string `json:"time"`
}

type CreateTokenResponse struct {
	Token string `json:"token"`
}
