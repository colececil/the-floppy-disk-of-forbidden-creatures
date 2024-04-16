package gen

import (
	"context"
	"github.com/sashabaranov/go-openai"
	"strings"
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

const creatureDescriptionPrompt = "You are the narrator for a game about summoning monsters. Your task is to " +
	"generate a description of the monster being summoned, based on several responses given by the player. The " +
	"description should be a single paragraph, which both narrates the appearance of the monster from the summoning " +
	"circle, and describes what the monster is like. It should also end with a narration explaining what becomes of " +
	"the player (who should be addressed as \"you\") once the monster they summoned has appeared." +
	"\n\n" +
	"The responses given by the player may be things that can directly apply to the monster's appearance, or they " +
	"indirectly provide an attribute of the monster. Please be creative and unpredictable in how the player's " +
	"responses influence what the monster is like. Also, it's better if the description brings up the things " +
	"influenced by the player responses in a different order than they are provided to you. It's also better if " +
	"the description doesn't include the exact wording of the player responses, but applies them in a more subtle " +
	"manner." +
	"\n\n" +
	"Please use descriptive language that paints a mental picture, and keep in mind that the game has a foreboding " +
	"and Lovecraftian tone. Your response should be a single paragraph no longer than 8 sentences. Do not include " +
	"anything other than the description in your response. The player responses are provided below, separated by " +
	"commas:" +
	"\n\n"

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
				Content: creatureDescriptionPrompt + creatureAttributesList,
			},
		},
	}

	response, err := g.openAiClient.CreateChatCompletion(ctx, request)
	if err != nil {
		return "The summoning has failed."
	}

	return response.Choices[0].Message.Content
}
