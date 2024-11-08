package usecase

import (
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
)

type ITrackerUsecase interface {
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
