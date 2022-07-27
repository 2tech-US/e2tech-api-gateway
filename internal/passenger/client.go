package passenger

import (
	"fmt"

	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	PassengerClient pb.PassengerServiceClient
}

func InitServiceClient(c *config.Config) pb.PassengerServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.PassengerSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewPassengerServiceClient(cc)
}
