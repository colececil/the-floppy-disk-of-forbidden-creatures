package gen

import (
	"context"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
	"github.com/sashabaranov/go-openai"
	"strings"
)

// CreatureGenerator generates creature descriptions and images.
type CreatureGenerator struct {
	messageProvider *messages.MessageProvider
	openAiClient    *openai.Client
}

// NewCreatureGenerator creates a new CreatureGenerator with the given OpenAI API key.
func NewCreatureGenerator(messageProvider *messages.MessageProvider, apiKey string) *CreatureGenerator {
	return &CreatureGenerator{
		messageProvider: messageProvider,
		openAiClient:    openai.NewClient(apiKey),
	}
}

// GenerateDescription generates a description of the creature being summoned, based on the given attributes.
func (g *CreatureGenerator) GenerateDescription(ctx context.Context, creatureAttributes []string) string {
	var creatureAttributesList string
	for i, creatureAttribute := range creatureAttributes {
		creatureAttributesList += strings.ReplaceAll(creatureAttribute, ",", " ")
		if i < len(creatureAttributes)-1 {
			creatureAttributesList += ", "
		}
	}

	request := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleUser,
				Content: g.messageProvider.GetMessage(messages.CreatureDescriptionPrompt) + creatureAttributesList,
			},
		},
	}

	response, err := g.openAiClient.CreateChatCompletion(ctx, request)
	if err != nil {
		return ""
	}

	return response.Choices[0].Message.Content
}
