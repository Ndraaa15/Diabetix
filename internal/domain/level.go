package domain

import "time"

type Level struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"type:varchar(255);not null"`
	TotalExp  uint64    `json:"totalExp" gorm:"type:integer;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
