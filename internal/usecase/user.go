package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/kataras/iris/v12"
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
	user, err := uc.userStore.GetProfile(ctx, req.UserID)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusNotFound).
			WithMessage("User not found").
			WithError(err)
	}

	if !user.IsActive {
		return errx.New().
			WithCode(iris.StatusForbidden).
			WithMessage("User is not active").
			WithError(errors.New("User is not active"))
	}

	frequenceSport, err := util.ParsePersonalizationFrequenceSport(req.FrequenceSport)
	if err != nil {
		return err
	}
	gender, err := util.ParsePersonalizationGender(req.Gender)
	if err != nil {
		return err
	}

	age := uint8(time.Now().Year() - user.Birth.Year())
	var maxGlucose float64
	if gender == domain.PersonalizationGenderMale {
		maxGlucose = 66.5 + (13.7 * req.Weight) + (5 * req.Weight) - (6.8 * float64(age))
	} else {
		maxGlucose = 65.5 + (9.6 * req.Weight) + (1.8 * req.Weight) - (4.7 * float64(age))
	}

	if req.DiabetesInheritance {
		maxGlucose *= 0.05
	} else {
		maxGlucose *= 0.1
	}

	personalization := domain.Personalization{
		UserID:              req.UserID,
		Gender:              gender,
		Age:                 age,
		FrequenceSport:      frequenceSport,
		MaxGlucose:          maxGlucose / 4.0,
		DiabetesInheritance: req.DiabetesInheritance,
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

	bmi := domain.BMI{
		UserID: req.UserID,
		Height: req.Height,
		Weight: req.Weight,
		BMI:    bmiFactor,
		Status: bmiStatus,
	}

	err = uc.userStore.WithTransaction(ctx, func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Model(&domain.Personalization{}).Create(&personalization).Error; err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to create personalization").
				WithError(err)
		}

		if err := tx.WithContext(ctx).Model(&domain.BMI{}).Create(&bmi).Error; err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to create BMI").
				WithError(err)
		}

		return nil
	})

	if err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to create personalization").
			WithError(err)
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

	latestArticle, err := uc.userStore.GetLatestArticle(ctx, userID)
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
