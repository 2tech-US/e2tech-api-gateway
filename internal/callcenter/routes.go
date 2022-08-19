package location

import (
	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/callcenter/routes"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
)

func RegisterRoutes(r *gin.Engine, c *config.Config, authSvc *auth.ServiceClient) *ServiceClient {
	a := auth.InitAuthMiddleware(authSvc)

	callcenterServiceClient := InitServiceClient(c)
	svc := &ServiceClient{
		CallCenterClient: callcenterServiceClient,
	}

	routes := r.Group("/callcenter")
	routes.Use(a.AuthRequired)
	routes.GET("/", svc.GetListEmployee)
	routes.GET("/:id", svc.GetEmployee)
	routes.PUT("/", svc.UpdateEmployee)

	return svc
}

func (svc *ServiceClient) GetEmployee(ctx *gin.Context) {
	routes.GetEmployee(ctx, svc.CallCenterClient)
}

func (svc *ServiceClient) GetListEmployee(ctx *gin.Context) {
	routes.GetListEmployee(ctx, svc.CallCenterClient)
}

func (svc *ServiceClient) CreateEmployee(ctx *gin.Context) {
	routes.CreateEmoloyee(ctx, svc.CallCenterClient)
}

func (svc *ServiceClient) UpdateEmployee(ctx *gin.Context) {
	routes.UpdateEmployee(ctx, svc.CallCenterClient)
}
