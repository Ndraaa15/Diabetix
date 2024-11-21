package usecase

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/internal/dto"
	"github.com/Ndraaa15/diabetix-server/internal/store"
	"github.com/Ndraaa15/diabetix-server/pkg/errx"
	"github.com/kataras/iris/v12"
)

type IArticleUsecase interface {
	GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.ArticleResponse, error)
	CreateLikes(ctx context.Context, userID string, articleID uint64) error
	DeleteLikes(ctx context.Context, userID string, articleID uint64) error
	GetArticleByID(ctx context.Context, articleID uint64) (domain.ArticleResponse, error)
}

type ArticleUsecase struct {
	articleStore store.IArticleStore
}

func NewArticleUsecase(articleStore store.IArticleStore) IArticleUsecase {
	return &ArticleUsecase{
		articleStore: articleStore,
	}
}

func (uc *ArticleUsecase) GetArticles(ctx context.Context, filter dto.GetArticlesFilter) ([]domain.ArticleResponse, error) {
	articles, err := uc.articleStore.GetArticles(ctx, filter)
	if err != nil {
		return nil, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get articles").
			WithError(err)
	}

	return articles, nil
}

func (uc *ArticleUsecase) CreateLikes(ctx context.Context, userID string, articleID uint64) error {
	articleLikes := domain.ArticleLike{
		UserID:    userID,
		ArticleID: articleID,
	}

	err := uc.articleStore.CreateLikes(ctx, articleLikes)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to create likes").
			WithError(err)
	}

	return nil
}

func (uc *ArticleUsecase) DeleteLikes(ctx context.Context, userID string, articleID uint64) error {
	articleLikes := domain.ArticleLike{
		UserID:    userID,
		ArticleID: articleID,
	}

	err := uc.articleStore.DeleteLikes(ctx, articleLikes)
	if err != nil {
		return errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to delete likes").
			WithError(err)
	}

	return nil
}

func (uc *ArticleUsecase) GetArticleByID(ctx context.Context, articleID uint64) (domain.ArticleResponse, error) {
	article, err := uc.articleStore.GetArticleByID(ctx, articleID)
	if err != nil {
		return domain.ArticleResponse{}, errx.New().
			WithCode(iris.StatusInternalServerError).
			WithMessage("Failed to get article by id").
			WithError(err)
	}

	return article, nil
}
