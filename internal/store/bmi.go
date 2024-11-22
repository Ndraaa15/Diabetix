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
	GetPersonalizationByUserID(ctx context.Context, userID string) (domain.Personalization, error)
	GetUserByID(ctx context.Context, userID string) (domain.User, error)
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
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

func (r *BMIStore) GetPersonalizationByUserID(ctx context.Context, userID string) (domain.Personalization, error) {
	var personalization domain.Personalization
	query := r.db.WithContext(ctx).Model(&domain.Personalization{}).Where("user_id = ?", userID).First(&personalization)
	if err := query.Error; err != nil {
		return domain.Personalization{}, err
	}

	return personalization, nil
}

func (r *BMIStore) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
	var profile domain.User
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Preload("Personalization").Preload("Level").Where("id = ?", userID).First(&profile).Error; err != nil {
		return domain.User{}, err
	}

	return profile, nil
}

func (r *BMIStore) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	tx := r.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	if err := fn(tx); err != nil {
		if rbErr := tx.Rollback().Error; rbErr != nil {
			return rbErr
		}
		return err
	}

	return tx.Commit().Error
}
