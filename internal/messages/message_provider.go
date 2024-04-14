package messages

import "math/rand"

// MessageProvider provides messages used in the game.
type MessageProvider struct {
	selectedPrompts    map[string]bool
	numPromptsSelected int
}

// NewMessageProvider creates a new MessageProvider.
func NewMessageProvider() *MessageProvider {
	return &MessageProvider{
		selectedPrompts: make(map[string]bool),
	}
}

// GetMessage returns the message for the given key.
func (p *MessageProvider) GetMessage(key MessageKey) string {
	return messages[key]
}

// GetPrompt returns a random prompt, ensuring that it has not already been selected.
func (p *MessageProvider) GetPrompt() string {
	if p.numPromptsSelected >= len(prompts) {
		panic("no more prompts available")
	}
	for {
		prompt := prompts[rand.Intn(len(prompts))]
		if !p.selectedPrompts[prompt] {
			p.selectedPrompts[prompt] = true
			p.numPromptsSelected++
			return prompt
		}
	}
}
