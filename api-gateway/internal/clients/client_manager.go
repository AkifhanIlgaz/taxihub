package clients

import (
	"context"
	"fmt"
	"time"

	driverPb "github.com/AkifhanIlgaz/taxihub/common/proto/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type ClientManager struct {
	DriverClient driverPb.DriverServiceClient

	driverConn *grpc.ClientConn
}

func NewClientManager(driverUrl string) (*ClientManager, error) {
	driverConn, err := createGRPCConnection(driverUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to driver service: %w", err)
	}

	return &ClientManager{
		DriverClient: driverPb.NewDriverServiceClient(driverConn),
		driverConn:   driverConn,
	}, nil
}

func createGRPCConnection(address string) (*grpc.ClientConn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(
		ctx,
		address,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024),
		),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                10 * time.Second,
			Timeout:             3 * time.Second,
			PermitWithoutStream: true,
		}),
	)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (cm *ClientManager) Close() {
	if cm.driverConn != nil {
		cm.driverConn.Close()
	}
}
