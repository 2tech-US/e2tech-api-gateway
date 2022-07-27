package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

// type CreatePassengerRequestBody struct {
// 	Phone       string `json:"phone"`
// 	Password    string `json:"password"`
// 	Name        string `json:"name"`
// 	DateOfBirth string `json:"date_of_birth"`
// }

// func CreatePassenger(ctx *gin.Context, c pb.PassengerServiceClient) {
// 	b := CreatePassengerRequestBody{}

// 	if err := ctx.BindJSON(&b); err != nil {
// 		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
// 		return
// 	}

// 	res, err := c.CreatePassenger(context.Background(), &pb.CreatePassengerRequest{
// 		Phone:       b.Phone,
// 		Password:    b.Password,
// 		Name:        b.Name,
// 		DateOfBirth: b.DateOfBirth,
// 	})

// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, &res)
// }

type getPassengerByPhoneRequestBody struct {
	Phone string `uri:"phone" binding:"required,min=8,max=15"`
}

func GetPassengerByPhone(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req getPassengerByPhoneRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetPassengerByPhone(context.Background(), &pb.GetPassengerByPhoneRequest{
		Phone: req.Phone,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type listPassengersRequestBody struct {
	Offset int32 `json:"offset" binding:"required,min=0"`
	Limit  int32 `json:"limit" binding:"required,min=1,max=100"`
}

func ListPassengers(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req listPassengersRequestBody

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	rsp, err := c.ListPassengers(context.Background(), &pb.ListPassengersRequest{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &rsp)
}

type updatePassengerRequestBody struct {
	ID          int64  `json:"id" binding:"required"`
	Phone       string `json:"phone" binding:"required,min=8,max=15"`
	Name        string `json:"name" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

func UpdatePassenger(ctx *gin.Context, c pb.PassengerServiceClient) {
	b := updatePassengerRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.UpdatePassenger(context.Background(), &pb.UpdatePassengerRequest{
		Id:          b.ID,
		Phone:       b.Phone,
		Name:        b.Name,
		DateOfBirth: b.DateOfBirth,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type deletePassengerRequestBody struct {
	Phone string `uri:"id" binding:"required"`
}

func DeletePassenger(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req deletePassengerRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.DeletePassenger(context.Background(), &pb.DeletePassengerRequest{
		Phone: req.Phone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
