package location

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/location/routes"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	locationServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		LocationClient: locationServiceClient,
	}

	routes := r.Group("/callcenter")
	routes.Use(a.AuthRequired)
	routes.GET("/address", svc.GetAddress)
	routes.GET("/address/search", svc.SearchAddress)
	routes.POST("/address", svc.CreateAddress)
	routes.PUT("/address", svc.UpdateAddress)

	routes.GET("/request/phone", svc.GetRecentPhoneCall)
	routes.GET("/request/:id", svc.GetRequest)
	routes.GET("/request", svc.GetListRequest)
	routes.POST("/request", svc.CreateRequest)
	routes.POST("/request/:id", svc.SendRequest)
	routes.PUT("/request/:id", svc.CancelRequest)

	return svc
}

func (svc *ServiceClient) GetAddress(ctx *gin.Context) {
	routes.GetAddress(ctx, svc.LocationClient)
}

func (svc *ServiceClient) SearchAddress(ctx *gin.Context) {
	routes.SearchAddress(ctx, svc.LocationClient)
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

func (svc *ServiceClient) GetRecentPhoneCall(ctx *gin.Context) {
	routes.GetRecentPhoneCall(ctx, svc.LocationClient)
}

func (svc *ServiceClient) CreateRequest(ctx *gin.Context) {
	routes.CreateRequest(ctx, svc.LocationClient)
}

func (svc *ServiceClient) CancelRequest(ctx *gin.Context) {
	routes.CancelRequest(ctx, svc.LocationClient)
}

func (svc *ServiceClient) SendRequest(ctx *gin.Context) {
	routes.SendRequest(ctx, svc.LocationClient)
}
