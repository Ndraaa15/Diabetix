package domain

import (
	"time"
)

type User struct {
	ID              string          `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name            string          `json:"name" gorm:"type:varchar(255);not null"`
	Email           string          `json:"email" gorm:"type:varchar(255);not null;unique"`
	Birth           time.Time       `json:"birth" gorm:"type:date;not null"`
	IsActive        bool            `json:"isActive" gorm:"type:boolean;not null;default:false"`
	Password        string          `json:"password" gorm:"type:varchar(255);not null"`
	CurrentExp      uint64          `json:"currentExp" gorm:"type:integer;not null"`
	LevelID         uint64          `json:"levelID" gorm:"type:integer;not null"`
	Level           Level           `json:"level" gorm:"foreignKey:LevelID;references:ID"`
	Personalization Personalization `json:"personalization" gorm:"foreignKey:UserID;references:ID"`
	CreatedAt       time.Time       `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt       time.Time       `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
