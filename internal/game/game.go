package game

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Game executes the game logic.
type Game struct {
	PreviousText string
	ActiveText   string
}

// NewGame creates a new Game.
func NewGame() *Game {
	return &Game{}
}

// Init performs initialization when the model is first created.
func (g *Game) Init() tea.Cmd {
	return g.updateGameState
}

// Update is called when a message is received, and it returns an updated model and an optional command.
func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return g, tea.Quit
		}
	}

	return g, g.updateGameState
}

// View renders the model as a string to be printed. This is called after every Update.
func (g *Game) View() string {
	return g.PreviousText + "\n\n" + g.ActiveText + "\n\n"
}

func (g *Game) updateGameState() tea.Msg {
	if (g.PreviousText == "") && (g.ActiveText == "") {
		g.ActiveText = messages[introKey]
		return nil
	}

	return nil

	//fmt.Println(messages[introKey])
	//fmt.Println()
	//
	//fmt.Println(messages[beginRitualKey])
	//promptGenerator := offerings.NewPromptGenerator()
	//fmt.Println(promptGenerator.GetPrompt())
}
