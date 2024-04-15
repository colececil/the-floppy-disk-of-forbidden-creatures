package gen

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"time"
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
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Generate a creature description.",
			},
		},
	}

	response, err := g.openAiClient.CreateChatCompletion(ctx, request)
	if err != nil {
		return "The summoning has failed."
	}

	return response.Choices[0].Message.Content
}
