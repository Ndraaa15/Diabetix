package dto

import "github.com/Ndraaa15/diabetix-server/internal/domain"

type CreateBMIRequest struct {
	Height float64 `json:"height" validate:"required"`
	Weight float64 `json:"weight" validate:"required"`
}

type BMIResponse struct {
	CurrentBMI      domain.BMI   `json:"currentBMI"`
	WeekPreviousBMI []domain.BMI `json:"weekPreviousBMI"`
	AllBMI          []domain.BMI `json:"allBMI"`
}
