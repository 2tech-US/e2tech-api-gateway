package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/booking/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type location struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type createRequestRequestBody struct {
	Type            string   `json:"type" binding:"required,oneof=call app"`
	Phone           string   `json:"phone" binding:"required"`
	PickUpLocation  location `json:"pick_up_location" binding:"required"`
	DropOffLocation location `json:"drop_off_location" binding:"required"`
}

func CreateRequest(ctx *gin.Context, c pb.BookingServiceClient) {
	var req createRequestRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, req.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.CreateRequest(context.Background(), &pb.CreateRequestRequest{
		Type:  req.Type,
		Phone: req.Phone,
		PickUpLocation: &pb.Location{
			Latitude:  req.PickUpLocation.Latitude,
			Longitude: req.PickUpLocation.Longitude,
		},
		DropOffLocation: &pb.Location{
			Latitude:  req.DropOffLocation.Latitude,
			Longitude: req.DropOffLocation.Longitude,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type closeRequestRequestBody struct {
	Phone string `uri:"phone" binding:"required"`
}

func CloseRequest(ctx *gin.Context, c pb.BookingServiceClient) {
	var req closeRequestRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, req.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.CloseRequest(context.Background(), &pb.CloseRequestRequest{
		Phone: req.Phone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type getResponseRequestBody struct {
	Phone string `uri:"phone" binding:"required"`
}

func GetResponse(ctx *gin.Context, c pb.BookingServiceClient) {
	var req getResponseRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, req.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetResponse(context.Background(), &pb.GetResponseRequest{
		Phone: req.Phone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type listHistoryRequestBody struct {
	Phone string `json:"phone" binding:"required"`
	Role  string `json:"role" binding:"required,oneof=driver passenger"`
}

func ListHistory(ctx *gin.Context, c pb.BookingServiceClient) {
	var req listHistoryRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, req.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.ListHistory(context.Background(), &pb.ListHistoryRequest{
		Phone: req.Phone,
		Role:  req.Role,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
