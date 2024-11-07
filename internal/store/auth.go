package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IAuthStore interface {
	CreateUser(ctx context.Context, data domain.User) error
	GetUserByID(ctx context.Context, id string) (domain.User, error)
	GetUserByEmail(ctx context.Context, email string) (domain.User, error)
	UpdateUser(ctx context.Context, data domain.User) error
}

type AuthStore struct {
	db *gorm.DB
}

func NewAuthStore(db *gorm.DB) IAuthStore {
	return &AuthStore{
		db: db,
	}
}

func (s *AuthStore) CreateUser(ctx context.Context, data domain.User) error {
	if err := s.db.Model(domain.User{}).Create(&data).Error; err != nil {
		return err
	}

	return nil
}

func (s *AuthStore) GetUserByID(ctx context.Context, id string) (domain.User, error) {
	var user domain.User
	if err := s.db.Model(domain.User{}).Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthStore) GetUserByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	if err := s.db.Model(domain.User{}).Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthStore) UpdateUser(ctx context.Context, data domain.User) error {
	if err := s.db.Model(domain.User{}).Where("id = ?", data.ID).Updates(&data).Error; err != nil {
		return err
	}

	return nil
}
