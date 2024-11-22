package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"gorm.io/gorm"
)

type IDoctorStore interface {
}

type DoctorStore struct {
	db *gorm.DB
}

func NewDoctorStore(db *gorm.DB) IDoctorStore {
	return &DoctorStore{
		db: db,
	}
}

func (r *DoctorStore) GetAllDoctor(ctx context.Context, filter dto.GetDoctorsFilter) ([]domain.Doctor, error) {
	var doctors []domain.Doctor

	queryBuilder := r.db.Model(&domain.Doctor{}).Find(&doctors)
	if filter.Keyword != "" {
		queryBuilder.Where("LOWER(name) LIKE ?", "%"+filter.Keyword+"%")
		queryBuilder.Where("LOWER(specialist) LIKE ?", "%"+filter.Keyword+"%")
	}

	return doctors, nil
}
