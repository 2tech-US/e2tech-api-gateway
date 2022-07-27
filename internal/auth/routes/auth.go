package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type registerRequestBody struct {
	Phone       string `json:"phone" binding:"required,min=8,max=15,e164"`
	Password    string `json:"password" binding:"required,min=6"`
	Name        string `json:"name" binding:"required,min=3,max=50"`
	Role        string `json:"role" binding:"required,oneof=admin passenger driver"`
	DateOfBirth string `json:"date_of_birth" binding:"required" time_format:"2006-01-02"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	var req registerRequestBody

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Phone:       req.Phone,
		Password:    req.Password,
		Name:        req.Name,
		Role:        req.Role,
		DateOfBirth: req.DateOfBirth,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

type loginRequestBody struct {
	Phone    string `json:"phone" binding:"required,min=8,max=15,e164"`
	Password string `json:"password" binding:"required,min=6"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	var req loginRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Phone:    req.Phone,
		Password: req.Password,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

type verifyRequestBody struct {
	Phone string `json:"phone" binding:"required,min=8,max=15,e164"`
	Token string `json:"token" binding:"required,min=32"`
}

func Verify(ctx *gin.Context, c pb.AuthServiceClient) {
	var req verifyRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Verify(context.Background(), &pb.VerifyRequest{
		Phone: req.Phone,
		Token: req.Token,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}
