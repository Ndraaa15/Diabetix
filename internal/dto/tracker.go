package dto

type CreateTrackerDetailRequest struct {
	FoodName     string  `json:"foodName" validate:"required"`
	FoodImage    string  `json:"foodImage" validate:"required"`
	Glucose      float64 `json:"glucose" validate:"required"`
	Calory       float64 `json:"calory" validate:"required"`
	Fat          float64 `json:"fat" validate:"required"`
	Protein      float64 `json:"protein" validate:"required"`
	Carbohydrate float64 `json:"carbohydrate" validate:"required"`
}

type PredictFoodResponse struct {
	FoodName       string  `json:"foodName"`
	Glucose        float64 `json:"glucose"`
	LevelGlucose   string  `json:"levelGlucose"`
	Calories       float64 `json:"calories"`
	Fat            float64 `json:"fat"`
	Carbohydrate   float64 `json:"carbohydrate"`
	Protein        float64 `json:"protein"`
	Advice         string  `json:"advice"`
	CurrentGlucose float64 `json:"currentGlucose"`
	MaxGlucose     float64 `json:"maxGlucose"`
}
