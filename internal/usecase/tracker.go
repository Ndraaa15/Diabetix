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
	PredictFood(ctx context.Context, fileHeader *multipart.FileHeader, file multipart.File, userID string) (interface{}, error)
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

func (uc *TrackerUsecase) PredictFood(ctx context.Context, fileHeader *multipart.FileHeader, file multipart.File, userID string) (interface{}, error) {
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

	data, err := uc.gemini.GenerateNutritionFood(ctx, imageData, tracker.TrackerDetails)
	if err != nil {
		return dto.PredictFoodResponse{}, err
	}

	return data, nil
}
