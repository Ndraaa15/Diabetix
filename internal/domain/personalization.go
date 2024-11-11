package domain

import "time"

type Personalization struct {
	UserID         string    `json:"userID" gorm:"type:varchar(255);not null;primaryKey"`
	Gender         string    `json:"gender" gorm:"type:varchar(255);not null"`
	Age            uint8     `json:"age" gorm:"type:integer;not null"`
	FrequenceSport uint8     `json:"frequenceSport" gorm:"type:integer;not null"`
	MaxGlucose     float64   `json:"maxGlucose" gorm:"type:integer;not null"`
	CreatedAt      time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt      time.Time `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
