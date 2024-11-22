package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"gorm.io/gorm"
)

type IReportStore interface {
	GetAllReport(ctx context.Context, userID string, filter dto.GetReportsFilter) ([]domain.Report, error)
}

type ReportStore struct {
	db *gorm.DB
}

func NewReportStore(db *gorm.DB) IReportStore {
	return &ReportStore{
		db: db,
	}
}

func (r *ReportStore) GetAllReport(ctx context.Context, userID string, filter dto.GetReportsFilter) ([]domain.Report, error) {
	var reports []domain.Report
	query := r.db.Where("user_id = ?", userID).Preload("UserMissions.Mission").Preload("Trackers.TrackerDetails").Order("created_at desc")

	if !filter.Date.IsZero() {
		query = query.Where("start_date >= ? AND end_date <= ?", filter.Date, filter.Date)
	}

	if err := query.Find(&reports).Error; err != nil {
		return nil, err
	}

	return reports, nil
}
