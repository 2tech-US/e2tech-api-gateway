package integration

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/booking"
	"github.com/lntvan166/e2tech-api-gateway/internal/callcenter"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/driver"
	"github.com/lntvan166/e2tech-api-gateway/internal/location"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger"
)

var c config.Config
var r *gin.Engine

type response struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
}

func TestMain(m *testing.M) {
	c = config.Config{
		Port:            ":3000",
		AuthSvcUrl:      "localhost:50051",
		PassengerSvcUrl: "localhost:50052",
		DriverSvcUrl:    "localhost:50053",
		BookingSvcUrl:   "localhost:50054",
	}

	gin.SetMode(gin.TestMode)
	r = gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	passenger.RegisterRoutes(r, &c, &authSvc)
	driver.RegisterRoutes(r, &c, &authSvc)
	booking.RegisterRoutes(r, &c, &authSvc)
	location.RegisterRoutes(r, &c, &authSvc)
	callcenter.RegisterRoutes(r, &c, &authSvc)

	os.Exit(m.Run())

}
