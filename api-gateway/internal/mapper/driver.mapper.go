package mapper

import (
	dtoRes "github.com/AkifhanIlgaz/taxihub/api-gateway/internal/dto/response"
	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
)

func ProtoToDriverResponse(driver *pb.Driver) *dtoRes.DriverResponse {
	return &dtoRes.DriverResponse{
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

func ProtoToGetDriversResponse(protoResp *pb.GetDriversResponse) *dtoRes.GetDriversResponse {
	return &dtoRes.GetDriversResponse{
		Drivers:            protoToDriverResponses(protoResp.Drivers),
		PaginationMetadata: protoToPaginationMetadataResponse(protoResp.Meta),
	}
}

func ProtoToGetNearbyDriversResponse(protoResp *pb.GetNearbyDriversResponse) *dtoRes.GetNearbyDriversResponse {
	return &dtoRes.GetNearbyDriversResponse{
		Drivers: protoToNearbyDriversResponse(protoResp.Drivers),
	}
}

func protoToNearbyDriversResponse(drivers []*pb.NearbyDriver) []*dtoRes.NearbyDriverResponse {
	responses := make([]*dtoRes.NearbyDriverResponse, len(drivers))
	for i, driver := range drivers {
		responses[i] = protoToNearbyDriverResponse(driver)
	}
	return responses
}

func protoToNearbyDriverResponse(driver *pb.NearbyDriver) *dtoRes.NearbyDriverResponse {
	return &dtoRes.NearbyDriverResponse{
		FirstName:  driver.FirstName,
		LastName:   driver.LastName,
		Plate:      driver.Plate,
		DistanceKm: driver.DistanceKm,
	}
}

func protoToDriverResponses(drivers []*pb.Driver) []*dtoRes.DriverResponse {
	responses := make([]*dtoRes.DriverResponse, len(drivers))
	for i, driver := range drivers {
		responses[i] = ProtoToDriverResponse(driver)
	}
	return responses
}

func protoToPaginationMetadataResponse(meta *pb.PaginationMetadata) *dtoRes.PaginationMetadataResponse {
	if meta == nil {
		return &dtoRes.PaginationMetadataResponse{}
	}
	return &dtoRes.PaginationMetadataResponse{
		Page:       meta.Page,
		PageSize:   meta.PageSize,
		TotalCount: meta.TotalCount,
		TotalPages: meta.TotalPages,
	}
}
