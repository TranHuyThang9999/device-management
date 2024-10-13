package usecase

import (
	"context"
	"device_management/common/enums"
	"device_management/common/log"
	"device_management/common/utils"
	"device_management/core/domain"
	"device_management/core/errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UseCaseFileStore struct {
	files domain.RepositoryFile
}

func NewUseCaseFileStore(files domain.RepositoryFile) *UseCaseFileStore {
	return &UseCaseFileStore{
		files: files,
	}
}

func (u *UseCaseFileStore) AddFileByObjectId(ctx context.Context, file []*domain.FileStore) errors.Error {
	return nil
}
func (u *UseCaseFileStore) RemoveFileByObjectId(ctx context.Context) errors.Error {
	return nil
}
func (u *UseCaseFileStore) GetListFilesByObjectId(ctx *gin.Context, id string) ([]*domain.FileStore, errors.Error) {

	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return nil, errors.NewSystemError("error parsing")
	}
	listFile, err := u.files.GetListFilesByObjectId(ctx, idNumber)
	if err != nil {
		return nil, errors.NewSystemError("error system getting")
	}
	return listFile, nil
}

func (u *UseCaseFileStore) DeleteFileById(ctx *gin.Context, id string) errors.Error {
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		return errors.ErrUnAuthorized
	}
	if role != enums.RoleAdmin {
		return errors.ErrUnAuthorized
	}
	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.NewSystemError("error parsing")
	}
	deletedFile, err := u.files.DeleteFileById(ctx, idNumber)
	if err != nil {
		return errors.NewSystemError("error system getting")
	}
	if deletedFile != nil {
		filePath := filepath.Join("publics", filepath.Base(deletedFile.URL))
		if err := os.Remove(filePath); err != nil {
			log.Error(err, "failed to delete the file from the filesystem")
		} else {
			log.Info(fmt.Sprintf("Successfully deleted file: %s", filePath))
		}
	}
	return nil
}
