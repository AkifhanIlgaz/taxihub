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

func (req *AddDriverRequest) ToProto() *pb.AddDriverRequest {
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

func (req *GetDriversRequest) ToProto() *pb.GetDriversRequest {
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

func (req *GetNearbyDriversRequest) ToProto() *pb.GetNearbyDriversRequest {
	return &pb.GetNearbyDriversRequest{
		Latitude:  req.Latitude,
		Longitude: req.Longitude,
		TaxiType:  req.TaxiType,
	}
}

type UpdateDriverRequest struct {
	FirstName *string   `json:"firstName,omitempty" validate:"omitempty,min=2,max=100"`
	LastName  *string   `json:"lastName,omitempty" validate:"omitempty,min=2,max=100"`
	Plate     *string   `json:"plate,omitempty" validate:"omitempty,min=2,max=100"`
	TaxiType  *string   `json:"taxiType,omitempty" validate:"omitempty,oneof=sari korsan uber tag"`
	CarModel  *string   `json:"carModel,omitempty" validate:"omitempty,min=2,max=100"`
	CarBrand  *string   `json:"carBrand,omitempty" validate:"omitempty,min=2,max=100"`
	Latitude  *float64  `json:"lat,omitempty" validate:"omitempty,latitude"`
	Longitude *float64  `json:"lon,omitempty" validate:"omitempty,longitude"`
	UpdatedAt time.Time `json:"-"`
}

func (req *UpdateDriverRequest) ToProto() *pb.UpdateDriverRequest {
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
