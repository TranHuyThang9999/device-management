package repos

import (
	"context"
	"device_management/core/domain"

	"gorm.io/gorm"
)

type FileStoreRepository struct {
	db *gorm.DB
}

func NewFileStoreRepository(db *gorm.DB) domain.RepositoryFile {
	return &FileStoreRepository{
		db: db,
	}
}

func (f *FileStoreRepository) AddFile(ctx context.Context, tx *gorm.DB, req []*domain.FileStore) error {
	result := tx.Create(&req)
	return result.Error
}

func (f *FileStoreRepository) DeleteFileById(ctx context.Context, id int64) error {
	result := f.db.Where("id = ?", id).Delete(&domain.FileStore{})
	return result.Error
}

func (f *FileStoreRepository) UpdateFile(ctx context.Context, req *domain.FileStore) error {
	result := f.db.Create(&req)
	return result.Error
}

func (f *FileStoreRepository) GetListFiles(ctx context.Context) ([]*domain.FileStore, error) {
	var files = make([]*domain.FileStore, 0)
	result := f.db.Find(&files)
	return files, result.Error
}

func (f *FileStoreRepository) GetListFilesByObjectId(ctx context.Context, id int64) ([]*domain.FileStore, error) {
	var files = make([]*domain.FileStore, 0)
	result := f.db.Where("any_id = ?", id).Find(&files)
	return files, result.Error
}
