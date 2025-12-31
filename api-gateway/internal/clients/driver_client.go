package clients

import (
	"context"
	"fmt"
	"time"

	pb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverClient struct {
	client pb.DriverServiceClient
	conn   *grpc.ClientConn
}

func NewDriverClient(address string) (*DriverClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to driver service: %w", err)
	}

	return &DriverClient{
		client: pb.NewDriverServiceClient(conn),
		conn:   conn,
	}, nil
}
