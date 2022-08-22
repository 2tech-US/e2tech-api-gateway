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
	routes.GET("/request/:phone", svc.GetRequest)
	routes.GET("/history", svc.ListHistory)

	routes.GET("/history/revenue/", svc.GetRevenue)

	routes.POST("/request", svc.CreateRequest)
	routes.POST("/request/:phone/close", svc.CloseRequest)
	routes.GET("/response/:phone", svc.GetResponse)

	routes.POST("request/accept", svc.AcceptRequest)
	routes.POST("request/reject", svc.RejectRequest)
	routes.POST("request/complete", svc.Complete)

	return svc
}

func (svc *ServiceClient) CreateRequest(ctx *gin.Context) {
	routes.CreateRequest(ctx, svc.BookingClient)
}

func (svc *ServiceClient) GetRequest(ctx *gin.Context) {
	routes.GetRequest(ctx, svc.BookingClient)
}

func (svc *ServiceClient) ListHistory(ctx *gin.Context) {
	routes.ListHistory(ctx, svc.BookingClient)
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

func (svc *ServiceClient) Complete(ctx *gin.Context) {
	routes.CompleteTrip(ctx, svc.BookingClient)
}

func (svc *ServiceClient) GetRevenue(ctx *gin.Context) {
	routes.GetRevenue(ctx, svc.BookingClient)
}
