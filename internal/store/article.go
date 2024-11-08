package store

import "gorm.io/gorm"

type IArticleStore interface {
}

type ArticleStore struct {
	db *gorm.DB
}

func NewArticleStore(db *gorm.DB) IArticleStore {
	return &ArticleStore{
		db: db,
	}
}
