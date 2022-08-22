package driver

import (
	"fmt"

	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/driver/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type DriverServiceClient struct {
	DriverClient pb.DriverServiceClient
}

func InitDiverServiceClient(c *config.Config) pb.DriverServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.DriverSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewDriverServiceClient(cc)
}
