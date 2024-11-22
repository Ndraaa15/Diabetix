package store

import "gorm.io/gorm"

type IDoctorStore interface {
}

type DoctorStore struct {
	db *gorm.DB
}

func NewDoctorStore(db *gorm.DB) IDoctorStore {
	return &DoctorStore{
		db: db,
	}
}
