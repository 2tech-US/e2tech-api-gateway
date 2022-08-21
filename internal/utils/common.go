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

func VerifyRole(ctx *gin.Context, roles []string) error {
	if contains(roles, ctx.GetString("role")) {
		return nil
	}
	return fmt.Errorf("you don't have permission to access this resource")
}

func VerifyAdminPermission(ctx *gin.Context) error {
	if ctx.GetString("role") != ADMIN {
		return fmt.Errorf("only admin can access this resource")
	}
	return nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
