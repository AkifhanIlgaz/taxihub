package dto

import (
	"time"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
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

func AddDriverRequestToProto(req *AddDriverRequest) *pb.AddDriverRequest {
	return &pb.AddDriverRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarBrand:  req.CarBrand,
		CarModel:  req.CarModel,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}

type GetDriversRequest struct {
	Page     int `query:"page" validate:"omitempty,gte=0"`
	PageSize int `query:"pageSize" validate:"omitempty,gte=0,max=100"`
}

func GetDriversRequestToProto(req *GetDriversRequest) *pb.GetDriversRequest {
	return &pb.GetDriversRequest{
		Page:     int32(req.Page),
		PageSize: int32(req.PageSize),
	}
}

type GetNearbyDriversRequest struct {
	Latitude  float64 `query:"lat" validate:"required,latitude"`
	Longitude float64 `query:"lon" validate:"required,longitude"`
	TaxiType  string  `query:"taxiType" validate:"required,oneof=sari korsan uber tag"`
}

func GetNearbyDriversRequestToProto(req *GetNearbyDriversRequest) *pb.GetNearbyDriversRequest {
	return &pb.GetNearbyDriversRequest{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		TaxiType:  req.TaxiType,
	}
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

func UpdateDriverRequestToProto(req *UpdateDriverRequest) *pb.UpdateDriverRequest {
	return &pb.UpdateDriverRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Plate:     req.Plate,
		TaxiType:  req.TaxiType,
		CarModel:  req.CarModel,
		CarBrand:  req.CarBrand,
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
	}
}
