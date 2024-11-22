package usecase

import (
	"context"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
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
	personalization, err := uc.bmiStore.GetPersonalizationByUserID(ctx, userID)
	if err != nil {
		return domain.BMI{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get personalization").
			WithError(err)
	}

	user, err := uc.bmiStore.GetUserByID(ctx, userID)
	if err != nil {
		return domain.BMI{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get user").
			WithError(err)
	}

	age := uint8(time.Now().Year() - user.Birth.Year())
	var maxGlucose float64
	if user.Personalization.Gender == domain.PersonalizationGenderMale {
		maxGlucose = 66.5 + (13.7 * req.Weight) + (5 * req.Weight) - (6.8 * float64(age))
	} else {
		maxGlucose = 65.5 + (9.6 * req.Weight) + (1.8 * req.Weight) - (4.7 * float64(age))
	}

	if user.Personalization.DiabetesInheritance {
		maxGlucose *= 0.05
	} else {
		maxGlucose *= 0.1
	}

	if user.Personalization.FrequenceSport == domain.PersonalizationFrequenceSportOncePerWeek {
		maxGlucose *= 1.2
	} else if user.Personalization.FrequenceSport == domain.PersonalizationFrequenceSportOnceToThreePerWeek {
		maxGlucose *= 1.3
	} else if user.Personalization.FrequenceSport == domain.PersonalizationFrequenceSportFourToFiveTimesPerWeek {
		maxGlucose *= 1.550
	} else if user.Personalization.FrequenceSport == domain.PersonalizationFrequenceSportFiveToSevenTimesPerWeek {
		maxGlucose *= 1.725
	}

	personalization.MaxGlucose = maxGlucose / 4.0

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

	data := domain.BMI{
		UserID: userID,
		Weight: req.Weight,
		Height: req.Height,
		BMI:    bmiFactor,
		Status: bmiStatus,
	}

	err = uc.bmiStore.WithTransaction(ctx, func(tx *gorm.DB) error {
		if err := tx.WithContext(ctx).Model(&domain.BMI{}).Create(&data).Error; err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to create BMI").
				WithError(err)
		}

		if err := tx.WithContext(ctx).Model(&domain.Personalization{}).Where("user_id = ?", userID).Updates(personalization).Error; err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to update personalization").
				WithError(err)
		}

		return nil
	})

	if err != nil {
		return domain.BMI{}, err
	}

	return data, nil
}
