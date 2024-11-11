package dto

import "github.com/Ndraaa15/diabetix-server/internal/domain"

type HomePageResponse struct {
	BMI          domain.BMI           `json:"bmi"`
	Tracker      domain.Tracker       `json:"tracker"`
	Articles     []domain.Article     `json:"articles"`
	UserMissions []domain.UserMission `json:"userMissions"`
}
