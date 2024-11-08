package store

import "gorm.io/gorm"

type IBMIStore interface {
}

type BMIStore struct {
	db *gorm.DB
}

func NewBMIStore(db *gorm.DB) IBMIStore {
	return &BMIStore{
		db: db,
	}
}
