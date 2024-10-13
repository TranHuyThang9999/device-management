package domain

import (
	"context"

	"gorm.io/gorm"
)

type Device struct {
	ID          int64  `json:"id"`          // ID của thiết bị
	DeviceName  string `json:"device_name"` // Tên thiết bị
	Quantity    int    `json:"quantity"`    //
	Description string `json:"description"` // Mô tả thiết bị, có thể là null
	Status      int    `json:"status"`      // Trạng thái, có thể là null
	CreatedAt   int64  `json:"created_at"`  // Thời gian tạo
	UpdatedAt   int64  `json:"updated_at"`  // Thời gian cập nhật
}
type RepositoryDevice interface {
	AddDevice(ctx context.Context, tx *gorm.DB, device *Device) error
	GetListDevice(ctx context.Context, limit, offset int) ([]*Device, error)
	UpdateDevice(ctx context.Context, tx *gorm.DB, device *Device) error
	DeleteDeviceById(ctx context.Context, id int64) error
	CheckDeviceByName(ctx context.Context, name string) bool
}
