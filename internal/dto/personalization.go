package dto

type CreatePersonalizationRequest struct {
	UserID              string  `json:"userID" validate:"required"`
	Gender              string  `json:"gender" validate:"required"`
	FrequenceSport      string  `json:"frequenceSport" validate:"required"`
	Height              float64 `json:"height" validate:"required"`
	Weight              float64 `json:"weight" validate:"required"`
	DiabetesInheritance bool    `json:"diabetesInheritance" validate:"required"`
}
