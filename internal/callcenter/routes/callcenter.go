package routes

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lntvan166/e2tech-api-gateway/internal/callcenter/pb"
	"github.com/lntvan166/e2tech-api-gateway/internal/utils"
)

type GetEmployeeParams struct {
	Phone string `uri:"phone" binding:"required,min=7,max=30"`
}

func GetEmployee(ctx *gin.Context, c pb.CallCenterServiceClient) {
	var p GetEmployeeParams
	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, p.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetEmployee(context.Background(), &pb.GetEmployeeRequest{
		Phone: p.Phone,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type GetListEmployeeQuery struct {
	Offset int32 `form:"offset" binding:"required,min=1"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=10"`
}

func GetListEmployee(ctx *gin.Context, c pb.CallCenterServiceClient) {
	var q GetListEmployeeQuery
	if err := ctx.ShouldBindQuery(&q); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyAdminPermission(ctx); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.GetListEmployee(context.Background(), &pb.GetListEmployeeRequest{
		Offset: q.Offset,
		Limit:  q.Limit,
	})

	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type CreateEmployeeBody struct {
	Phone string `json:"phone" binding:"required,min=7,max=30"`
	Role  string `json:"role" binding:"required"`
}

func CreateEmoloyee(ctx *gin.Context, c pb.CallCenterServiceClient) {
	var b CreateEmployeeBody
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}
	res, err := c.CreateEmployee(context.Background(), &pb.CreateCallCenterEmployeeRequest{
		Phone: b.Phone,
		Role:  b.Role,
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}

type UpdateEmployeeBody struct {
	Name        string `json:"name" binding:"min=3"`
	Role        string `json:"role"`
	UrlImage    string `json:"url_image"`
	DateOfBirth string `json:"date_of_birth"`
}

type updateEmployeeParam struct {
	Phone string `uri:"phone" binding:"required,min=7,max=30"`
}

func UpdateEmployee(ctx *gin.Context, c pb.CallCenterServiceClient) {
	var p updateEmployeeParam
	var b UpdateEmployeeBody
	if err := ctx.ShouldBindJSON(&b); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := ctx.ShouldBindUri(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, utils.ErrorResponse(err))
		return
	}

	if err := utils.VerifyPermission(ctx, p.Phone); err != nil {
		ctx.JSON(http.StatusForbidden, utils.ErrorResponse(err))
		return
	}

	res, err := c.UpdateEmployee(context.Background(), &pb.UpdateCallCenterEmployeeRequest{
		Employee: &pb.CallCenterEmployee{
			Phone:       p.Phone,
			Name:        b.Name,
			DateOfBirth: b.DateOfBirth,
			UrlImage:    "",
			Role:        b.Role,
		},
	})
	if err != nil {
		ctx.JSON(http.StatusBadGateway, utils.ErrorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, &res)
}
