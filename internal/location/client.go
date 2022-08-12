package location

import (
	"fmt"

	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/location/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	LocationClient pb.LocationServiceClient
}

func InitServiceClient(c *config.Config) pb.LocationServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.LocationSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewLocationServiceClient(cc)
}
