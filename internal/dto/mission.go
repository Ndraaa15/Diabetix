package dto

type UpdateUserMissionRequest struct {
	Status string `json:"status" validate:"required"`
}
