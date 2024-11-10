package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
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
	mission, err := uc.missionStore.GetUserMission(ctx, userID, missionID)
	if err != nil {
		return err
	}

	if req.Status == "accepted" {
		mission.IsDone = true
	} else {
		return errx.New().WithCode(iris.StatusBadRequest).WithMessage("Status not valid")
	}

	if err := uc.missionStore.UpdateUserMission(ctx, mission); err != nil {
		return err
	}

	return nil
}
