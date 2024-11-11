package dto

type CreatePersonalizationRequest struct {
	UserID         string  `json:"userID" validate:"required"`
	Gender         string  `json:"gender" validate:"required"`
	Age            uint8   `json:"age" validate:"required"`
	FrequenceSport uint8   `json:"frequenceSport" validate:"required"`
	Height         float64 `json:"height" validate:"required"`
	Weight         float64 `json:"weight" validate:"required"`
}
