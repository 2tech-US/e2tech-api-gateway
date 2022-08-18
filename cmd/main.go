package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth"
	"github.com/lntvan166/e2tech-api-gateway/internal/config"
	"github.com/lntvan166/e2tech-api-gateway/internal/driver"
	"github.com/lntvan166/e2tech-api-gateway/internal/location"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()
	r.Use(CORSMiddleware())

	authSvc := *auth.RegisterRoutes(r, &c)
	passenger.RegisterRoutes(r, &c, &authSvc)
	driver.RegisterRoutes(r, &c, &authSvc)
	location.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
