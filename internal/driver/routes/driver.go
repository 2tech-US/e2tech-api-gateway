package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/driver/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type getDriverByPhoneRequestBody struct {
	Phone string `uri:"phone" binding:"required,min=8,max=15"`
}

func GetDriverByPhone(ctx *gin.Context, c pb.DriverServiceClient) {
	var req getDriverByPhoneRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := verifyPermission(ctx, req.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetDriverByPhone(context.Background(), &pb.GetDriverByPhoneRequest{
		Phone: req.Phone,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type listDriversRequestBody struct {
	Offset int32 `json:"offset" binding:"min=0"`
	Limit  int32 `json:"limit" binding:"required,min=1,max=100"`
}

func ListDrivers(ctx *gin.Context, c pb.DriverServiceClient) {
	var req listDriversRequestBody

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if ctx.GetString("role") != utils.ADMIN {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(fmt.Errorf("only admin can list drivers")))
		return
	}

	rsp, err := c.ListDrivers(context.Background(), &pb.ListDriversRequest{
		Offset: req.Offset,
		Limit:  req.Limit,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &rsp)
}

type updateDriverRequestBody struct {
	ID          int64  `json:"id" binding:"required"`
	Phone       string `json:"phone" binding:"required,min=8,max=15"`
	Name        string `json:"name" binding:"required"`
	DateOfBirth string `json:"date_of_birth" binding:"required"`
}

func UpdateDriver(ctx *gin.Context, c pb.DriverServiceClient) {
	b := updateDriverRequestBody{}

	if err := ctx.BindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := verifyPermission(ctx, b.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.UpdateDriver(context.Background(), &pb.UpdateDriverRequest{
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

type deleteDriverRequestBody struct {
	Phone string `uri:"phone" binding:"required"`
}

func DeleteDriver(ctx *gin.Context, c pb.DriverServiceClient) {
	var req deleteDriverRequestBody
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := verifyPermission(ctx, req.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.DeleteDriver(context.Background(), &pb.DeleteDriverRequest{
		Phone: req.Phone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
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
