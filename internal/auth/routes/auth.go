package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/auth/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type registerRequestBody struct {
	Phone    string `json:"phone" binding:"required,min=8,max=15"`
	Password string `json:"password" binding:"required,min=8"`
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Role     string `json:"role" binding:"required,oneof=admin passenger driver callcenter"`
}

func Register(ctx *gin.Context, c pb.AuthServiceClient) {
	var req registerRequestBody

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Register(context.Background(), &pb.RegisterRequest{
		Phone:    req.Phone,
		Password: req.Password,
		Name:     req.Name,
		Role:     req.Role,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type loginRequestBody struct {
	Phone       string `json:"phone" binding:"required,min=8,max=15"`
	Password    string `json:"password" binding:"required,min=8"`
	DeviceToken string `json:"device_token" binding:"required,min=1"`
}

type loginRequestUri struct {
	Role string `uri:"role" binding:"required,oneof=admin passenger driver callcenter"`
}

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	var req loginRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	var reqUri loginRequestUri
	if err := ctx.ShouldBindUri(&reqUri); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Phone:       req.Phone,
		Password:    req.Password,
		Role:        reqUri.Role,
		DeviceToken: req.DeviceToken,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type verifyRequestBody struct {
	Phone string `json:"phone" binding:"required,min=8,max=15"`
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
