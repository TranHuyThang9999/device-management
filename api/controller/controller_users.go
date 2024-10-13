package controller

import (
	"device_management/core/entities"
	"device_management/core/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type ControllerUsers struct {
	base *baseController
	user *usecase.UseCaseUser
	jwt  *usecase.UseCaseJwt
}

func NewControllerUser(base *baseController, user *usecase.UseCaseUser, jwt *usecase.UseCaseJwt) *ControllerUsers {
	return &ControllerUsers{
		base: base,
		user: user,
		jwt:  jwt,
	}
}

func (h *ControllerUsers) AddUser(context *gin.Context) {
	var req entities.User
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.user.AddUser(context, &req); err != nil {
		h.base.ErrorData(context, err)
		return
	}

	h.base.Success(context, nil)
}

func (u *ControllerUsers) Login(ctx *gin.Context) {
	username := ctx.Query("username")
	password := ctx.Query("password")
	resp, err := u.user.Login(ctx, username, password)
	if err != nil {
		u.base.ErrorData(ctx, err)
		return
	}
	u.base.Success(ctx, resp)
}

func (u *ControllerUsers) GetListUser(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offsert := ctx.Query("offset")
	users, err := u.user.GetListUser(ctx, limit, offsert)
	if err != nil {
		u.base.ErrorData(ctx, err)
		return
	}
	u.base.Success(ctx, users)
}

func (u *ControllerUsers) CheckToken(context *gin.Context) {
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

	_, err := u.jwt.Verify(tokenString)
	if err != nil {
		context.JSON(401, gin.H{"error": "invalid authorization"})
		context.Abort()
		return
	}
	u.base.Success(context, nil)
}

func (u *ControllerUsers) DeleteUserById(ctx *gin.Context) {
	id := ctx.Query("id")
	err := u.user.DeleteUserById(ctx, id)
	if err != nil {
		u.base.ErrorData(ctx, err)
		return
	}
	u.base.Success(ctx, nil)
}

func (h *ControllerUsers) UpdateUserById(ctx *gin.Context) {

	var req entities.UserRequestUpdate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.user.UpdatedUserById(ctx, &req); err != nil {
		h.base.ErrorData(ctx, err)
		return
	}

	h.base.Success(ctx, nil)
}

func (u *ControllerUsers) ResetPasswordUser(ctx *gin.Context) {
	id := ctx.Query("id")
	err := u.user.ResetPasswordUser(ctx, id)
	if err != nil {
		u.base.ErrorData(ctx, err)
		return
	}
	u.base.Success(ctx, nil)
}
