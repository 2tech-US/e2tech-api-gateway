package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/driver/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	driverServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		DriverClient: driverServiceClient,
	}

	routes := r.Group("/drivers")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.ListDrivers)
	routes.GET("/:phone", svc.GetDriverByPhone)
	routes.GET("/:phone/location", svc.GetLocation)
	routes.PUT("/", svc.UpdateDriver)
	routes.PUT("/location", svc.UpdateLocation)
	routes.PUT("/status", svc.UpdateDriverStatus)

	return svc
}

func (svc *ServiceClient) ListDrivers(ctx *gin.Context) {
	routes.ListDrivers(ctx, svc.DriverClient)
}

func (svc *ServiceClient) GetDriverByPhone(ctx *gin.Context) {
	routes.GetDriverByPhone(ctx, svc.DriverClient)
}

func (svc *ServiceClient) GetLocation(ctx *gin.Context) {
	routes.GetLocation(ctx, svc.DriverClient)
}

func (svc *ServiceClient) UpdateDriver(ctx *gin.Context) {
	routes.UpdateDriver(ctx, svc.DriverClient)
}

func (svc *ServiceClient) DeleteDriver(ctx *gin.Context) {
	routes.DeleteDriver(ctx, svc.DriverClient)
}

func (svc *ServiceClient) UpdateLocation(ctx *gin.Context) {
	routes.UpdateLocation(ctx, svc.DriverClient)
}

func (svc *ServiceClient) UpdateDriverStatus(ctx *gin.Context) {
	routes.UpdateDriverStatus(ctx, svc.DriverClient)
}
