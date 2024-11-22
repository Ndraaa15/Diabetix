package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
)

type IBMIUsecase interface {
	GetCurrentBMI(ctx context.Context, userID string) (dto.BMIResponse, error)
	CreateBMI(ctx context.Context, req dto.CreateBMIRequest, userID string) (domain.BMI, error)
}

type BMIUsecase struct {
	bmiStore store.IBMIStore
}

func NewBMIUsecase(bmiStore store.IBMIStore) IBMIUsecase {
	return &BMIUsecase{
		bmiStore: bmiStore,
	}
}

func (uc *BMIUsecase) GetCurrentBMI(ctx context.Context, userID string) (dto.BMIResponse, error) {
	currentBmi, err := uc.bmiStore.GetCurrentBMI(ctx, userID)
	if err != nil {
		return dto.BMIResponse{}, err
	}

	weekPreviousBmi, err := uc.bmiStore.GetWeekPreviousBMI(ctx, userID)
	if err != nil {
		return dto.BMIResponse{}, err
	}

	allBmi, err := uc.bmiStore.GetAllBMI(ctx, userID)
	if err != nil {
		return dto.BMIResponse{}, err
	}

	return dto.BMIResponse{
		CurrentBMI:      currentBmi,
		WeekPreviousBMI: weekPreviousBmi,
		AllBMI:          allBmi,
	}, nil
}

func (uc *BMIUsecase) CreateBMI(ctx context.Context, req dto.CreateBMIRequest, userID string) (domain.BMI, error) {
	data := domain.BMI{
		UserID: userID,
		Weight: req.Weight,
		Height: req.Height,
	}

	bmiFactor := req.Weight / ((req.Height / 100) * (req.Height / 100))
	var bmiStatus domain.BMIStatus
	switch {
	case bmiFactor < 18.5:
		bmiStatus = domain.BMIStatusUnderweight
	case bmiFactor >= 18.5 && bmiFactor < 24.9:
		bmiStatus = domain.BMIStatusNormal
	case bmiFactor >= 25 && bmiFactor < 29.9:
		bmiStatus = domain.BMIStatusOverweight
	case bmiFactor >= 30 && bmiFactor <= 34.9:
		bmiStatus = domain.BMIStatusObesityI
	case bmiFactor >= 35 && bmiFactor <= 39.9:
		bmiStatus = domain.BMIStatusObesityII
	case bmiFactor > 39.9:
		bmiStatus = domain.BMIStatusObesityIII
	}

	data.BMI = bmiFactor
	data.Status = bmiStatus

	bmi, err := uc.bmiStore.CreateBMI(ctx, data)
	if err != nil {
		return domain.BMI{}, err
	}

	return bmi, nil
}
