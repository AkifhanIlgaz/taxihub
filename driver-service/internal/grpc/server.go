package grpc

import (
	"context"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/mappers"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/services"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	id, err := s.service.AddDriver(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to add driver: %v", err)
	}

	return &pb.AddDriverResponse{
		Id: id.Hex(),
	}, nil
}

func (s *DriverServer) UpdateDriver(ctx context.Context, req *pb.UpdateDriverRequest) (*pb.UpdateDriverResponse, error) {
	err := s.service.UpdateDriver(ctx, req.Id, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to update driver: %v", err)
	}

	return &pb.UpdateDriverResponse{
		Message: "Driver updated successfully",
	}, nil
}

func (s *DriverServer) GetDrivers(ctx context.Context, req *pb.GetDriversRequest) (*pb.GetDriversResponse, error) {
	drivers, metadata, err := s.service.GetDrivers(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get drivers: %v", err)
	}

	res := &pb.GetDriversResponse{
		Drivers: mappers.DriversToProto(drivers),
		Meta:    mappers.PaginationMetadataToProto(metadata),
	}

	return res, nil
}

func (s *DriverServer) GetNearbyDrivers(ctx context.Context, req *pb.GetNearbyDriversRequest) (*pb.GetNearbyDriversResponse, error) {
	drivers, err := s.service.GetNearbyDrivers(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get nearby drivers: %v", err)
	}

	res := &pb.GetNearbyDriversResponse{
		Drivers: mappers.NearbyDriversToProto(drivers),
	}

	return res, nil
}
