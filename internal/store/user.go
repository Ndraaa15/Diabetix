package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IUserStore interface {
	CreatePersonalizationAndBMI(ctx context.Context, personalization domain.Personalization, bmi domain.BMI) error
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) IUserStore {
	return &UserStore{
		db: db,
	}
}

func (r *UserStore) CreatePersonalizationAndBMI(ctx context.Context, personalization domain.Personalization, bmi domain.BMI) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Debug().Model(domain.Personalization{}).Create(&personalization).Error; err != nil {
			return err
		}

		if err := tx.Debug().Model(domain.BMI{}).Create(&bmi).Error; err != nil {
			return err
		}

		return nil
	})
}
