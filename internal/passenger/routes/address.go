package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type createAddressRequestBody struct {
	PassengerID int64  `json:"passenger_id" binding:"required,min=1"`
	Detail      string `json:"detail" binding:"required"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street" binding:"required"`
	Ward        string `json:"ward" binding:"required"`
	District    string `json:"district" binding:"required"`
	City        string `json:"city" binding:"required"`
}

func CreateAddress(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req createAddressRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.CreateAddress(context.Background(), &pb.CreateAddressRequest{
		PassengerId: req.PassengerID,
		Detail:      req.Detail,
		HouseNumber: req.HouseNumber,
		Street:      req.Street,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusCreated, &res)
}

type getAddressRequestBody struct {
	PassengerID int64 `uri:"passenger_id" binding:"required,min=1"`
}

func GetAddress(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req getAddressRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetAddress(context.Background(), &pb.GetAddressRequest{
		PassengerId: req.PassengerID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type getLocationRequestBody struct {
	PassengerID int64 `uri:"passenger_id" binding:"required,min=1"`
}

func GetLocation(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req getLocationRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetLocation(context.Background(), &pb.GetLocationRequest{
		PassengerId: req.PassengerID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type updateAddressRequestBody struct {
	ID          int64  `json:"id" binding:"required,min=1"`
	Detail      string `json:"detail" binding:"required"`
	HouseNumber string `json:"house_number"`
	Street      string `json:"street" binding:"required"`
	Ward        string `json:"ward" binding:"required"`
	District    string `json:"district" binding:"required"`
	City        string `json:"city" binding:"required"`
}

func UpdateAddress(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req updateAddressRequestBody
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.UpdateAddress(context.Background(), &pb.UpdateAddressRequest{
		Id:          req.ID,
		Detail:      req.Detail,
		HouseNumber: req.HouseNumber,
		Street:      req.Street,
		Ward:        req.Ward,
		District:    req.District,
		City:        req.City,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type deleteAddressRequestBody struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func DeleteAddress(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req deleteAddressRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.DeleteAddress(context.Background(), &pb.DeleteAddressRequest{
		Id: req.ID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
