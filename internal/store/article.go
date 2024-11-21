package store

import (
	"context"
	"strings"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"gorm.io/gorm"
)

type IArticleStore interface {
	GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.ArticleResponse, error)
	CreateLikes(ctx context.Context, data domain.ArticleLike) error
	DeleteLikes(ctx context.Context, data domain.ArticleLike) error
	GetArticleByID(ctx context.Context, articleID uint64) (domain.ArticleResponse, error)
}

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) IArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (r *ArticleStore) GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.ArticleResponse, error) {
	var articles []domain.ArticleResponse

	queryBuilder := r.db.Table("articles").
		Select("articles.*, COALESCE(COUNT(article_likes.article_id), 0) AS likes, "+
			"CASE WHEN article_likes.user_id = ? THEN true ELSE false END AS is_liked_by_current_user",
			filter.UserID).
		Joins("LEFT JOIN article_likes ON article_likes.article_id = articles.id").
		Group("articles.id, article_likes.user_id").
		Scan(&articles)

	if filter.Keyword != "" {
		queryBuilder = queryBuilder.Where("LOWER(articles.title) LIKE ?", "%"+strings.ToLower(filter.Keyword)+"%")
	}

	err := queryBuilder.Error
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (r *ArticleStore) CreateLikes(ctx context.Context, data domain.ArticleLike) error {
	err := r.db.WithContext(ctx).Model(&domain.ArticleLike{}).Create(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ArticleStore) DeleteLikes(ctx context.Context, data domain.ArticleLike) error {
	err := r.db.WithContext(ctx).Model(&domain.ArticleLike{}).Where("article_id = ? AND user_id = ?", data.ArticleID, data.UserID).Delete(&data).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *ArticleStore) GetArticleByID(ctx context.Context, articleID uint64) (domain.ArticleResponse, error) {
	var article domain.ArticleResponse

	err := r.db.Table("articles").
		Select("articles.*, COALESCE(COUNT(article_likes.article_id), 0) AS likes").
		Joins("LEFT JOIN article_likes ON article_likes.article_id = articles.id").
		Where("articles.id = ?", articleID).
		Group("articles.id").
		Scan(&article).Error

	if err != nil {
		return domain.ArticleResponse{}, err
	}

	return article, nil
}
