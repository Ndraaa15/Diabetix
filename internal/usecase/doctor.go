package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
)

type IDoctorUsecase interface {
	GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error)
	GetDoctorByID(ctx context.Context, doctorID uint64) (domain.Doctor, error)
}

type DoctorUsecase struct {
	doctorStore store.IDoctorStore
}

func NewDoctorUsecase(doctorStore store.IDoctorStore) IDoctorUsecase {
	return &DoctorUsecase{
		doctorStore: doctorStore,
	}
}

func (uc *DoctorUsecase) GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error) {
	return uc.doctorStore.GetAllDoctor(ctx, filter)
}

func (uc *DoctorUsecase) GetDoctorByID(ctx context.Context, doctorID uint64) (domain.Doctor, error) {
	return uc.doctorStore.GetDoctorByID(ctx, doctorID)
}
