package location

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/location/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	// a := auth.InitAuthMiddleware(authSvc)

	locationServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		LocationClient: locationServiceClient,
	}

	routes := r.Group("/location")
	// routes.Use(a.AuthRequired)
	// routes.GET("/", svc.listAddress)
	routes.GET("/", svc.GetAddress)
	routes.POST("/", svc.CreateAddress)
	routes.PUT("/", svc.UpdateAddress)

	return svc
}

func (svc *ServiceClient) CreateAddress(ctx *gin.Context) {
	routes.CreateAddress(ctx, svc.LocationClient)
}

func (svc *ServiceClient) GetAddress(ctx *gin.Context) {
	routes.GetAddress(ctx, svc.LocationClient)
}

func (svc *ServiceClient) GetAddressList(ctx *gin.Context) {
	routes.GetAddressList(ctx, svc.LocationClient)
}

func (svc *ServiceClient) UpdateAddress(ctx *gin.Context) {
	routes.UpdateAddress(ctx, svc.LocationClient)
}
