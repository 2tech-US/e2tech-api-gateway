package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/location/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type GetAddressRequestBody struct {
	City     string `json:"city" binding:"required,min=1,max=30"`
	District string `json:"district" binding:"required,min=1,max=30"`
	Ward     string `json:"ward" binding:"required,min=1,max=30"`
	Street   string `json:"street" binding:"required,min=1,max=30"`
	Home     string `json:"home" binding:"required,min=1,max=30"`
}

func GetAddress(ctx *gin.Context, c pb.LocationServiceClient) {
	var b GetAddressRequestBody
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

func CreateAddress(ctx *gin.Context, c pb.LocationServiceClient) {
	var b GetAddressRequestBody
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

type SearchAddressQuery struct {
	Offset        int32  `form:"offset" binding:"required,min=1"`
	Limit         int32  `form:"limit" binding:"required,min=1,max=30"`
	SearchAddress string `form:"search" binding:"required,min=1"`
}

func SearchAddress(ctx *gin.Context, c pb.LocationServiceClient) {
	var q SearchAddressQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := c.SearchAddress(context.Background(), &pb.SearchAddressRequest{
		Offset:        q.Offset,
		Limit:         q.Limit,
		SearchAddress: q.SearchAddress,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type GetLocationRequestBody struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
}

type UpdateAddressRequest struct {
	Address  GetAddressRequestBody  `json:"address" binding:"required"`
	Location GetLocationRequestBody `json:"location" binding:"required"`
}

func UpdateAddress(ctx *gin.Context, c pb.LocationServiceClient) {
	var b UpdateAddressRequest
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	adddress, location := b.Address, b.Location
	res, err := c.UpdateAddress(context.Background(), &pb.UpdateAddressRequest{
		Address: &pb.AddressKey{
			City:     adddress.City,
			District: adddress.District,
			Ward:     adddress.Ward,
			Street:   adddress.Street,
			Home:     adddress.Home,
		},
		Location: &pb.LocationKey{
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func UpdateLocation(ctx *gin.Context, c pb.LocationServiceClient) {
	var location GetLocationRequestBody
	if err := ctx.ShouldBindJSON(&location); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.UpdateLocation(context.Background(), &pb.UpdateLocationRequest{
		Location: &pb.LocationKey{
			Latitude:  location.Latitude,
			Longitude: location.Longitude,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type CreateRequestBody struct {
	Phone           string                `json:"phone" binding:"required"`
	EmployeeId      string                `json:"employeeId" binding:"required"`
	Type            string                `json:"type" binding:"required"`
	PickingAddress  GetAddressRequestBody `json:"picking" binding:"required"`
	ArrivingAddress GetAddressRequestBody `json:"arriving" binding:"required"`
}

func CreateRequest(ctx *gin.Context, c pb.LocationServiceClient) {
	var b CreateRequestBody
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	picking, arriving := b.PickingAddress, b.ArrivingAddress

	res, err := c.CreateRequest(context.Background(), &pb.CreateCallCenterRequest{
		Request: &pb.CallCenterRequestCreation{
			Id:         0,
			Phone:      b.Phone,
			Type:       b.Type,
			EmployeeId: b.EmployeeId,
			PickingAddress: &pb.AddressKey{
				City:     picking.City,
				District: picking.District,
				Ward:     picking.Ward,
				Street:   picking.Street,
				Home:     picking.Home,
			},
			ArrivingAddress: &pb.AddressKey{
				City:     arriving.City,
				District: arriving.District,
				Ward:     arriving.Ward,
				Street:   arriving.Street,
				Home:     arriving.Home,
			},
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type GetRequestBody struct {
	Id int64 `uri:"id" binding:"required"`
}

func GetRequest(ctx *gin.Context, c pb.LocationServiceClient) {
	var b GetRequestBody
	if err := ctx.ShouldBindUri(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetRequest(context.Background(), &pb.GetCallCenterRequest{
		RequestId: b.Id,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type GetListRequestQuery struct {
	Offset int32  `form:"offset" binding:"required,min=1"`
	Limit  int32  `form:"limit" binding:"required,min=1,max=30"`
	Phone  string `form:"phone"`
}

func GetListRequest(ctx *gin.Context, c pb.LocationServiceClient) {
	var q GetListRequestQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := c.GetListRequest(context.Background(), &pb.GetListCallCenterRequest{
		Offset: q.Offset,
		Limit:  q.Limit,
		Phone:  q.Phone,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

func verifyPermission(ctx *gin.Context, phoneBody string) error {
	if ctx.GetString("role") == utils.ADMIN || ctx.GetString("role") == utils.CALLCENTER {
		return nil
	}
	if ctx.GetString("phone") != phoneBody {
		return fmt.Errorf("you don't have permission to access this resource")
	}
	return nil
}
