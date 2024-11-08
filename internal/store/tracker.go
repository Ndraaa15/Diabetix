package store

import (
	"gorm.io/gorm"
)

type ITrackerStore interface {
}

type TrackerStore struct {
	db *gorm.DB
}

func NewTrackerStore(db *gorm.DB) ITrackerStore {
	return &AuthStore{
		db: db,
	}
}
