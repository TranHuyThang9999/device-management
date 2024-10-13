package controller

import (
	"device_management/core/usecase"

	"github.com/gin-gonic/gin"
)

type ControllerFileStore struct {
	fileCtr *usecase.UseCaseFileStore
	*baseController
}

func NewControllerFileStore(fileCtr *usecase.UseCaseFileStore, base *baseController) *ControllerFileStore {
	return &ControllerFileStore{
		fileCtr:        fileCtr,
		baseController: base,
	}
}

func (u *ControllerFileStore) GetListFilesByObjectId(ctx *gin.Context) {
	id := ctx.Query("id")
	listFile, err := u.fileCtr.GetListFilesByObjectId(ctx, id)
	if err != nil {
		u.baseController.ErrorData(ctx, err)
		return
	}
	u.baseController.Success(ctx, listFile)
}

func (u *ControllerFileStore) DeleteFileById(ctx *gin.Context) {
	id := ctx.Query("id")
	err := u.fileCtr.DeleteFileById(ctx, id)
	if err != nil {
		u.baseController.ErrorData(ctx, err)
		return
	}
	u.baseController.Success(ctx, nil)
}
