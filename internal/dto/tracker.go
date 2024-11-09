package dto

type AddFoodRequest struct {
	Name string `json:"name"`
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
