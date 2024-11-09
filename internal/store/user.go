package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IUserStore interface {
	CreatePersonalization(ctx context.Context, data domain.Personalization) error
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) IUserStore {
	return &UserStore{
		db: db,
	}
}

func (r *UserStore) CreatePersonalization(ctx context.Context, data domain.Personalization) error {
	if err := r.db.Debug().Model(domain.Personalization{}).Create(&data).Error; err != nil {
		return err
	}

	return nil
}
