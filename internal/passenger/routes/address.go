package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type createAddressRequestBody struct {
	Phone       string `json:"phone" binding:"required,min=1"`
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

	phone := ctx.GetString("phone")
	if req.Phone != phone {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("you can only create address from your own passenger")))
		return
	}
	passengerID, err := getPassengerIDFromCtx(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.CreateAddress(context.Background(), &pb.CreateAddressRequest{
		PassengerId: passengerID,
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

type getAddressRequestBody struct {
	Phone string `uri:"phone" binding:"required,min=8"`
}

func GetAddress(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req getAddressRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	phone := ctx.GetString("phone")
	if req.Phone != phone {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("you can only get address from your own passenger")))
		return
	}
	passengerID, err := getPassengerIDFromCtx(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetAddress(context.Background(), &pb.GetAddressRequest{
		PassengerId: passengerID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type getLocationRequestBody struct {
	Phone string `uri:"phone" binding:"required,min=1"`
}

func GetLocation(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req getLocationRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	phone := ctx.GetString("phone")
	if req.Phone != phone {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("you can only get location for your own passenger")))
		return
	}
	passengerID, err := getPassengerIDFromCtx(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetLocation(context.Background(), &pb.GetLocationRequest{
		PassengerId: passengerID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type updateAddressRequestBody struct {
	Phone       string `json:"phone" binding:"required,min=1"`
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

	phone := ctx.GetString("phone")
	if req.Phone != phone {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("you can only update address for your own passenger")))
		return
	}
	passengerID, err := getPassengerIDFromCtx(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.UpdateAddress(context.Background(), &pb.UpdateAddressRequest{
		PassengerId: passengerID,
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
	Phone string `uri:"phone" binding:"required,min=1"`
}

func DeleteAddress(ctx *gin.Context, c pb.PassengerServiceClient) {
	var req deleteAddressRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	phone := ctx.GetString("phone")
	if req.Phone != phone {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(fmt.Errorf("you can only delete address for your own passenger")))
		return
	}
	passengerID, err := getPassengerIDFromCtx(ctx, c)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.DeleteAddress(context.Background(), &pb.DeleteAddressRequest{
		PassengerId: passengerID,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func getPassengerIDFromCtx(ctx *gin.Context, c pb.PassengerServiceClient) (int64, error) {
	phone := ctx.GetString("phone")
	passengerRsp, err := c.GetPassengerByPhone(ctx, &pb.GetPassengerByPhoneRequest{
		Phone: phone,
	})
	if err != nil {
		return 0, err
	}

	if passengerRsp.Passenger == nil {
		return 0, fmt.Errorf("passenger not found")
	}

	return passengerRsp.Passenger.Id, nil
}
