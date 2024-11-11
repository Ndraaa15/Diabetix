package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/gemini"
)

type IReportUsecase interface {
	GetAllReport(ctx context.Context, userID string, filter dto.GetReportsFilter) ([]domain.Report, error)
}

type ReportUsecase struct {
	reportStore store.IReportStore
	gemini      *gemini.Gemini
}

func NewReportUsecase(reportStore store.IReportStore, gemini *gemini.Gemini) IReportUsecase {
	return &ReportUsecase{
		reportStore: reportStore,
		gemini:      gemini,
	}
}

func (uc *ReportUsecase) GetAllReport(ctx context.Context, userID string, filter dto.GetReportsFilter) ([]domain.Report, error) {
	reports, err := uc.reportStore.GetAllReport(ctx, userID, filter)
	if err != nil {
		return nil, err
	}

	return reports, nil
}
