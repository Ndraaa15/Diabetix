package usecase

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
	"gorm.io/gorm"
)

type ITrackerUsecase interface {
	PredictFood(ctx context.Context, fileHeader *multipart.FileHeader, file multipart.File, userID string) (dto.PredictFoodResponse, error)
	AddFood(ctx context.Context, req dto.CreateTrackerDetailRequest, userID string) error
	GetAllTracker(ctx context.Context, userID string) ([]domain.Tracker, error)
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
	tracker, err := uc.trackerStore.GetCurrentTracker(ctx, userID, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local))
	if err != nil {
		return dto.PredictFoodResponse{}, err
	}

	dataByte, err := fileHeader.Open()
	if err != nil {
		return dto.PredictFoodResponse{}, err
	}
	defer dataByte.Close()

	imageData := make([]byte, fileHeader.Size)
	_, err = dataByte.Read(imageData)
	if err != nil {
		return dto.PredictFoodResponse{}, err
	}

	resultGenerate, err := uc.gemini.GenerateNutritionFood(ctx, imageData, tracker.TrackerDetails)
	if err != nil {
		return dto.PredictFoodResponse{}, err
	}

	personalization, err := uc.trackerStore.GetPersonalization(ctx, userID)
	if err != nil {
		return dto.PredictFoodResponse{}, err
	}
	var levelGlucose string
	if resultGenerate.Glucose < personalization.MaxGlucose*0.7 {
		levelGlucose = "Low"
	} else if resultGenerate.Glucose < personalization.MaxGlucose {
		levelGlucose = "Normal"
	} else {
		levelGlucose = "High"
	}

	response := dto.PredictFoodResponse{
		FoodName:       resultGenerate.FoodName,
		Glucose:        resultGenerate.Glucose,
		Calories:       resultGenerate.Calories,
		Fat:            resultGenerate.Fat,
		Carbohydrate:   resultGenerate.Carbohydrate,
		Protein:        resultGenerate.Protein,
		Advice:         resultGenerate.Advice,
		MaxGlucose:     personalization.MaxGlucose,
		CurrentGlucose: tracker.TotalGlucose,
		LevelGlucose:   levelGlucose,
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

		report, err = uc.trackerStore.GetCurrentReport(ctx, userID, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			report, err = uc.trackerStore.CreateReport(ctx, domain.Report{
				UserID:    userID,
				StartDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
				EndDate:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day()+7, 23, 59, 59, 0, time.Local),
				Advice:    "",
			})

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		fmt.Println("report", report)

		tracker, err = uc.trackerStore.GetCurrentTracker(ctx, userID, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local))
		if errors.Is(err, gorm.ErrRecordNotFound) {
			tracker, err = uc.trackerStore.CreateTracker(ctx, domain.Tracker{
				TotalGlucose: req.Glucose,
				Status:       "Low",
				UserID:       userID,
				ReportID:     report.ID,
			})

			if err != nil {
				return err
			}
		} else if err != nil {
			return err
		}

		fmt.Println("tracker", tracker)

		personalization, err := uc.trackerStore.GetPersonalization(ctx, userID)
		if err != nil {
			return err
		}

		fmt.Println("personalization", personalization)

		var status string
		if req.Glucose < personalization.MaxGlucose*0.7 {
			status = "Low"
		} else if req.Glucose < personalization.MaxGlucose {
			status = "Normal"
		} else {
			status = "High"
		}

		tracker.TotalGlucose = tracker.TotalGlucose + req.Glucose
		tracker.Status = status

		data := domain.TrackerDetail{
			FoodName:     req.FoodName,
			FoodImage:    req.FoodImage,
			Glucose:      req.Glucose,
			Calory:       req.Calory,
			Fat:          req.Fat,
			Protein:      req.Protein,
			Carbohydrate: req.Carbohydrate,
			TrackerID:    tracker.ID,
		}

		if err := uc.trackerStore.UpdateTracker(ctx, tracker); err != nil {
			return err
		}
		fmt.Println("tracker", data)
		if err := uc.trackerStore.CreateTrackerDetail(ctx, data); err != nil {
			return err
		}

		return nil
	})
}

func (r *TrackerUsecase) GetAllTracker(ctx context.Context, userID string) ([]domain.Tracker, error) {
	trackers, err := r.trackerStore.GetAllTracker(ctx, userID)
	if err != nil {
		return nil, err
	}

	return trackers, nil
}
