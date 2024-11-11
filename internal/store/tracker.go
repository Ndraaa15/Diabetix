package store

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type ITrackerStore interface {
	GetCurrentTracker(ctx context.Context, userID string, date time.Time) (domain.Tracker, error)
	GetPersonalization(ctx context.Context, userID string) (domain.Personalization, error)
	CreateTrackerDetail(ctx context.Context, trackerDetail domain.TrackerDetail) error
	CreateTracker(ctx context.Context, tracker domain.Tracker) (domain.Tracker, error)
	GetAllTracker(ctx context.Context, userID string) ([]domain.Tracker, error)
	GetCurrentReport(ctx context.Context, userID string, date time.Time) (domain.Report, error)
	CreateReport(ctx context.Context, report domain.Report) (domain.Report, error)
	UpdateTracker(ctx context.Context, tracker domain.Tracker) error
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

type TrackerStore struct {
	db *gorm.DB
}

func NewTrackerStore(db *gorm.DB) ITrackerStore {
	return &TrackerStore{
		db: db,
	}
}

func (r *TrackerStore) GetCurrentTracker(ctx context.Context, userID string, date time.Time) (domain.Tracker, error) {
	var tracker domain.Tracker

	err := r.db.Where("user_id = ? AND DATE(created_at) = ?", userID, date.Format("2006-01-02")).Preload("TrackerDetails").First(&tracker).Error
	if err != nil {
		return domain.Tracker{}, err
	}

	return tracker, nil
}

func (r *TrackerStore) GetAllTracker(ctx context.Context, userID string) ([]domain.Tracker, error) {
	var trackers []domain.Tracker

	err := r.db.Where("user_id = ?", userID).Preload("TrackerDetails").Find(&trackers).Error
	if err != nil {
		return nil, err
	}

	return trackers, nil
}

func (r *TrackerStore) GetCurrentReport(ctx context.Context, userID string, date time.Time) (domain.Report, error) {
	var report domain.Report

	err := r.db.Where("user_id = ? AND DATE(created_at) = ?", userID, date.Format("2006-01-02")).Preload("Trackers.TrackerDetails").First(&report).Error
	if err != nil {
		return domain.Report{}, err
	}

	return report, nil
}

func (r *TrackerStore) GetPersonalization(ctx context.Context, userID string) (domain.Personalization, error) {
	var personalization domain.Personalization

	err := r.db.Where("user_id = ?", userID).First(&personalization).Error
	if err != nil {
		return domain.Personalization{}, err
	}

	return personalization, nil
}

func (r *TrackerStore) CreateTracker(ctx context.Context, tracker domain.Tracker) (domain.Tracker, error) {
	err := r.db.Model(&domain.Tracker{}).Create(&tracker).Error
	if err != nil {
		return domain.Tracker{}, err
	}

	return tracker, nil
}

func (r *TrackerStore) UpdateTracker(ctx context.Context, tracker domain.Tracker) error {
	err := r.db.Where("id = ?", tracker.ID).Updates(&tracker).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TrackerStore) CreateTrackerDetail(ctx context.Context, trackerDetail domain.TrackerDetail) error {
	err := r.db.Model(&domain.TrackerDetail{}).Create(&trackerDetail).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *TrackerStore) CreateReport(ctx context.Context, report domain.Report) (domain.Report, error) {
	err := r.db.Model(&domain.Report{}).Create(&report).Error
	if err != nil {
		return domain.Report{}, err
	}

	return report, nil
}

func (r *TrackerStore) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}
