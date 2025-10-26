package scopes

import "gorm.io/gorm"

func Paginate(db *gorm.DB, page, limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset((page - 1) * limit).Limit(limit + 1)
	}
}
