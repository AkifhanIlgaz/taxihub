package grpc

import (
	"context"
	"errors"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/mappers"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/services"
	"go.mongodb.org/mongo-driver/v2/mongo"
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
		if mongo.IsDuplicateKeyError(err) {
			return nil, status.Error(codes.AlreadyExists, "driver with plate already exists")
		}
		return nil, status.Errorf(codes.Internal, "failed to add driver: %v", err)
	}

	return &pb.AddDriverResponse{
		Id: id.Hex(),
	}, nil
}

func (s *DriverServer) UpdateDriver(ctx context.Context, req *pb.UpdateDriverRequest) (*pb.UpdateDriverResponse, error) {
	err := s.service.UpdateDriver(ctx, req.Id, req)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, status.Error(codes.NotFound, "driver not found")
		}
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
