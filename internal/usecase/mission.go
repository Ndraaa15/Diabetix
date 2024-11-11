package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"gorm.io/gorm"
)

type IMissionUsecase interface {
	GetAllMissionUser(ctx context.Context, userID string) ([]domain.UserMission, error)
	UpdateUserMission(ctx context.Context, userID string, missionID uint64, req dto.UpdateUserMissionRequest) error
}

type MissionUsecase struct {
	missionStore store.IMissionStore
}

func NewMissionUsecase(missionStore store.IMissionStore) IMissionUsecase {
	return &MissionUsecase{
		missionStore: missionStore,
	}
}

func (uc *MissionUsecase) GetAllMissionUser(ctx context.Context, userID string) ([]domain.UserMission, error) {
	mission, err := uc.missionStore.GetAllMissionUser(ctx, userID)
	if err != nil {
		return nil, err
	}

	return mission, nil
}

func (uc *MissionUsecase) UpdateUserMission(ctx context.Context, userID string, missionID uint64, req dto.UpdateUserMissionRequest) error {
	userMission, err := uc.missionStore.GetUserMission(ctx, userID, missionID)
	if err != nil {
		return err
	}

	if req.Status == "accepted" {
		userMission.IsDone = true
	}

	user, err := uc.missionStore.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	user.CurrentExp += userMission.Mission.Exp

	if user.CurrentExp >= user.Level.TotalExp {
		// Todo : make sure the level is the next level (and consist in database)
		user.LevelID = user.LevelID + 1
	}

	err = uc.missionStore.WithTransaction(ctx, func(tx *gorm.DB) error {
		if err := uc.missionStore.UpdateUser(ctx, user); err != nil {
			return err
		}

		if err := uc.missionStore.UpdateUserMission(ctx, userMission); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
