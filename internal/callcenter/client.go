package location

import (
	"fmt"

	"github.com/lntvan166/e2tech-api-gateway/internal/callcenter/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	CallCenterClient pb.CallCenterServiceClient
}

func InitServiceClient(c *config.Config) pb.CallCenterServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.CallCenterSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCallCenterServiceClient(cc)
}
