package booking

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/booking/routes"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	passengerServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		PassengerClient: passengerServiceClient,
	}

	routes := r.Group("/booking")
	routes.Use(a.AuthRequired)
	routes.POST("/request", svc.CreateRequest)

	return svc
}

func (svc *ServiceClient) CreateRequest(ctx *gin.Context) {
	routes.CreateRequest(ctx, svc.PassengerClient)
}
