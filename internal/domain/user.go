package domain

import (
	"database/sql"
	"time"
)

type User struct {
	ID        string       `json:"id" gorm:"type:varchar(255);primaryKey"`
	Name      string       `json:"name" gorm:"type:varchar(255);not null"`
	Email     string       `json:"email" gorm:"type:varchar(255);not null;unique"`
	Birth     time.Time    `json:"birth" gorm:"type:date;not null"`
	IsActive  bool         `json:"isActive" gorm:"type:boolean;not null;default:false"`
	Password  string       `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt sql.NullTime `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt sql.NullTime `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}
