package models

type ListDriversResponse struct {
	Drivers []Driver  `json:"drivers"`
	Meta    *Metadata `json:"meta,omitempty"`
}

type Metadata struct {
	Page       int `json:"page"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
	TotalPages int `json:"totalPages"`
}

type NearbyDrivers struct {
	FirstName  string  `json:"firstName" bson:"firstName"`
	LastName   string  `json:"lastName" bson:"lastName"`
	Plate      string  `json:"plate" bson:"plate"`
	DistanceKm float64 `json:"distanceKm" bson:"distanceKm"`
}
