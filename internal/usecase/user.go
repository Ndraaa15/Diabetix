package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"gorm.io/gorm"
)

type IUserUsecase interface {
	CreatePersonalization(ctx context.Context, req dto.CreatePersonalizationRequest) error
	GetProfile(ctx context.Context, userID string) (domain.User, error)
	GetDataForHomePage(ctx context.Context, userID string) (dto.HomePageResponse, error)
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
	personalization := domain.Personalization{
		UserID:         req.UserID,
		Gender:         req.Gender,
		Age:            req.Age,
		FrequenceSport: req.FrequenceSport,
		MaxGlucose:     50.0,
	}

	bmiFactor := req.Weight / ((req.Height / 100) * (req.Height / 100))
	var bmiStatus string
	switch {
	case bmiFactor < 18.5:
		bmiStatus = "Underweight"
	case bmiFactor >= 18.5 && bmiFactor < 24.9:
		bmiStatus = "Normal"
	case bmiFactor >= 25:
		bmiStatus = "Overweight"
	}

	bmi := domain.BMI{
		UserID: req.UserID,
		Height: req.Height,
		Weight: req.Weight,
		BMI:    bmiFactor,
		Status: bmiStatus,
	}

	err := uc.userStore.WithTransaction(ctx, func(tx *gorm.DB) error {
		if _, err := uc.userStore.CreatePersonalization(ctx, personalization); err != nil {
			return err
		}

		if _, err := uc.userStore.CreateBMI(ctx, bmi); err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (uc *UserUsecase) GetProfile(ctx context.Context, userID string) (domain.User, error) {
	profile, err := uc.userStore.GetProfile(ctx, userID)
	if err != nil {
		return domain.User{}, err
	}

	return profile, nil
}

func (uc *UserUsecase) GetDataForHomePage(ctx context.Context, userID string) (dto.HomePageResponse, error) {
	bmi, err := uc.userStore.GetCurrentBMI(ctx, userID)
	if err != nil {
		return dto.HomePageResponse{}, err
	}

	latestArticle, err := uc.userStore.GetLatestArticle(ctx)
	if err != nil {
		return dto.HomePageResponse{}, err
	}

	latestUserMission, err := uc.userStore.GetLatestUserMission(ctx, userID)
	if err != nil {
		return dto.HomePageResponse{}, err
	}

	tracker, err := uc.userStore.GetCurrentTracker(ctx, userID)
	if err != nil {
		return dto.HomePageResponse{}, err
	}

	response := dto.HomePageResponse{
		BMI:          bmi,
		Articles:     latestArticle,
		UserMissions: latestUserMission,
		Tracker:      tracker,
	}

	return response, nil
}
