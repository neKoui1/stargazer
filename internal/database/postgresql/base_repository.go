package postgresql

import (
	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func NewBaseRepository() *BaseRepository {
	return &BaseRepository{
		db: GetPostgre(),
	}
}

// 通用的基础操作，不涉及具体表结构
func (r *BaseRepository) GetDB() *gorm.DB {
	return r.db
}

// 通用的事务操作
func (r *BaseRepository) Transaction(fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

// 通用的Raw SQL执行
func (r *BaseRepository) Raw(query string, args ...interface{}) *gorm.DB {
	return r.db.Raw(query, args...)
}

// 通用的计数查询
func (r *BaseRepository) CountRaw(query string, args ...interface{}) (int64, error) {
	var count int64
	err := r.db.Raw(query, args...).Scan(&count).Error
	return count, err
}
