package domain

type Article struct {
	ID int64 `json:"id" gorm:"primaryKey"`
}
