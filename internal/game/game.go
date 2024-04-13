package game

import "fmt"
import "github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/offerings"

// Game executes the game logic.
type Game struct {
}

// NewGame creates a new Game.
func NewGame() *Game {
	return &Game{}
}

// Start starts the game.
func (g *Game) Start() {
	fmt.Println(messages[introKey])
	fmt.Println()

	promptGenerator := offerings.NewPromptGenerator()
	fmt.Println(promptGenerator.GetPrompt())
}
