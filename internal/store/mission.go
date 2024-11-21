package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IMissionStore interface {
	GetAllUserMission(ctx context.Context, userID string) ([]domain.UserMission, error)
	GetUserMission(ctx context.Context, userID string, missionID uint64) (domain.UserMission, error)
	UpdateUserMission(ctx context.Context, mission domain.UserMission) error
	GetUserByID(ctx context.Context, userID string) (domain.User, error)
	UpdateUser(ctx context.Context, user domain.User) error
	GetLevelByID(ctx context.Context, levelID uint64) (domain.Level, error)

	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
}

type MissionStore struct {
	db *gorm.DB
}

func NewMissionStore(db *gorm.DB) IMissionStore {
	return &MissionStore{
		db: db,
	}
}

func (r *MissionStore) GetAllUserMission(ctx context.Context, userID string) ([]domain.UserMission, error) {
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

func (r *MissionStore) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	var user domain.User
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Preload("Personalization").Preload("Level").Where("id = ?", userID).First(&user).Error; err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func (r *MissionStore) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *MissionStore) UpdateUser(ctx context.Context, user domain.User) error {
	if err := r.db.Model(domain.User{}).Where("id = ?", user.ID).Updates(&user).Error; err != nil {
		return err
	}

	return nil
}

func (r *MissionStore) GetLevelByID(ctx context.Context, levelID uint64) (domain.Level, error) {
	var level domain.Level
	if err := r.db.WithContext(ctx).Model(&domain.Level{}).Where("id = ?", levelID).First(&level).Error; err != nil {
		return domain.Level{}, err
	}

	return level, nil
}
