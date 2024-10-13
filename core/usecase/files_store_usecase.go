package usecase

import (
	"context"
	"device_management/common/enums"
	"device_management/common/log"
	savefiles "device_management/common/save_files"
	"device_management/common/utils"
	"device_management/core/domain"
	"device_management/core/errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

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
		return errors.NewSystemError("invalid permisson")
	}
	idNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return errors.NewSystemError("error parsing")
	}
	err = u.files.DeleteFileById(ctx, idNumber)
	if err != nil {
		return errors.NewSystemError("error system getting")
	}
	return nil
}

func (u *UseCaseFileStore) TriggerClearFiles(ctx context.Context) errors.Error {
	files := savefiles.ListFilesInDirectory("publics")

	listFilesFromDb, err := u.files.GetListFiles(ctx)
	if err != nil {
		log.Error(err, "error getting files from database")
		return errors.ErrSystem
	}

	dbFileMap := make(map[string]struct{})
	for _, v := range listFilesFromDb {
		dbFileMap[filepath.Base(v.URL)] = struct{}{}
	}

	for _, file := range files {
		fileName := filepath.Base(file)
		if _, exists := dbFileMap[fileName]; !exists {
			filePath := filepath.Join("publics", fileName)
			if err := os.Remove(filePath); err != nil {
				log.Error(err, fmt.Sprintf("failed to delete file: %s", filePath))
				continue
			}
			log.Infof("Deleted file: %s", filePath)
		}
	}

	return nil
}

func (u *UseCaseFileStore) AutoCleanUp(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := u.TriggerClearFiles(ctx); err != nil {
				log.Error(err, "error occurred during file cleanup")
			}
		}
	}
}
