package middleware

import (
	"device_management/common/log"
	"device_management/core/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	jwt  *usecase.UseCaseJwt
	user *usecase.UseCaseUser
}

func NewMiddleware(jwt *usecase.UseCaseJwt, user *usecase.UseCaseUser) *Middleware {
	return &Middleware{
		jwt:  jwt,
		user: user,
	}
}

func (m *Middleware) Authenticate() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(401, gin.H{"error": "request does not contain an access token"})
			context.Abort()
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			context.JSON(401, gin.H{"error": "invalid authorization header format 2"})
			context.Abort()
			return
		}

		tokenString := tokenParts[1]

		data, err := m.jwt.Verify(tokenString)
		if err != nil {
			context.JSON(401, gin.H{"error": "invalid authorization"})
			context.Abort()
			return
		}
		user, err := m.user.GetInforUser(context, data.UserName)
		if err != nil {
			context.JSON(401, gin.H{"error": "invalid authorization"})
			context.Abort()
			return
		}
		if data.UpdatedAt != user.UpdatedAt {
			context.JSON(401, gin.H{"error": "invalid authorization"})
			context.Abort()
			return
		}
		log.Infof("data ", data.UpdatedAt, user.UpdatedAt)
		context.Set("userId", data.UserId)
		context.Set("role", data.Role)

		context.Next()
	}
}
