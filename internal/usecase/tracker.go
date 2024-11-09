package usecase

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
)

type ITrackerUsecase interface {
	PredictFood(ctx context.Context, fileHeader *multipart.FileHeader, file multipart.File, userID string) (dto.PredictFoodResponse, error)
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

// func (uc *TrackerUsecase) AddFood(ctx context.Context, req dto.CreateTrackerDetailRequest, userID string) error {
// 	tracker, err := uc.trackerStore.GetCurrentTracker(ctx, userID, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local))
// 	if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return err
// 	}

// 	report, err := uc.trackerStore.GetCurrentReport(ctx, userID, time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local))

// 	if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		return err
// 	}

// 	report, err := uc.trackerStore.CreateReport(ctx, domain.Report{
// 		UserID:    userID,
// 		StartDate: time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 0, 0, 0, 0, time.Local),
// 		EndDate:   time.Date(time.Now().Year(), time.Now().Month(), time.Now().Day(), 23, 59, 59, 0, time.Local),
// 	})

// 	tracker = domain.Tracker{
// 		TotalGlucose: req.Glucose,
// 		Status:       "Low",
// 		UserID:       userID,
// 	}

// 	data := domain.TrackerDetail{
// 		FoodName:     req.FoodName,
// 		FoodImage:    req.FoodImage,
// 		Glucose:      req.Glucose,
// 		Calory:       req.Calory,
// 		Fat:          req.Fat,
// 		Protein:      req.Protein,
// 		Carbohydrate: req.Carbohydrate,
// 	}

// 	if err := uc.trackerStore.CreateTrackerDetail(ctx, data); err != nil {
// 		return err
// 	}

// 	return nil
// }
