package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/Ndraaa15/diabetix-server/pkg/midtrans"
	"github.com/Ndraaa15/diabetix-server/pkg/util"
	"github.com/kataras/iris/v12"
)

type IDoctorUsecase interface {
	GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error)
	GetDoctorByID(ctx context.Context, doctorID uint64) (domain.Doctor, error)
	CreateConsultation(ctx context.Context, doctorScheduleID uint64, userID string) (string, error)
}

type DoctorUsecase struct {
	doctorStore store.IDoctorStore
	midtrans    *midtrans.Midtrans
}

func NewDoctorUsecase(doctorStore store.IDoctorStore, midtrans *midtrans.Midtrans) IDoctorUsecase {
	return &DoctorUsecase{
		doctorStore: doctorStore,
		midtrans:    midtrans,
	}
}

func (uc *DoctorUsecase) GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error) {
	return uc.doctorStore.GetAllDoctor(ctx, filter)
}

func (uc *DoctorUsecase) GetDoctorByID(ctx context.Context, doctorID uint64) (domain.Doctor, error) {
	return uc.doctorStore.GetDoctorByID(ctx, doctorID)
}

func (uc *DoctorUsecase) CreateConsultation(ctx context.Context, doctorScheduleID uint64, userID string) (string, error) {
	consultation := domain.Consultation{
		UserID:           userID,
		DoctorScheduleID: doctorScheduleID,
		IsDone:           false,
	}

	doctorSchedule, err := uc.doctorStore.GetDoctorScheduleByID(ctx, doctorScheduleID)
	if err != nil {
		return "", errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get doctor schedule").
			WithError(err)
	}

	if !doctorSchedule.IsOpen {
		return "", errx.New().
			WithCode(iris.StatusBadRequest).
			WithMessage("Doctor schedule is already booked")
	}

	doctorSchedule.IsOpen = false

	err = uc.doctorStore.WithTransaction(ctx, func(ctx context.Context) error {
		if err := uc.doctorStore.CreateConsultation(ctx, consultation); err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to create consultation").
				WithError(err)
		}

		if err := uc.doctorStore.UpdateDoctorSchedule(ctx, doctorSchedule); err != nil {
			return errx.New().
				WithCode(iris.StatusInternalServerError).
				WithMessage("Failed to update doctor schedule").
				WithError(err)
		}

		return nil
	})

	res, err := uc.midtrans.CreateTransaction(ctx, util.GenerateRandomString(9), 100000)
	if err != nil {
		return "", err
	}

	return res, nil
}
