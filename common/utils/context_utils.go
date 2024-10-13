package utils

import (
	"device_management/common/log"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRoleFromContext(c *gin.Context) (int, error) {
	role, exists := c.Get("role")
	if !exists {
		log.Info("role not found")
		return 0, fmt.Errorf("role not found in context")
	}
	return role.(int), nil
}
