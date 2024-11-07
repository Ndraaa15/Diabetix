package gemini

import (
	"context"

	"github.com/Ndraaa15/diabetix-server/pkg/env"
	"github.com/google/generative-ai-go/genai"
	"go.uber.org/zap"
	"google.golang.org/api/option"
)

type Gemini struct {
	model *genai.GenerativeModel
}

func NewGemini(env env.Env) *Gemini {
	ctx := context.Background()

	client, err := genai.NewClient(ctx, option.WithAPIKey(env.GeminiApiKey))
	if err != nil {
		zap.S().Fatal(err)
	}

	model := client.GenerativeModel(env.GeminiModel)

	return &Gemini{
		model: model,
	}
}

func (g *Gemini) GenerateResponseForProblem(ctx context.Context, text string, picture []byte) error {

	return nil
}
