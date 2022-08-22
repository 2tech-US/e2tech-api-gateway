package driver

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/driver/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *DriverServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	driverServiceClient := InitDiverServiceClient(c)
	svc := &DriverServiceClient{
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

func (svc *DriverServiceClient) ListDrivers(ctx *gin.Context) {
	routes.ListDrivers(ctx, svc.DriverClient)
}

func (svc *DriverServiceClient) GetDriverByPhone(ctx *gin.Context) {
	routes.GetDriverByPhone(ctx, svc.DriverClient)
}

func (svc *DriverServiceClient) GetLocation(ctx *gin.Context) {
	routes.GetLocation(ctx, svc.DriverClient)
}

func (svc *DriverServiceClient) UpdateDriver(ctx *gin.Context) {
	routes.UpdateDriver(ctx, svc.DriverClient)
}

func (svc *DriverServiceClient) DeleteDriver(ctx *gin.Context) {
	routes.DeleteDriver(ctx, svc.DriverClient)
}

func (svc *DriverServiceClient) UpdateLocation(ctx *gin.Context) {
	routes.UpdateLocation(ctx, svc.DriverClient)
}

func (svc *DriverServiceClient) UpdateDriverStatus(ctx *gin.Context) {
	routes.UpdateDriverStatus(ctx, svc.DriverClient)
}
