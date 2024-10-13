package usecase

import (
	"device_management/common/enums"
	"device_management/common/log"
	"device_management/common/utils"
	"device_management/core/domain"
	"device_management/core/entities"
	"device_management/core/errors"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UseCaseDevice struct {
	device domain.RepositoryDevice
	file   domain.RepositoryFile
	trans  domain.TransactionHelper
}

func NewUseCaseDevice(device domain.RepositoryDevice,
	file domain.RepositoryFile,
	trans domain.TransactionHelper) *UseCaseDevice {
	return &UseCaseDevice{
		device: device,
		file:   file,
		trans:  trans,
	}
}

func (u *UseCaseDevice) AddDevice(ctx *gin.Context, req *entities.Devices) errors.Error {

	deviceId := utils.GenerateUniqueKey()
	var listFile = make([]*domain.FileStore, 0)

	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		return errors.NewSystemError("invalid permission: role not found in context")
	}

	if role != enums.RoleAdmin {
		return errors.NewSystemError("invalid permisson")
	}
	existed := u.device.CheckDeviceByName(ctx, req.DeviceName)
	if existed {
		return errors.ErrConflict
	}
	err = u.trans.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {
		err = u.device.AddDevice(ctx, tx, &domain.Device{
			ID:          deviceId,
			DeviceName:  strings.TrimSpace(req.DeviceName),
			Quantity:    req.Quantity,
			Description: req.Description,
			Status:      enums.STATUS_DEVICE_ARE_USING,
			CreatedAt:   utils.GenerateTimestamp(),
			UpdatedAt:   utils.GenerateTimestamp(),
		})
		if err != nil {
			log.Error(err, "error creating device")
			return errors.NewSystemError("error system")
		}
		if len(req.Files) != 0 {
			for _, v := range req.Files {
				listFile = append(listFile, &domain.FileStore{
					ID:    utils.GenerateUniqueKey(),
					AnyID: deviceId,
					URL:   v,
				})
			}
			err = u.file.AddFile(ctx, tx, listFile)
			if err != nil {
				log.Error(err, "error add file")
				return errors.NewSystemError("error system")
			}
		}

		return nil
	})
	if err != nil {
		log.Error(err, "error system")
		return errors.NewSystemError("error system")
	}

	return nil
}

func (u *UseCaseDevice) GetListDevice(ctx *gin.Context, limitStr, offsetStr string) ([]*domain.Device, errors.Error) {
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		return nil, errors.NewSystemError("invalid permission: role not found in context")
	}

	if role != enums.RoleAdmin {
		return nil, errors.NewSystemError("invalid permisson")
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil || offset < 0 {
		offset = 0
	}
	listDevice, err := u.device.GetListDevice(ctx, limit, offset)
	if err != nil {
		log.Error(err, "error system")
		return nil, errors.NewSystemError("error system")
	}
	return listDevice, nil
}

func (u *UseCaseDevice) DeleteDeviceById(ctx *gin.Context, id string) errors.Error {
	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		return errors.NewSystemError("invalid permission: role not found in context")
	}

	if role != enums.RoleAdmin {
		return errors.NewSystemError("invalid permisson")
	}
	deviceId, err := strconv.Atoi(id)
	if err != nil {
		return errors.NewBadRequestWithCode("invalid device id")
	}
	err = u.device.DeleteDeviceById(ctx, int64(deviceId))
	if err != nil {
		return errors.NewSystemError("error system")
	}
	return nil
}

func (u *UseCaseDevice) UpdatedDeviceById(ctx *gin.Context, req *entities.DeviceReqUpdate) errors.Error {

	var listFile = make([]*domain.FileStore, 0)

	role, err := utils.GetRoleFromContext(ctx)
	if err != nil {
		return errors.NewSystemError("invalid permission: role not found in context")
	}

	if role != enums.RoleAdmin {
		return errors.NewSystemError("invalid permisson")
	}

	err = u.trans.ExecuteInTransaction(ctx, func(tx *gorm.DB) error {

		err := u.device.UpdateDevice(ctx, tx, &domain.Device{
			ID:          req.ID,
			DeviceName:  req.DeviceName,
			Quantity:    req.Quantity,
			Description: req.Description,
			UpdatedAt:   utils.GenerateTimestamp(),
		})
		if err != nil {
			log.Error(err, "error updating device")
			return errors.NewSystemError("error system")
		}

		if len(req.Files) > 0 {
			for _, v := range req.Files {
				listFile = append(listFile, &domain.FileStore{
					ID:    utils.GenerateUniqueKey(),
					AnyID: req.ID,
					URL:   v,
				})
			}
			err = u.file.AddFile(ctx, tx, listFile)
			if err != nil {
				log.Error(err, "error add file")
				return errors.NewSystemError("error system")
			}
		}
		return nil
	})
	if err != nil {
		log.Error(err, "error updating device")
		return errors.NewSystemError("error system")
	}

	return nil

}
