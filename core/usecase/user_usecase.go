package usecase

import (
	"context"
	"device_management/common/enums"
	"device_management/common/log"
	"device_management/common/utils"
	"device_management/core/domain"
	"device_management/core/entities"
	"device_management/core/errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type UseCaseUser struct {
	user domain.RepositoryUser
	jwt  *UseCaseJwt
}

func NewUseCaseUser(user domain.RepositoryUser, jwt *UseCaseJwt) *UseCaseUser {
	return &UseCaseUser{
		user: user,
		jwt:  jwt,
	}
}

func (u *UseCaseUser) AddUser(ctx *gin.Context, req *entities.User) errors.Error {

	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		return errors.ErrUnAuthorized
	}

	if role != enums.RoleAdmin {
		return errors.NewSystemError("invalid permisson")
	}
	user, err := u.user.GetUserByUserName(ctx, req.UserName)
	if err != nil {
		log.Error(err, "error system")
		return errors.NewSystemError("error system")
	}
	if user != nil {
		return errors.ErrConflict
	}

	err = u.user.AddUser(ctx, &domain.User{
		ID:         utils.GenerateUniqueKey(),
		UserName:   strings.TrimSpace(req.UserName),
		Password:   utils.GenPassWord(),
		Avatar:     req.Avatar,
		Age:        req.Age,
		Role:       enums.RoleUser,
		Department: req.Department,
		CreatedAt:  utils.GenerateTimestamp(),
		UpdatedAt:  utils.GenerateTimestamp(),
	})
	if err != nil {
		log.Error(err, "error system")
		return errors.NewSystemError("error system")
	}
	return nil
}

func (u *UseCaseUser) GetInforUser(ctx context.Context, name string) (*domain.User, errors.Error) {
	user, err := u.user.GetUserByUserName(ctx, name)
	if err != nil {
		log.Error(err, "error system")
		return nil, errors.NewSystemError("error system")
	}
	return user, nil
}

func (u *UseCaseUser) Login(ctx context.Context, userName, passWord string) (*entities.UseCaseJwt, errors.Error) {

	user, err := u.user.GetUserByUserName(ctx, userName)
	if err != nil {
		log.Error(err, "error system")
		return nil, errors.NewSystemError("user or password invaid")
	}
	if user == nil {
		log.Error(err, "error system")
		return nil, errors.ErrConflict
	}

	if user.Password != passWord {
		return nil, errors.ErrConflict
	}

	token, err := u.jwt.GenToken(user.ID, user.Role, user.UpdatedAt, user.UserName)
	if err != nil {
		log.Error(err, "error system")
		return nil, errors.NewSystemError("user or password invaid")
	}
	return &entities.UseCaseJwt{
		Token:     token.Token,
		UserName:  userName,
		UserId:    user.ID,
		UpdatedAt: utils.GenerateTimestamp(),
	}, nil
}

func (u *UseCaseUser) GetListUser(ctx *gin.Context, limitStr, offsetStr string) ([]*entities.GetUsers, errors.Error) {
	var users = make([]*entities.GetUsers, 0)
	role, err := utils.GetRoleFromContext(ctx)

	if err != nil {
		log.Error(err, "error system")
		return nil, errors.NewSystemError("invalid permission")
	}

	if role != enums.RoleAdmin {
		return nil, errors.NewSystemError("invalid permission")
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}

	listUser, err := u.user.GetListUser(ctx, limit, offset)
	if err != nil {
		log.Error(err, "error system")
		return nil, errors.NewSystemError("error system")
	}
	for _, v := range listUser {
		if v.Role == enums.RoleUser {
			users = append(users, &entities.GetUsers{
				ID:         v.ID,
				UserName:   v.UserName,
				Avatar:     v.Avatar,
				Age:        v.Age,
				Department: v.Department,
				Password:   v.Password,
				UpdatedAt:  v.CreatedAt,
			})
		}
	}
	return users, nil
}

func (u *UseCaseUser) DeleteUserById(ctx *gin.Context, id string) errors.Error {

	role, err := utils.GetRoleFromContext(ctx)

	if err != nil {
		return errors.ErrUnAuthorized
	}

	if role != enums.RoleAdmin {
		return errors.ErrUnAuthorized
	}
	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.ErrBadRequest
	}
	err = u.user.DeleteUserById(ctx, idNumber)
	if err != nil {
		return errors.ErrSystem
	}
	return nil
}

func (u *UseCaseUser) UpdatedUserById(ctx *gin.Context, req *entities.UserRequestUpdate) errors.Error {
	role, err := utils.GetRoleFromContext(ctx)

	if err != nil {
		return errors.ErrUnAuthorized
	}

	if role != enums.RoleAdmin {
		return errors.ErrUnAuthorized
	}

	err = u.user.UpdateUserById(ctx, &domain.User{
		ID:         req.ID,
		UserName:   req.UserName,
		Avatar:     req.Avatar,
		Age:        req.Age,
		Department: req.Department,
		UpdatedAt:  utils.GenerateTimestamp(),
	})
	if err != nil {
		return errors.ErrSystem
	}
	return nil
}

func (u *UseCaseUser) ResetPasswordUser(ctx *gin.Context, idUser string) errors.Error {
	role, err := utils.GetRoleFromContext(ctx)

	if err != nil {
		return errors.ErrUnAuthorized
	}

	if role != enums.RoleAdmin {
		return errors.ErrUnAuthorized
	}
	idUserSf, err := strconv.ParseInt(idUser, 10, 64)
	if err != nil {
		return errors.NewBadRequestWithCode("Invalid user ID")
	}
	err = u.user.UpdateUserById(ctx, &domain.User{
		ID:        idUserSf,
		Password:  utils.GenPassWord(),
		UpdatedAt: utils.GenerateTimestamp(),
	})
	if err != nil {
		return errors.ErrSystem
	}
	return nil
}
