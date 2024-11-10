package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
)

type IArticleUsecase interface {
	GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.Article, error)
}

type ArticleUsecase struct {
	articleStore store.IArticleStore
}

func NewArticleUsecase(articleStore store.IArticleStore) IArticleUsecase {
	return &ArticleUsecase{
		articleStore: articleStore,
	}
}

func (uc *ArticleUsecase) GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.Article, error) {
	articles, err := uc.articleStore.GetArticles(ctx, filter)
	if err != nil {
		return nil, err
	}

	return articles, nil
}
