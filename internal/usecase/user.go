package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
)

type IUserUsecase interface {
	CreatePersonalization(ctx context.Context, req dto.CreatePersonalizationRequest) error
}

type UserUsecase struct {
	userStore store.IUserStore
}

func NewUserUsecase(userStore store.IUserStore) IUserUsecase {
	return &UserUsecase{
		userStore: userStore,
	}
}

func (uc *UserUsecase) CreatePersonalization(ctx context.Context, req dto.CreatePersonalizationRequest) error {
	data := domain.Personalization{
		UserID:         req.UserID,
		Gender:         req.Gender,
		Age:            req.Age,
		FrequenceSport: req.FrequenceSport,
	}

	bmiFactor := req.Weight / ((req.Height / 100) * (req.Height / 100))
	var bmiStatus string
	switch {
	case bmiFactor < 18.5:
		bmiStatus = "Low"
	case bmiFactor >= 18.5 && bmiFactor < 24.9:
		bmiStatus = "Normal"
	case bmiFactor >= 25:
		bmiStatus = "High"
	}

	bmi := domain.BMI{
		UserID: req.UserID,
		Height: req.Height,
		Weight: req.Weight,
		BMI:    bmiFactor,
		Status: bmiStatus,
	}

	if err := uc.userStore.CreatePersonalizationAndBMI(ctx, data, bmi); err != nil {
		return err
	}

	return nil
}
