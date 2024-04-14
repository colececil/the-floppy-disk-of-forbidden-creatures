package game

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
	"github.com/muesli/reflow/wordwrap"
)

// Game executes the game logic.
type Game struct {
	messageProvider *messages.MessageProvider
	terminalWidth   int
	terminalHeight  int
	previousText    string
	activeText      string
	input           textinput.Model
}

// NewGame creates a new Game.
func NewGame() *Game {
	return &Game{
		messageProvider: messages.NewMessageProvider(),
		input:           textinput.New(),
	}
}

// Init performs initialization when the model is first created.
func (g *Game) Init() tea.Cmd {
	return tea.Batch(g.updateGameState, textinput.Blink)
}

// Update is called when a message is received, and it returns an updated model and an optional command.
func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		g.terminalWidth, g.terminalHeight = msg.Width, msg.Height
		return g, nil
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return g, tea.Quit
		}
	}

	var inputCmd tea.Cmd
	g.input, inputCmd = g.input.Update(msg)
	return g, tea.Batch(g.updateGameState, inputCmd)
}

// View renders the model as a string to be printed. This is called after every Update.
func (g *Game) View() string {
	return wordwrap.String(g.previousText, g.terminalWidth) + "\n\n" +
		wordwrap.String(g.activeText, g.terminalWidth) + "\n\n" +
		g.input.View()
}

// updateGameState updates the game state.
func (g *Game) updateGameState() tea.Msg {
	if (g.previousText == "") && (g.activeText == "") {
		g.previousText = g.messageProvider.GetMessage(messages.IntroMessage)
		g.activeText = g.messageProvider.GetMessage(messages.BeginRitualMessage)
	}

	return nil
}
