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

func (r *TrackerStore) GetPersonalization(ctx context.Context, userID string) (domain.Personalization, error) {
	var personalization domain.Personalization

	err := r.db.Where("user_id = ?", userID).First(&personalization).Error
	if err != nil {
		return domain.Personalization{}, err
	}

	return personalization, nil
}
