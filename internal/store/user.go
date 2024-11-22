package store

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"gorm.io/gorm"
)

type IUserStore interface {
	GetUserByID(ctx context.Context, userID string) (domain.User, error)

	CreateBMI(ctx context.Context, bmi domain.BMI) (domain.BMI, error)
	GetCurrentBMI(ctx context.Context, userID string) (domain.BMI, error)

	CreatePersonalization(ctx context.Context, personalization domain.Personalization) (domain.Personalization, error)

	WithTransaction(ctx context.Context, fn func(tx *gorm.DB) error) error

	GetLatestArticle(ctx context.Context, userID string) ([]domain.ArticleResponse, error)
	GetLatestUserMission(ctx context.Context, userID string) ([]domain.UserMission, error)
	GetCurrentTracker(ctx context.Context, userID string) (domain.Tracker, error)
	UpdateUser(ctx context.Context, user domain.User) error
}

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) IUserStore {
	return &UserStore{
		db: db,
	}
}

func (r *UserStore) GetUserByID(ctx context.Context, userID string) (domain.User, error) {
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

func (r *UserStore) GetLatestArticle(ctx context.Context, userID string) ([]domain.ArticleResponse, error) {
	var articles []domain.ArticleResponse
	queryBuilder := r.db.Table("articles").
		Select("articles.*, COALESCE(COUNT(article_likes.article_id), 0) AS likes, "+
			"CASE WHEN article_likes.user_id = ? THEN true ELSE false END AS is_liked_by_current_user", userID).
		Joins("LEFT JOIN article_likes ON article_likes.article_id = articles.id").
		Group("articles.id, article_likes.user_id").
		Scan(&articles).Limit(3).Order("created_at desc")

	if err := queryBuilder.Error; err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *UserStore) GetLatestUserMission(ctx context.Context, userID string) ([]domain.UserMission, error) {
	var missions []domain.UserMission
	if err := r.db.WithContext(ctx).Model(&domain.UserMission{}).Preload("Mission").Where("user_id = ?", userID).Order("created_at desc").Limit(3).Find(&missions).Error; err != nil {
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

func (r *UserStore) UpdateUser(ctx context.Context, user domain.User) error {
	return r.db.WithContext(ctx).Model(&domain.User{}).Where("id = ?", user.ID).Updates(&user).Error
}
