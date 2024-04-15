package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/game"
)

func main() {
	teaProgram := tea.NewProgram(game.New(), tea.WithAltScreen())
	_, err := teaProgram.Run()
	if err != nil {
		panic(err)
	}
}
