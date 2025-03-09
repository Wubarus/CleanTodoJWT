package repo

import "gorm.io/gorm"

// Storage that contains DB
// Storage also may contain other DBs, kv-storages
type Storage struct {
	db *gorm.DB
}

// New Storage instance
func NewStorage(db *gorm.DB) *Storage {
	return &Storage{
		db: db,
	}
}
