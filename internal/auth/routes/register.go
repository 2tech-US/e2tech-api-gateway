package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type RegisterRequestBody struct {
	Phone       string `json:"phone"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	DateOfBirth string `json:"date_of_birth"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	b := RegisterRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Phone:       b.Phone,
		Password:    b.Password,
		Name:        b.Name,
		DateOfBirth: b.DateOfBirth,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
