package repos

import (
	"context"
	"device_management/core/domain"
	"fmt"

	"gorm.io/gorm"
)

type DBHelper struct {
	db *gorm.DB
}

func NewDBHelper(db *gorm.DB) domain.TransactionHelper {
	return &DBHelper{db: db}
}

func (dh *DBHelper) beginTx(ctx context.Context) (*gorm.DB, error) {
	tx := dh.db.WithContext(ctx).Begin()
	if tx.Error != nil {
		return nil, tx.Error
	}
	return tx, nil
}

func (dh *DBHelper) ExecuteInTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx, err := dh.beginTx(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := fn(tx); err != nil {
		tx.Rollback()
		return fmt.Errorf("transaction error: %w", err)
	}

	return tx.Commit().Error
}
