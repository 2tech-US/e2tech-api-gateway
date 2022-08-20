package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/booking/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type acceptRequestRequest struct {
	PassengerPhone string `json:"passenger_phone" binding:"required,min=8,max=15"`
	DriverPhone    string `json:"driver_phone" binding:"required"`
}

func AcceptRequest(ctx *gin.Context, c pb.BookingServiceClient) {
	var req acceptRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, req.DriverPhone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.AcceptRequest(context.Background(), &pb.AcceptRequestRequest{
		PassengerPhone: req.PassengerPhone,
		DriverPhone:    req.DriverPhone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type RejectRequestRequest struct {
	PassengerPhone string `json:"passenger_phone" binding:"required"`
	DriverPhone    string `json:"driver_phone" binding:"required"`
}

func RejectRequest(ctx *gin.Context, c pb.BookingServiceClient) {
	var req RejectRequestRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, req.DriverPhone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.RejectRequest(context.Background(), &pb.RejectRequestRequest{
		PassengerPhone: req.PassengerPhone,
		DriverPhone:    req.DriverPhone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
