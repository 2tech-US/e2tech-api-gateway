package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/passenger/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

func ListPassengers(ctx *gin.Context, c pb.PassengerServiceClient) {
	res, err := c.ListPassengers(context.Background(), &pb.ListPassengersRequest{})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func GetPassengerByPhone(ctx *gin.Context, c pb.PassengerServiceClient) {
	phone := ctx.Param("phone")

	res, err := c.GetPassengerByPhone(context.Background(), &pb.GetPassengerByPhoneRequest{
		Phone: phone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	js, err := json.MarshalIndent(res, "", "  ")
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
	}
	fmt.Println(string(js))

	ctx.JSON(http.StatusOK, &res)
}
