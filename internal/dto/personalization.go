package dto

type CreatePersonalizationRequest struct {
	UserID         string `json:"userID" validate:"required"`
	Gender         string `json:"gender"	validate:"required"`
	Age            uint8  `json:"age" validate:"required"`
	FrequenceSport string `json:"frequenceSport" validate:"required"`
}
