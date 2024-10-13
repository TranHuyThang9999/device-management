package repos

import (
	"context"
	"device_management/core/domain"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.RepositoryUser {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) AddUser(ctx context.Context, req *domain.User) error {
	if err := u.db.WithContext(ctx).Create(req).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) DeleteUserById(ctx context.Context, id int64) error {
	if err := u.db.WithContext(ctx).Delete(&domain.User{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetListUser(ctx context.Context, limit int, offset int) ([]*domain.User, error) {
	var users []*domain.User
	if err := u.db.WithContext(ctx).Limit(limit).Offset(offset).Order("updated_at DESC").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) UpdateUserById(ctx context.Context, req *domain.User) error {
	if err := u.db.WithContext(ctx).Model(req).Where("id = ?", req.ID).Updates(req).Error; err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) GetUserByUserName(ctx context.Context, username string) (*domain.User, error) {
	var user *domain.User
	result := u.db.Where("user_name = ?", username).First(&user)
	if result.Error == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return user, result.Error
}
