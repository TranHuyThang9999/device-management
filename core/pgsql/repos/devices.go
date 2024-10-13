package repos

import (
	"context"
	"device_management/core/domain"

	"gorm.io/gorm"
)

type DevicesRepository struct {
	db *gorm.DB
}

func NewDevicesRepository(db *gorm.DB) domain.RepositoryDevice {
	return &DevicesRepository{
		db: db,
	}
}

func (d *DevicesRepository) AddDevice(ctx context.Context, tx *gorm.DB, device *domain.Device) error {
	result := tx.Create(&device)
	return result.Error
}

func (d *DevicesRepository) DeleteDeviceById(ctx context.Context, id int64) error {
	result := d.db.Where("id=?", id).Delete(&domain.Device{})
	return result.Error
}

func (d *DevicesRepository) GetListDevice(ctx context.Context, limit, offset int) ([]*domain.Device, error) {
	var devices = make([]*domain.Device, 0)

	result := d.db.WithContext(ctx).
		Order("updated_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&devices)

	return devices, result.Error
}

func (d *DevicesRepository) UpdateDevice(ctx context.Context, tx *gorm.DB, device *domain.Device) error {
	result := tx.Where("id = ?", device.ID).UpdateColumns(&domain.Device{
		ID:          device.ID,
		DeviceName:  device.DeviceName,
		Quantity:    device.Quantity,
		Description: device.Description,
		Status:      device.Status,
		CreatedAt:   device.CreatedAt,
		UpdatedAt:   device.UpdatedAt,
	})
	return result.Error
}

func (d *DevicesRepository) CheckDeviceByName(ctx context.Context, name string) bool {
	var exists bool
	result := d.db.Model(&domain.Device{}).Select("count(1) > 0").Where("device_name = ?", name).Find(&exists)
	if result.Error != nil || result.RowsAffected == 0 {
		return false
	}
	return exists
}
