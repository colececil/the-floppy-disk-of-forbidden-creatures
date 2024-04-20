package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/game"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/gen"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
)

// This value is overridden at build time using `-ldflags`.
var apiKey = "change me"

func main() {
	messageProvider := messages.NewMessageProvider()
	creatureGenerator := gen.NewCreatureGenerator(messageProvider, apiKey)
	teaProgram := tea.NewProgram(game.New(messageProvider, creatureGenerator), tea.WithAltScreen())
	_, err := teaProgram.Run()
	if err != nil {
		panic(err)
	}
}
