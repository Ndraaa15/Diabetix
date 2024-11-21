package domain

import "time"

type ArticleResponse struct {
	ID                   uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Image                string          `json:"image" gorm:"type:text;not null"`
	Author               string          `json:"author" gorm:"type:varchar(255);not null"`
	Title                string          `json:"title" gorm:"type:varchar(255);not null"`
	Category             ArticleCategory `json:"category" gorm:"type:varchar(255);not null"`
	Likes                uint64          `json:"likes" gorm:"likes"`
	Body                 string          `json:"body" gorm:"type:text;not null"`
	Date                 time.Time       `json:"date" gorm:"type:timestamp;not null;autoCreateTime"`
	CreatedAt            time.Time       `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt            time.Time       `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
	IsLikedByCurrentUser bool            `json:"isLikedByCurrentUser" gorm:"isLikedByCurrentUser"`
}

type Article struct {
	ID        uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Image     string          `json:"image" gorm:"type:text;not null"`
	Author    string          `json:"author" gorm:"type:varchar(255);not null"`
	Title     string          `json:"title" gorm:"type:varchar(255);not null"`
	Category  ArticleCategory `json:"category" gorm:"type:varchar(255);not null"`
	Body      string          `json:"body" gorm:"type:text;not null"`
	Date      time.Time       `json:"date" gorm:"type:timestamp;not null;autoCreateTime"`
	CreatedAt time.Time       `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time       `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type ArticleCategory string

const (
	ArticleCategorySport     ArticleCategory = "Sport"
	ArticleCategoryFood      ArticleCategory = "Food"
	ArticleCategoryMindset   ArticleCategory = "Mindset"
	ArticleCategoryMedical   ArticleCategory = "Medical"
	ArticleCategoryKnowledge ArticleCategory = "Knowledge"
)

type ArticleLike struct {
	UserID    string    `json:"userID" gorm:"type:varchar(19);primaryKey"`
	ArticleID uint64    `json:"articleID" gorm:"type:integer;primaryKey"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
}
