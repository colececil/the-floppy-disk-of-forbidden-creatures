package gen

import (
	"context"
	"github.com/sashabaranov/go-openai"
)

// CreatureGenerator generates creature descriptions and images.
type CreatureGenerator struct {
	openAiClient *openai.Client
}

// NewCreatureGenerator creates a new CreatureGenerator with the given OpenAI API key.
func NewCreatureGenerator(apiKey string) *CreatureGenerator {
	client := openai.NewClient(apiKey)
	return &CreatureGenerator{
		openAiClient: client,
	}
}

func (g *CreatureGenerator) GenerateDescription() string {
	// Todo: Set timeout.
	ctx := context.Background()

	request := openai.CompletionRequest{
		Model:     openai.GPT3Dot5Turbo,
		Prompt:    "Write a description of a creature",
		MaxTokens: 200,
	}

	response, err := g.openAiClient.CreateCompletion(ctx, request)
	if err != nil {
		return "Your summoning has failed."
	}

	return response.Choices[0].Text
}
