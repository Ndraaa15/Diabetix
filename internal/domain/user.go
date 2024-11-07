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
	IsActive  bool         `json:"is_active" gorm:"type:boolean;not null;default:false"`
	Password  string       `json:"password" gorm:"type:varchar(255);not null"`
	Gender    Gender       `json:"gender" gorm:"type:integer;not null"`
	CreatedAt sql.NullTime `json:"created_at" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt sql.NullTime `json:"updated_at" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type Gender uint8

const (
	GenderUnknown Gender = 0
	GenderMale    Gender = 1
	GenderFemale  Gender = 2
)

var GenderMap = map[Gender]string{
	GenderMale:   "Male",
	GenderFemale: "Female",
}

func (s Gender) String() string {
	return GenderMap[s]
}

func (s Gender) Value() uint8 {
	return uint8(s)
}
