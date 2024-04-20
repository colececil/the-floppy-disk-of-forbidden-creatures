package gen

import (
	"context"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
	"github.com/sashabaranov/go-openai"
	"strings"
	"time"
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
func (g *CreatureGenerator) GenerateDescription(creatureAttributes []string) string {
	var creatureAttributesList string
	for i, creatureAttribute := range creatureAttributes {
		creatureAttributesList += strings.ReplaceAll(creatureAttribute, ",", " ")
		if i < len(creatureAttributes)-1 {
			creatureAttributesList += ", "
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
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
		return g.messageProvider.GetMessage(messages.SummoningErrorMessage)
	}

	return response.Choices[0].Message.Content
}
