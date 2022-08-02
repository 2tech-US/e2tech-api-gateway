package passenger

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	passengerServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		PassengerClient: passengerServiceClient,
	}

	routes := r.Group("/passengers")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.ListPassengers)
	routes.GET("/:phone", svc.GetPassengerByPhone)
	routes.PUT("/", svc.UpdatePassenger)
	routes.DELETE("/:phone", svc.DeletePassenger)

	addressRoutes := r.Group("/passengers/addresses")
	addressRoutes.Use(a.AuthRequired)
	addressRoutes.POST("/", svc.CreateAddress)
	addressRoutes.GET("/:phone", svc.GetAddress)
	addressRoutes.GET("/location/:phone", svc.GetLocation)
	addressRoutes.PUT("/", svc.UpdateAddress)
	addressRoutes.DELETE("/:phone", svc.DeleteAddress)

	return svc
}

func (svc *ServiceClient) ListPassengers(ctx *gin.Context) {
	routes.ListPassengers(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) GetPassengerByPhone(ctx *gin.Context) {
	routes.GetPassengerByPhone(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) UpdatePassenger(ctx *gin.Context) {
	routes.UpdatePassenger(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) DeletePassenger(ctx *gin.Context) {
	routes.DeletePassenger(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) CreateAddress(ctx *gin.Context) {
	routes.CreateAddress(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) GetAddress(ctx *gin.Context) {
	routes.GetAddress(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) GetLocation(ctx *gin.Context) {
	routes.GetLocation(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) UpdateAddress(ctx *gin.Context) {
	routes.UpdateAddress(ctx, svc.PassengerClient)
}

func (svc *ServiceClient) DeleteAddress(ctx *gin.Context) {
	routes.DeleteAddress(ctx, svc.PassengerClient)
}
