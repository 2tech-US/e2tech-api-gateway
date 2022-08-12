package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/location/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type getAddressRequestBody struct {
	City     string `json:"city" binding:required,min=3,max=15`
	District string `json:"district" binding:required,min=3,max=15`
	Ward     string `json:"ward" binding:required,min=3,max=15`
	Street   string `json:"street" binding:required,min=3,max=15`
	Home     string `json:"home" binding:required,min=3,max=15`
}

func GetAddress(ctx *gin.Context, c pb.LocationServiceClient) {
	var b getAddressRequestBody
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetAddress(context.Background(), &pb.GetAddressRequest{
		AddressKey: &pb.AddressKey{
			City:     b.City,
			District: b.District,
			Ward:     b.Ward,
			Street:   b.Street,
			Home:     b.Home,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)

}
func GetAddressList(ctx *gin.Context, c pb.LocationServiceClient) {

}

func CreateAddress(ctx *gin.Context, c pb.LocationServiceClient) {
	var b getAddressRequestBody
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.CreateAddress(context.Background(), &pb.CreateAddressRequest{
		AddressKey: &pb.AddressKey{
			City:     b.City,
			District: b.District,
			Ward:     b.Ward,
			Street:   b.Street,
			Home:     b.Home,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func UpdateAddress(ctx *gin.Context, c pb.LocationServiceClient) {

}

func verifyPermission(ctx *gin.Context, phoneBody string) error {
	if ctx.GetString("role") == utils.ADMIN {
		return nil
	}
	if ctx.GetString("phone") != phoneBody {
		return fmt.Errorf("you don't have permission to access this resource")
	}
	return nil
}
