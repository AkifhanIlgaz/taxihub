package main

import (
	"fmt"
	"log"
	"net"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/config"
	grpcServer "github.com/AkifhanIlgaz/taxihub/driver-service/internal/grpc"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/repositories"
	"github.com/AkifhanIlgaz/taxihub/driver-service/internal/services"
	"github.com/AkifhanIlgaz/taxihub/driver-service/pkg/database"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mongodb, err := database.ConnectMongo(config.Mongo)
	if err != nil {
		log.Fatal(err)
	}

	err = database.SeedDrivers(mongodb)
	if err != nil {
		log.Fatal(err)
	}

	driverRepo := repositories.NewDriverRepository(mongodb)
	driverService := services.NewDriverService(driverRepo)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcSrv := grpc.NewServer()
	pb.RegisterDriverServiceServer(grpcSrv, grpcServer.NewDriverServer(driverService))

	reflection.Register(grpcSrv)

	log.Printf("Driver Service gRPC server listening on :%d", config.Port)
	if err := grpcSrv.Serve(lis); err != nil {
		log.Fatalf("failed to serve:  %v", err)
	}
}
