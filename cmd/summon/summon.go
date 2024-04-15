package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/game"
)

// This value is overridden at build time using `-ldflags`.
var apiKey = "change me"

func main() {
	teaProgram := tea.NewProgram(game.New(apiKey), tea.WithAltScreen())
	_, err := teaProgram.Run()
	if err != nil {
		panic(err)
	}
}
