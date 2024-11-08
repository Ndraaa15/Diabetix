package usecase

import "github.com/Ndraaa15/diabetix-server/internal/store"

type IArticleUsecase interface {
}

type ArticleUsecase struct {
	articleStore store.IArticleStore
}

func NewArticleUsecase(articleStore store.IArticleStore) IArticleUsecase {
	return &ArticleUsecase{
		articleStore: articleStore,
	}
}
