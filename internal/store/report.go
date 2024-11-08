package store

import "gorm.io/gorm"

type IReportStore interface {
}

type ReportStore struct {
	db *gorm.DB
}

func NewReportStore(db *gorm.DB) IReportStore {
	return &ReportStore{
		db: db,
	}
}
