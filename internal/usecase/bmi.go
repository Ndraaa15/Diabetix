package usecase

import (
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
)

type IBMIUsecase interface {
}

type BMIUsecase struct {
	bmiStore store.IBMIStore
	gemini   *gemini.Gemini
}

func NewBMIUsecase(bmiStore store.IBMIStore, gemini *gemini.Gemini) IBMIUsecase {
	return &BMIUsecase{
		bmiStore: bmiStore,
		gemini:   gemini,
	}
}
