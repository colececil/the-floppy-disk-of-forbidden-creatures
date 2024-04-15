package game

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/ui"
)

// Game executes the game logic. It implements tea.Model.
type Game struct {
	messageProvider *messages.MessageProvider
	currentState    gameState
	uiMessages      []ui.Message
}

// New creates a new Game.
func New() *Game {
	return &Game{
		messageProvider: messages.NewMessageProvider(),
	}
}

// gameState represents the current state of the game.
type gameState int

const (
	introState gameState = iota
	promptingState
)

// addUiMessage adds a new message to the UI.
type addUiMessageMsg struct {
	uiMessage ui.Message
}

// Init implements tea.Model by returning a tea.Cmd that updates the game state.
func (g *Game) Init() tea.Cmd {
	return tea.Batch(g.updateGameState, textinput.Blink)
}

// Update implements tea.Model by updating the model based on the given message.
func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.Type == tea.KeyEsc {
			return g, tea.Quit
		}
	case addUiMessageMsg:
		g.uiMessages = append(g.uiMessages, msg.uiMessage)
		return g, msg.uiMessage.Init()
	case ui.AcknowledgeMsg:
		return g, g.updateGameState
	}

	var cmd tea.Cmd
	for i, uiMessage := range g.uiMessages {
		updatedUiMessage, uiMessageCmd := uiMessage.Update(msg)
		g.uiMessages[i] = updatedUiMessage.(ui.Message)
		cmd = tea.Batch(cmd, uiMessageCmd)
	}
	return g, cmd
}

// View implements tea.Model by returning the model as a string to be rendered.
func (g *Game) View() string {
	var view string
	for _, uiMessage := range g.uiMessages {
		view += uiMessage.View()
	}
	return view
}

// updateGameState advances the game state.
func (g *Game) updateGameState() tea.Msg {
	switch g.currentState {
	case introState:
		switch len(g.uiMessages) {
		case 0:
			uiMessage := ui.NewMessage(
				len(g.uiMessages),
				g.messageProvider.GetMessage(messages.IntroMessage),
				g.messageProvider.GetMessage(messages.AwaitingAcknowledgementMessage),
			)
			return addUiMessageMsg{uiMessage: uiMessage}
		case 1:
			uiMessage := ui.NewMessage(
				len(g.uiMessages),
				g.messageProvider.GetMessage(messages.BeginRitualMessage),
				g.messageProvider.GetMessage(messages.AwaitingAcknowledgementMessage),
			)
			return addUiMessageMsg{uiMessage: uiMessage}
		default:
			g.currentState = promptingState
			uiMessage := ui.NewMessage(
				len(g.uiMessages),
				g.messageProvider.GetPrompt(),
				g.messageProvider.GetMessage(messages.AwaitingAcknowledgementMessage),
			)
			return addUiMessageMsg{uiMessage: uiMessage}
		}
	case promptingState:
	}

	return nil
}
