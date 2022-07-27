package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type LoginRequestBody struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	b := LoginRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Phone:    b.Phone,
		Password: b.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
