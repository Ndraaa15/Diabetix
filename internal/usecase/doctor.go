package usecase

import "github.com/Ndraaa15/diabetix-server/internal/store"

type IDoctorUsecase interface {
}

type DoctorUsecase struct {
	doctorStore store.IDoctorStore
}

func NewDoctorUsecase(doctorStore store.IDoctorStore) IDoctorUsecase {
	return &DoctorUsecase{
		doctorStore: doctorStore,
	}
}
