package controller

import (
	"device_management/core/entities"
	"device_management/core/usecase"

	"github.com/gin-gonic/gin"
)

type ControllerDevices struct {
	base   *baseController
	device *usecase.UseCaseDevice
}

func NewControllerDevices(base *baseController, device *usecase.UseCaseDevice) *ControllerDevices {
	return &ControllerDevices{
		base:   base,
		device: device,
	}
}

func (h *ControllerDevices) AddDevice(context *gin.Context) {
	var req entities.Devices
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.device.AddDevice(context, &req); err != nil {
		h.base.ErrorData(context, err)
		return
	}

	h.base.Success(context, nil)
}
func (u *ControllerDevices) GetListDevice(ctx *gin.Context) {
	limit := ctx.Query("limit")
	offsert := ctx.Query("offset")
	devices, err := u.device.GetListDevice(ctx, limit, offsert)
	if err != nil {
		u.base.ErrorData(ctx, err)
		return
	}
	u.base.Success(ctx, devices)
}

func (u *ControllerDevices) DeleteDeviceById(ctx *gin.Context) {
	id := ctx.Query("id")
	err := u.device.DeleteDeviceById(ctx, id)
	if err != nil {
		u.base.ErrorData(ctx, err)
		return
	}
	u.base.Success(ctx, nil)
}

func (h *ControllerDevices) UpdatedDeviceById(ctx *gin.Context) {
	var req entities.DeviceReqUpdate
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := h.device.UpdatedDeviceById(ctx, &req); err != nil {
		h.base.ErrorData(ctx, err)
		return
	}

	h.base.Success(ctx, nil)
}
