package passenger

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	passengerServiceClient, addressServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		PassengerClient: passengerServiceClient,
		AddressClient:   addressServiceClient,
	}

	routes := r.Group("/passengers")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.ListPassengers)
	routes.GET("/:phone", svc.GetPassengerByPhone)

	return svc
}

func (svc *ServiceClient) ListPassengers(ctx *gin.Context) {
	routes.ListPassengers(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) GetPassengerByPhone(ctx *gin.Context) {
	routes.GetPassengerByPhone(ctx, svc.PassengerClient)
}
