package domain

import "time"

type Article struct {
	ID        uint64          `json:"id" gorm:"primaryKey;autoIncrement"`
	Image     string          `json:"image" gorm:"type:text;not null"`
	Title     string          `json:"title" gorm:"type:varchar(255);not null"`
	Category  ArticleCategory `json:"category" gorm:"type:varchar(255);not null"`
	Body      string          `json:"body" gorm:"type:text;not null"`
	Likes     uint64          `json:"likes" gorm:"type:integer;not null"`
	Date      time.Time       `json:"date" gorm:"type:timestamp;not null;autoCreateTime"`
	CreatedAt time.Time       `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time       `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type ArticleCategory uint8

const (
	ArticleCategoryUnknown   ArticleCategory = 0
	ArticleCategoryHealth    ArticleCategory = 1
	ArticleCategorySport     ArticleCategory = 2
	ArticleCategoryFood      ArticleCategory = 3
	ArticleCategoryLifeStyle ArticleCategory = 4
)

var ArticleCategoryMap = map[ArticleCategory]string{
	ArticleCategoryHealth:    "Health",
	ArticleCategorySport:     "Sport",
	ArticleCategoryFood:      "Food",
	ArticleCategoryLifeStyle: "LifeStyle",
}

func (s ArticleCategory) String() string {
	return ArticleCategoryMap[s]
}

func (s ArticleCategory) Value() uint8 {
	return uint8(s)
}

type GetArticlesFilter struct {
	Keyword string
}
