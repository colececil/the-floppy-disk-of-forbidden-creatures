package game

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/ui"
)

// Game executes the game logic.
type Game struct {
	messageProvider *messages.MessageProvider
	uiMessages      []*ui.Message
	uiInput         textinput.Model
}

// NewGame creates a new Game.
func NewGame() *Game {
	return &Game{
		messageProvider: messages.NewMessageProvider(),
		uiInput:         textinput.New(),
	}
}

// addUiMessage adds a new message to the UI.
type addUiMessageMsg struct {
	uiMessage *ui.Message
}

// Init performs initialization when the model is first created.
func (g *Game) Init() tea.Cmd {
	return tea.Batch(g.updateGameState, textinput.Blink)
}

// Update is called when a message is received, and it returns an updated model and an optional command.
func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC {
			return g, tea.Quit
		}
	case addUiMessageMsg:
		g.uiMessages = append(g.uiMessages, msg.uiMessage)
		return g, msg.uiMessage.Init()
	}

	cmd := g.updateGameState

	for _, uiMessage := range g.uiMessages {
		_, uiMessageCmd := uiMessage.Update(msg)
		cmd = tea.Batch(cmd, uiMessageCmd)
	}

	var inputCmd tea.Cmd
	g.uiInput, inputCmd = g.uiInput.Update(msg)
	cmd = tea.Batch(cmd, inputCmd)

	return g, cmd
}

// View renders the model as a string to be printed. This is called after every Update.
func (g *Game) View() string {
	var view string
	for _, uiMessage := range g.uiMessages {
		view += uiMessage.View()
	}
	return view
}

// updateGameState updates the game state.
func (g *Game) updateGameState() tea.Msg {
	if len(g.uiMessages) == 0 {
		uiIntroMessage := ui.NewMessage(
			g.messageProvider.GetMessage(messages.IntroMessage),
			g.messageProvider.GetMessage(messages.AwaitingAcknowledgementMessage),
		)
		return addUiMessageMsg{uiMessage: uiIntroMessage}
	}

	return nil
}
