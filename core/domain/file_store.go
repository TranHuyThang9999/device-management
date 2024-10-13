package domain

import (
	"context"

	"gorm.io/gorm"
)

type FileStore struct {
	ID    int64  `json:"id"`
	AnyID int64  `json:"any_id"`
	URL   string `json:"url"`
}

type RepositoryFile interface {
	AddFile(ctx context.Context, tx *gorm.DB, req []*FileStore) error

	DeleteFileById(ctx context.Context, id int64) (*FileStore, error)

	UpdateFile(ctx context.Context, req *FileStore) error

	GetListFiles(ctx context.Context) ([]*FileStore, error)

	GetListFilesByObjectId(ctx context.Context, id int64) ([]*FileStore, error)
}
