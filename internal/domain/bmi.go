package domain

import "time"

type BMI struct {
	ID        uint64    `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    string    `json:"userID" gorm:"type:varchar(19);not null"`
	Height    float64   `json:"height" gorm:"type:numeric;not null"`
	Weight    float64   `json:"weight" gorm:"type:numeric;not null"`
	Status    BMIStatus `json:"status" gorm:"type:varchar(255);not null"`
	BMI       float64   `json:"bmi" gorm:"type:numeric;not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"type:timestamp;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"type:timestamp;not null;autoUpdateTime"`
}

type BMIStatus string

const (
	BMIStatusUnderweight BMIStatus = "Underweight"
	BMIStatusNormal      BMIStatus = "Normal"
	BMIStatusOverweight  BMIStatus = "Overweight"
	BMIStatusObesityI    BMIStatus = "Obesity I"
	BMIStatusObesityII   BMIStatus = "Obesity II"
	BMIStatusObesityIII  BMIStatus = "Obesity III"
)
