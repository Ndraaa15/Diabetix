package store

import (
	"context"
	"strings"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"gorm.io/gorm"
)

type IArticleStore interface {
	GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.Article, error)
}

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) IArticleStore {
	return &ArticleStore{
		db: db,
	}
}

func (r *ArticleStore) GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.Article, error) {
	var articles []domain.Article
	query := r.db.WithContext(ctx).Model(&domain.Article{}).Find(&articles)
	if filter.Keyword != "" {
		query.Where("LOWER(keyword) = ?", strings.ToLower(filter.Keyword))
	}

	if err := query.Error; err != nil {
		return nil, err
	}

	return articles, nil
}
