package db

import (
	communication "github.com/AlyssonT/CheckpointBackend/communication/dtos"
	"gorm.io/gorm"
)

func Paginate(pagination *communication.PaginationRequest) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if pagination.Page <= 0 {
			pagination.Page = 1
		}
		if pagination.PageSize <= 0 {
			pagination.PageSize = 10
		}
		if pagination.PageSize > 100 {
			pagination.PageSize = 100
		}

		offset := (pagination.Page - 1) * pagination.PageSize

		return db.Offset(offset).Limit(pagination.PageSize)
	}
}
