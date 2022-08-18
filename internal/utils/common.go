package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func VerifyPermission(ctx *gin.Context, phoneBody string) error {
	if ctx.GetString("role") == ADMIN {
		return nil
	}
	if ctx.GetString("phone") != phoneBody {
		return fmt.Errorf("you don't have permission to access this resource")
	}
	return nil
}

func VerifyAdminPermission(ctx *gin.Context) error {
	if ctx.GetString("role") != ADMIN {
		return fmt.Errorf("only admin can access this resource")
	}
	return nil
}
