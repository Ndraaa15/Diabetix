package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IBMIStore interface {
	GetCurrentBMI(ctx context.Context, userID string) (domain.BMI, error)
	CreateBMI(ctx context.Context, bmi domain.BMI) (domain.BMI, error)
	GetWeekPreviousBMI(ctx context.Context, userID string) ([]domain.BMI, error)
	GetAllBMI(ctx context.Context, userID string) ([]domain.BMI, error)
}

type BMIStore struct {
	db *gorm.DB
}

func NewBMIStore(db *gorm.DB) IBMIStore {
	return &BMIStore{
		db: db,
	}
}

func (r *BMIStore) GetCurrentBMI(ctx context.Context, userID string) (domain.BMI, error) {
	var bmi domain.BMI
	query := r.db.WithContext(ctx).Model(&domain.BMI{}).Order("created_at desc").Limit(1).Where("user_id = ?", userID).First(&bmi)
	if err := query.Error; err != nil {
		return domain.BMI{}, err
	}

	return bmi, nil
}

func (r *BMIStore) CreateBMI(ctx context.Context, bmi domain.BMI) (domain.BMI, error) {
	if err := r.db.WithContext(ctx).Model(&domain.BMI{}).Create(&bmi).Error; err != nil {
		return domain.BMI{}, err
	}

	return bmi, nil
}

func (r *BMIStore) GetWeekPreviousBMI(ctx context.Context, userID string) ([]domain.BMI, error) {
	var bmi []domain.BMI
	query := r.db.WithContext(ctx).Model(&domain.BMI{}).Order("created_at desc").Where("user_id = ?", userID).Limit(7).Find(&bmi)
	if err := query.Error; err != nil {
		return nil, err
	}

	return bmi, nil
}

func (r *BMIStore) GetAllBMI(ctx context.Context, userID string) ([]domain.BMI, error) {
	var bmi []domain.BMI
	query := r.db.WithContext(ctx).Model(&domain.BMI{}).Where("user_id = ?", userID).Find(&bmi)
	if err := query.Error; err != nil {
		return nil, err
	}

	return bmi, nil
}
