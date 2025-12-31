package grpc

import (
	"context"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/models"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type DriverServer struct {
	pb.UnimplementedDriverServiceServer
	service *services.DriverService
}

func NewDriverServer(service *services.DriverService) *DriverServer {
	return &DriverServer{
		service: service,
	}
}

func (s *DriverServer) AddDriver(ctx context.Context, req *pb.AddDriverRequest) (*pb.AddDriverResponse, error) {
	id, err := s.service.AddDriver(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add driver: %v", err)
	}

	return &pb.AddDriverResponse{
		Id: id.Hex(),
	}, nil
}

func (s *DriverServer) UpdateDriver(ctx context.Context, req *pb.UpdateDriverRequest) (*pb.UpdateDriverResponse, error) {
	err := s.service.UpdateDriver(req.Id, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update driver: %v", err)
	}

	return &pb.UpdateDriverResponse{
		Message: "Driver updated successfully",
	}, nil
}

func (s *DriverServer) GetDrivers(ctx context.Context, req *pb.GetDriversRequest) (*pb.GetDriversResponse, error) {
	res, err := s.service.GetDrivers(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get drivers: %v", err)
	}

	return res, nil
}

func (s *DriverServer) GetNearbyDrivers(ctx context.Context, req *pb.GetNearbyDriversRequest) (*pb.GetNearbyDriversResponse, error) {
	res, err := s.service.GetNearbyDrivers(req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get nearby drivers: %v", err)
	}

	return res, nil
}

func driverToProto(d models.Driver) *pb.Driver {
	return &pb.Driver{
		Id:        d.Id.Hex(),
		FirstName: d.FirstName,
		LastName:  d.LastName,
		Plate:     d.Plate,
		TaxiType:  d.TaxiType,
		CarBrand:  d.CarBrand,
		CarModel:  d.CarModel,
		Latitude:  d.Latitude,
		Longitude: d.Longitude,
		CreatedAt: timestamppb.New(d.CreatedAt),
		UpdatedAt: timestamppb.New(d.UpdatedAt),
	}
}
