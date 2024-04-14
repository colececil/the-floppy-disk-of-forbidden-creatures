package game

import (
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
}

// NewGame creates a new Game.
func NewGame() *Game {
	return &Game{
		messageProvider: messages.NewMessageProvider(),
	}
}

// Init performs initialization when the model is first created.
func (g *Game) Init() tea.Cmd {
	return g.updateGameState
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
		return g, nil
	default:
		return g, g.updateGameState
	}
}

// View renders the model as a string to be printed. This is called after every Update.
func (g *Game) View() string {
	return wordwrap.String(g.previousText, g.terminalWidth) + "\n\n" +
		wordwrap.String(g.activeText, g.terminalWidth) + "\n\n"
}

func (g *Game) updateGameState() tea.Msg {
	if (g.previousText == "") && (g.activeText == "") {
		g.activeText = g.messageProvider.GetMessage(messages.IntroMessage)
	}

	return nil
}
