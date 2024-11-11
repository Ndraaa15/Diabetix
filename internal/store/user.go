package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IUserStore interface {
	GetProfile(ctx context.Context, userID string) (domain.User, error)
	CreateBMI(ctx context.Context, bmi domain.BMI) (domain.BMI, error)
	CreatePersonalization(ctx context.Context, personalization domain.Personalization) (domain.Personalization, error)
	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error
	GetLatestArticle(ctx context.Context) ([]domain.Article, error)
	GetLatestUserMission(ctx context.Context, userID string) ([]domain.UserMission, error)
	GetCurrentBMI(ctx context.Context, userID string) (domain.BMI, error)
	GetCurrentTracker(ctx context.Context, userID string) (domain.Tracker, error)
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) IUserStore {
	return &UserStore{
		db: db,
	}
}

func (r *UserStore) GetProfile(ctx context.Context, userID string) (domain.User, error) {
	var profile domain.User
	if err := r.db.WithContext(ctx).Model(&domain.User{}).Preload("Personalization").Preload("Level").Where("id = ?", userID).First(&profile).Error; err != nil {
		return domain.User{}, err
	}

	return profile, nil
}

func (r *UserStore) CreateBMI(ctx context.Context, bmi domain.BMI) (domain.BMI, error) {
	if err := r.db.WithContext(ctx).Model(&domain.BMI{}).Create(&bmi).Error; err != nil {
		return domain.BMI{}, err
	}

	return bmi, nil
}

func (r *UserStore) CreatePersonalization(ctx context.Context, personalization domain.Personalization) (domain.Personalization, error) {
	if err := r.db.WithContext(ctx).Model(&domain.Personalization{}).Create(&personalization).Error; err != nil {
		return domain.Personalization{}, err
	}

	return personalization, nil
}

func (r *UserStore) WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error {
	return r.db.Transaction(fn)
}

func (r *UserStore) GetLatestArticle(ctx context.Context) ([]domain.Article, error) {
	var articles []domain.Article
	if err := r.db.WithContext(ctx).Model(&domain.Article{}).Order("created_at desc").Limit(3).Find(&articles).Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *UserStore) GetLatestUserMission(ctx context.Context, userID string) ([]domain.UserMission, error) {
	var missions []domain.UserMission
	if err := r.db.WithContext(ctx).Model(&domain.UserMission{}).Where("user_id = ?", userID).Order("created_at desc").Limit(3).Find(&missions).Error; err != nil {
		return nil, err
	}

	return missions, nil
}

func (r *UserStore) GetCurrentBMI(ctx context.Context, userID string) (domain.BMI, error) {
	var bmi domain.BMI
	if err := r.db.WithContext(ctx).Model(&domain.BMI{}).Where("user_id = ?", userID).Order("created_at desc").First(&bmi).Error; err != nil {
		return domain.BMI{}, err
	}

	return bmi, nil
}

func (r *UserStore) GetCurrentTracker(ctx context.Context, userID string) (domain.Tracker, error) {
	var tracker domain.Tracker
	if err := r.db.WithContext(ctx).Model(&domain.Tracker{}).Where("user_id = ?", userID).Order("created_at desc").First(&tracker).Error; err != nil {
		return domain.Tracker{}, err
	}

	return tracker, nil
}
