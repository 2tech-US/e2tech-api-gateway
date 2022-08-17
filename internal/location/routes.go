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
	routes.GET("/address", svc.GetAddress)
	routes.POST("/address", svc.CreateAddress)
	routes.PUT("/address", svc.UpdateAddress)
	routes.GET("/address/search", svc.GetAddressList)

	routes.GET("/request/:phone", svc.GetRequest)
	routes.GET("/requests", svc.GetListRequest)
	routes.GET("/requests/:phone", svc.GetListRequestBYPhone)

	routes.POST("/request/", svc.CreateRequest)

	return svc
}

func (svc *ServiceClient) GetAddress(ctx *gin.Context) {
	routes.GetAddress(ctx, svc.LocationClient)
}

func (svc *ServiceClient) GetAddressList(ctx *gin.Context) {
	routes.GetAddressList(ctx, svc.LocationClient)
}

func (svc *ServiceClient) CreateAddress(ctx *gin.Context) {
	routes.CreateAddress(ctx, svc.LocationClient)
}

func (svc *ServiceClient) UpdateAddress(ctx *gin.Context) {
	routes.UpdateAddress(ctx, svc.LocationClient)
}

func (svc *ServiceClient) GetRequest(ctx *gin.Context) {
	routes.GetRequest(ctx, svc.LocationClient)
}

func (svc *ServiceClient) GetListRequest(ctx *gin.Context) {
	routes.GetListRequest(ctx, svc.LocationClient)
}

func (svc *ServiceClient) GetListRequestBYPhone(ctx *gin.Context) {
	routes.UpdateAddress(ctx, svc.LocationClient)
}
func (svc *ServiceClient) CreateRequest(ctx *gin.Context) {
	routes.CreateRequest(ctx, svc.LocationClient)
}
