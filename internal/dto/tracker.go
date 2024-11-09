package dto

type CreateTrackerDetailRequest struct {
	FoodName     string  `json:"foodName"`
	FoodImage    string  `json:"foodImage"`
	Glucose      float64 `json:"glucose"`
	Calory       float64 `json:"calory"`
	Fat          float64 `json:"fat"`
	Protein      float64 `json:"protein"`
	Carbohydrate float64 `json:"carbohydrate"`
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
