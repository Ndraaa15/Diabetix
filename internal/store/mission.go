package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IMissionStore interface {
	GetAllMissionUser(ctx context.Context, userID string) ([]domain.UserMission, error)
	GetUserMission(ctx context.Context, userID string, missionID uint64) (domain.UserMission, error)
	UpdateUserMission(ctx context.Context, mission domain.UserMission) error
}

type MissionStore struct {
	db *gorm.DB
}

func NewMissionStore(db *gorm.DB) IMissionStore {
	return &MissionStore{
		db: db,
	}
}

func (r *MissionStore) GetAllMissionUser(ctx context.Context, userID string) ([]domain.UserMission, error) {
	var mission []domain.UserMission
	query := r.db.WithContext(ctx).Preload("Mission").Model(&domain.UserMission{}).Where("user_id = ?", userID).Where("is_active = ?", true).Find(&mission)
	if err := query.Error; err != nil {
		return nil, err
	}

	return mission, nil
}

func (r *MissionStore) GetUserMission(ctx context.Context, userID string, missionID uint64) (domain.UserMission, error) {
	var mission domain.UserMission
	query := r.db.WithContext(ctx).Preload("Mission").Model(&domain.UserMission{}).Where("user_id = ?", userID).Where("mission_id = ?", missionID).First(&mission)
	if err := query.Error; err != nil {
		return domain.UserMission{}, err
	}

	return mission, nil
}

func (r *MissionStore) UpdateUserMission(ctx context.Context, mission domain.UserMission) error {
	if err := r.db.WithContext(ctx).Model(&domain.UserMission{}).Where("user_id = ?", mission.UserID).Where("mission_id = ?", mission.MissionID).Updates(&mission).Error; err != nil {
		return err
	}

	return nil
}
