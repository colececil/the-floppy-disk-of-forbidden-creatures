package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/game"
)

const apiKey = "API key must be set at build time"

func main() {
	teaProgram := tea.NewProgram(game.New(apiKey), tea.WithAltScreen())
	_, err := teaProgram.Run()
	if err != nil {
		panic(err)
	}
}
