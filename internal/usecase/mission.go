package usecase

import (
	"context"
	"errors"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type IMissionUsecase interface {
	GetAllUserMission(ctx context.Context, userID string) ([]domain.UserMission, error)
	UpdateUserMission(ctx context.Context, userID string, missionID uint64) error
}

type MissionUsecase struct {
	missionStore store.IMissionStore
}

func NewMissionUsecase(missionStore store.IMissionStore) IMissionUsecase {
	return &MissionUsecase{
		missionStore: missionStore,
	}
}

func (uc *MissionUsecase) GetAllUserMission(ctx context.Context, userID string) ([]domain.UserMission, error) {
	mission, err := uc.missionStore.GetAllUserMission(ctx, userID)
	if err != nil {
		return nil, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get all user mission").
			WithError(err)
	}

	return mission, nil
}

func (uc *MissionUsecase) UpdateUserMission(ctx context.Context, userID string, missionID uint64) error {
	userMission, err := uc.missionStore.GetUserMission(ctx, userID, missionID)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get user mission").
			WithError(err)
	}

	if userMission.IsDone {
		return errx.New().
			WithCode(iris.StatusBadRequest).
			WithMessage("Mission has been done").
			WithError(errors.New("Mission has been done"))
	}

	userMission.IsDone = true

	user, err := uc.missionStore.GetUserByID(ctx, userID)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get user").
			WithError(err)
	}

	if user.Level.NextLevel == 0 {
		return errx.New().
			WithCode(iris.StatusBadRequest).
			WithMessage("User has reached the maximum level").
			WithError(errors.New("User has reached the maximum level"))
	}

	user.CurrentExp += userMission.Mission.Exp

	for user.CurrentExp >= user.Level.TotalExp {
		level, err := uc.missionStore.GetLevelByID(ctx, user.LevelID)
		if err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to get level").
				WithError(err)
		}

		user.CurrentExp -= user.Level.TotalExp
		user.LevelID = level.NextLevel

		if level.NextLevel == 0 {
			user.CurrentExp = 0
			break
		}
	}

	err = uc.missionStore.WithTransaction(ctx, func(tx *gorm.DB) error {
		if err := uc.missionStore.UpdateUser(ctx, user); err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to update user").
				WithError(err)
		}

		if err := uc.missionStore.UpdateUserMission(ctx, userMission); err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to update user mission").
				WithError(err)
		}

		return nil
	})

	return nil
}
