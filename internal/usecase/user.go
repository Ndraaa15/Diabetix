package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
)

type IUserUsecase interface {
	CreatePersonalization(ctx context.Context, req dto.CreatePersonalizationRequest) error
}

type UserUsecase struct {
	userStore store.IUserStore
}

func NewUserUsecase(userStore store.IUserStore) IUserUsecase {
	return &UserUsecase{
		userStore: userStore,
	}
}

func (uc *UserUsecase) CreatePersonalization(ctx context.Context, req dto.CreatePersonalizationRequest) error {
	data := domain.Personalization{
		UserID:         req.UserID,
		Gender:         req.Gender,
		Age:            req.Age,
		FrequenceSport: req.FrequenceSport,
	}

	if err := uc.userStore.CreatePersonalization(ctx, data); err != nil {
		return err
	}

	return nil
}
