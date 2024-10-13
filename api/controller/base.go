package controller

import (
	"device_management/api/resource"
	"device_management/common/log"
	"device_management/core/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type baseController struct {
	validate *validator.Validate
}

func NewBaseController(
	validate *validator.Validate,
) *baseController {
	return &baseController{
		validate: validate,
	}
}
func (b *baseController) validateRequest(request interface{}) error {
	err := b.validate.Struct(request)
	if err != nil {
		for _, errValidate := range err.(validator.ValidationErrors) {
			log.Debugf("query invalid, err:[%s]", errValidate)
		}
		return err
	}
	return nil
}
func (b *baseController) GenRequestId() string {
	return uuid.New().String()
}

func (b *baseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, resource.NewResponseResource(errors.Success, "success", data))
}

func (b *baseController) ErrorData(c *gin.Context, err errors.Error) {
	c.JSON(err.GetHttpCode(), resource.NewResponseErr(err))
}

func (b *baseController) BadRequest(c *gin.Context, message string) {
	b.ErrorData(c, errors.NewCustomHttpError(http.StatusBadRequest, errors.BadRequest, message))
}

func (b *baseController) DefaultBadRequest(c *gin.Context) {
	b.ErrorData(c, errors.ErrBadRequest)
}
