package booking

import (
	"fmt"

	"github.com/lntvan166/e2tech-api-gateway/internal/booking/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	BookingClient pb.BookingServiceClient
}

func InitServiceClient(c *config.Config) pb.BookingServiceClient {
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.BookingSvcUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewBookingServiceClient(cc)
}
