package booking

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/booking/routes"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	bookingServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		BookingClient: bookingServiceClient,
	}

	routes := r.Group("/booking")
	routes.Use(a.AuthRequired)
	routes.POST("/request", svc.CreateRequest)
	routes.POST("/request/:phone/close", svc.CloseRequest)
	routes.GET("/response/:phone", svc.GetResponse)
	routes.POST("request/accept", svc.AcceptRequest)
	routes.POST("request/reject", svc.RejectRequest)

	return svc
}

func (svc *ServiceClient) CreateRequest(ctx *gin.Context) {
	routes.CreateRequest(ctx, svc.BookingClient)
}

func (svc *ServiceClient) GetResponse(ctx *gin.Context) {
	routes.GetResponse(ctx, svc.BookingClient)
}

func (svc *ServiceClient) CloseRequest(ctx *gin.Context) {
	routes.CloseRequest(ctx, svc.BookingClient)
}

func (svc *ServiceClient) AcceptRequest(ctx *gin.Context) {
	routes.AcceptRequest(ctx, svc.BookingClient)
}

func (svc *ServiceClient) RejectRequest(ctx *gin.Context) {
	routes.RejectRequest(ctx, svc.BookingClient)
}
