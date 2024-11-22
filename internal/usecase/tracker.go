package usecase

import (
	"context"
	"errors"
	"mime/multipart"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/kataras/iris/v12"
	"gorm.io/gorm"
)

type ITrackerUsecase interface {
	PredictFood(ctx context.Context, fileHeader *multipart.FileHeader, file multipart.File, userID string) (dto.PredictFoodResponse, error)
	AddFood(ctx context.Context, req dto.CreateTrackerDetailRequest, userID string) error
	GetAllTracker(ctx context.Context, userID string) (dto.TrackerResponse, error)
}

type TrackerUsecase struct {
	trackerStore store.ITrackerStore
	gemini       *gemini.Gemini
}

func NewTrackerUsecase(trackerStore store.ITrackerStore, gemini *gemini.Gemini) ITrackerUsecase {
	return &TrackerUsecase{
		trackerStore: trackerStore,
		gemini:       gemini,
	}
}

func (uc *TrackerUsecase) PredictFood(ctx context.Context, fileHeader *multipart.FileHeader, file multipart.File, userID string) (dto.PredictFoodResponse, error) {
	tracker, err := uc.trackerStore.GetCurrentTracker(ctx, userID, util.GetCurrentDate())
	if err != nil {
		return dto.PredictFoodResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get current tracker").
			WithError(err)
	}

	dataByte, err := fileHeader.Open()
	if err != nil {
		return dto.PredictFoodResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to open file").
			WithError(err)

	}
	defer dataByte.Close()

	imageData := make([]byte, fileHeader.Size)
	_, err = dataByte.Read(imageData)
	if err != nil {
		return dto.PredictFoodResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to read file").
			WithError(err)
	}

	resultGenerate, err := uc.gemini.GenerateNutritionFood(ctx, imageData, tracker.TrackerDetails)
	if err != nil {
		return dto.PredictFoodResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to generate nutrition food").
			WithError(err)
	}

	personalization, err := uc.trackerStore.GetPersonalization(ctx, userID)
	if err != nil {
		return dto.PredictFoodResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get personalization").
			WithError(err)
	}

	var levelGlycemic string
	if resultGenerate.IndexGlycemic <= 55 {
		levelGlycemic = string(domain.TrackerDetailLevelGlycemicLow)
	} else if resultGenerate.IndexGlycemic >= 56 && resultGenerate.IndexGlycemic <= 69 {
		levelGlycemic = string(domain.TrackerDetailLevelGlycemicNormal)
	} else {
		levelGlycemic = string(domain.TrackerDetailLevelGlycemicHigh)
	}

	response := dto.PredictFoodResponse{
		FoodName:       resultGenerate.FoodName,
		Glucose:        resultGenerate.Glucose,
		Calories:       resultGenerate.Calories,
		Fat:            resultGenerate.Fat,
		Carbohydrate:   resultGenerate.Carbohydrate,
		Protein:        resultGenerate.Protein,
		IndexGlycemic:  resultGenerate.IndexGlycemic,
		Advice:         resultGenerate.Advice,
		MaxGlucose:     personalization.MaxGlucose,
		CurrentGlucose: tracker.TotalGlucose,
		LevelGlycemic:  levelGlycemic,
	}

	return response, nil
}

func (uc *TrackerUsecase) AddFood(ctx context.Context, req dto.CreateTrackerDetailRequest, userID string) error {
	return uc.trackerStore.WithTransaction(ctx, func(tx *gorm.DB) error {
		var (
			report  domain.Report
			tracker domain.Tracker
			err     error
		)

		report, err = uc.trackerStore.GetCurrentReport(ctx, userID, util.GetCurrentDate())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			report, err = uc.trackerStore.CreateReport(ctx, domain.Report{
				UserID:    userID,
				StartDate: util.GetCurrentDate(),
				EndDate:   util.GetCurrentDate().AddDate(0, 0, 7),
				Advice:    "",
			})

			if err != nil {
				return errx.New().
					WithCode(iris.StatusInternalServerError).
					WithMessage("Failed to create report").
					WithError(err)
			}
		} else if err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to get current report").
				WithError(err)
		}

		personalization, err := uc.trackerStore.GetPersonalization(ctx, userID)
		if err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to get personalization").
				WithError(err)
		}

		tracker, err = uc.trackerStore.GetCurrentTracker(ctx, userID, util.GetCurrentDate())
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tracker, err = uc.trackerStore.CreateTracker(ctx, domain.Tracker{
				TotalGlucose: req.Glucose,
				MaxGlucose:   personalization.MaxGlucose,
				Status:       "Low",
				UserID:       userID,
				ReportID:     report.ID,
			})

			if err != nil {
				return errx.New().
					WithCode(iris.StatusInternalServerError).
					WithMessage("Failed to create tracker").
					WithError(err)
			}
		} else if err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to get current tracker").
				WithError(err)
		}

		var levelGlycemic string
		if req.IndexGlycemic <= 55 {
			levelGlycemic = string(domain.TrackerDetailLevelGlycemicLow)
		} else if req.IndexGlycemic >= 56 && req.IndexGlycemic <= 69 {
			levelGlycemic = string(domain.TrackerDetailLevelGlycemicNormal)
		} else {
			levelGlycemic = string(domain.TrackerDetailLevelGlycemicHigh)
		}

		// Todo : change status with proper calculation
		var status string
		if tracker.TotalGlucose+req.Glucose > personalization.MaxGlucose {
			status = string(domain.TrackerStatusHigh)
		} else if tracker.TotalGlucose+req.Glucose <= personalization.MaxGlucose {
			status = string(domain.TrackerStatusNormal)
		} else if tracker.TotalGlucose+req.Glucose < personalization.MaxGlucose {
			status = string(domain.TrackerStatusLow)
		}

		tracker.TotalGlucose = tracker.TotalGlucose + req.Glucose
		tracker.Status = status

		data := domain.TrackerDetail{
			FoodName:      req.FoodName,
			FoodImage:     req.FoodImage,
			Glucose:       req.Glucose,
			Calory:        req.Calory,
			Fat:           req.Fat,
			Protein:       req.Protein,
			IndexGlycemic: req.IndexGlycemic,
			LevelGlycemic: levelGlycemic,
			Carbohydrate:  req.Carbohydrate,
			TrackerID:     tracker.ID,
		}

		if err := uc.trackerStore.UpdateTracker(ctx, tracker); err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to update tracker").
				WithError(err)
		}

		if err := uc.trackerStore.CreateTrackerDetail(ctx, data); err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to create tracker detail").
				WithError(err)
		}

		return nil
	})
}

func (r *TrackerUsecase) GetAllTracker(ctx context.Context, userID string) (dto.TrackerResponse, error) {
	trackers, err := r.trackerStore.GetAllTracker(ctx, userID)
	if err != nil {
		return dto.TrackerResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get all tracker").
			WithError(err)
	}

	currentTracker, err := r.trackerStore.GetCurrentTracker(ctx, userID, util.GetCurrentDate())
	if err != nil {
		return dto.TrackerResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get current tracker").
			WithError(err)
	}

	return dto.TrackerResponse{
		CurrentTracker: currentTracker,
		Trackers:       trackers,
	}, nil
}
