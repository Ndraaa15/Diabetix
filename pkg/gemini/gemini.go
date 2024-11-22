package gemini

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/Ndraaa15/diabetix-server/internal/domain"
	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/google/generative-ai-go/genai"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type PredictFoodResponseWithGenai struct {
	FoodName      string  `json:"foodName"`
	Glucose       float64 `json:"glucose"`
	Calories      float64 `json:"calories"`
	Fat           float64 `json:"fat"`
	Carbohydrate  float64 `json:"carbohidrate"`
	Protein       float64 `json:"protein"`
	IndexGlycemic float64 `json:"indexGlycemic"`
	Advice        string  `json:"advice"`
}

type ReportAdvice struct {
	Advice string `json:"advice"`
}

type Gemini struct {
	model *genai.GenerativeModel
}

func NewGemini(env *env.Env) *Gemini {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(env.GeminiApiKey))
	if err != nil {
		zap.S().Fatalf("Failed to create Gemini client: %v", err)
	}

	model := client.GenerativeModel(env.GeminiModel)

	return &Gemini{
		model: model,
	}
}

func (g *Gemini) GenerateNutritionFood(ctx context.Context, picture []byte, previousFoods []domain.TrackerDetail) (PredictFoodResponseWithGenai, error) {
	previousFoodsJSON, err := json.Marshal(previousFoods)
	if err != nil {
		return PredictFoodResponseWithGenai{}, err
	}

	prompt := fmt.Sprintf(`
Analyze the food item in the provided image to predict its nutritional information, including its name, glucose level, calories, fat, carbohydrate, and protein content. 
Additionally, provide advice based on the user's food consumption history for today, which is provided as an array of JSON objects.

Here is the JSON structure for the response:
{
		"foodName": "string",
		"glucose": "float64",
		"calories": "float64",
		"fat": "float64",
		"carbohidrate": "float64",
		"protein": "float64",
		"indexGlycemic": "float64",
		"advice": "string"
}

Previous food consumption history (JSON format):
%s

Analyze the provided image and return a response in the JSON structure above. Provide dietary advice based on the cumulative nutritional values of today's meals. 
The advice should address balancing glucose, calorie, fat, carbohydrate, and protein intake.
`, previousFoodsJSON)

	genaiParts := []genai.Part{
		genai.Text(prompt),
		genai.ImageData("jpg", picture),
		genai.Text("Please provide using Indonesia Language"),
	}

	content, err := g.model.GenerateContent(ctx, genaiParts...)
	if err != nil {
		return PredictFoodResponseWithGenai{}, err
	}

	part := content.Candidates[0].Content.Parts[0]
	jsonByte, err := json.Marshal(part)
	if err != nil {
		return PredictFoodResponseWithGenai{}, nil
	}

	jsonStr, err := strconv.Unquote(string(jsonByte))
	if err != nil {
		return PredictFoodResponseWithGenai{}, err
	}

	jsonStr = strings.Replace(jsonStr, "```json", "", -1)
	jsonStr = strings.Replace(jsonStr, "```", "", -1)

	var response PredictFoodResponseWithGenai
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return PredictFoodResponseWithGenai{}, err
	}

	return response, nil
}

func (g *Gemini) GenerateReportAdvice(ctx context.Context, previousReports domain.Report) (ReportAdvice, error) {
	previousReportJSON, err := json.Marshal(previousReports)
	if err != nil {
		return ReportAdvice{}, err
	}

	prompt := fmt.Sprintf(`Analyze the user's daily reports to provide advice based on the user's daily glucose levels. The user's daily reports are provided as an array of JSON objects.

	Here is the JSON structure for the response:
	{
		"advice": "string"
	}

	Previous activity and previous food consumption (JSON format):
	%s
	`, previousReportJSON)

	genaiParts := []genai.Part{
		genai.Text(prompt),
		genai.Text("Please provide using Indonesia Language"),
	}

	content, err := g.model.GenerateContent(ctx, genaiParts...)
	if err != nil {
		return ReportAdvice{}, err
	}

	part := content.Candidates[0].Content.Parts[0]
	jsonByte, err := json.Marshal(part)
	if err != nil {
		return ReportAdvice{}, nil
	}

	jsonStr, err := strconv.Unquote(string(jsonByte))
	if err != nil {
		return ReportAdvice{}, err
	}

	jsonStr = strings.Replace(jsonStr, "```json", "", -1)
	jsonStr = strings.Replace(jsonStr, "```", "", -1)

	var response ReportAdvice
	err = json.Unmarshal([]byte(jsonStr), &response)
	if err != nil {
		return ReportAdvice{}, err
	}

	return response, nil
}
