package offerings

import "math/rand"

// PromptGenerator generates prompts asking the player about their offerings.
type PromptGenerator struct {
	selectedPrompts    map[string]bool
	numPromptsSelected int
}

// NewPromptGenerator creates a new PromptGenerator.
func NewPromptGenerator() *PromptGenerator {
	return &PromptGenerator{
		selectedPrompts: make(map[string]bool),
	}
}

// GetPrompt returns a random prompt, ensuring that it has not already been selected.
func (g *PromptGenerator) GetPrompt() string {
	if g.numPromptsSelected >= len(prompts) {
		panic("no more prompts available")
	}
	for {
		prompt := prompts[rand.Intn(len(prompts))]
		if !g.selectedPrompts[prompt] {
			g.selectedPrompts[prompt] = true
			g.numPromptsSelected++
			return prompt
		}
	}
}
