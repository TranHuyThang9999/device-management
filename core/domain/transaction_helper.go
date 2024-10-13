package domain

import (
	"context"

	"gorm.io/gorm"
)

type TransactionHelper interface {
	ExecuteInTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}
