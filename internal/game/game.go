package game

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/gen"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/messages"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/ui"
)

// Game executes the game logic. It implements tea.Model.
type Game struct {
	messageProvider   *messages.MessageProvider
	creatureGenerator *gen.CreatureGenerator
	currentState      gameState
	uiMessages        []ui.Message
	uiSummoningCircle ui.SummoningCircle
	playerResponses   []string
}

// New creates a new Game.
func New(messageProvider *messages.MessageProvider, creatureGenerator *gen.CreatureGenerator) *Game {
	return &Game{
		messageProvider:   messageProvider,
		creatureGenerator: creatureGenerator,
	}
}

// gameState represents the current state of the game.
type gameState int

const (
	introState gameState = iota
	promptingState
	summoningState
)

// addUiMessage adds a new message to the UI.
type addUiMessageMsg struct {
	uiMessage ui.Message
}

// beginSummoning initializes the summoning circle.
type beginSummoningMsg struct{}

// exitGameMsg exits the game.
type exitGameMsg struct{}

// Init implements tea.Model by returning a tea.Cmd that updates the game state.
func (g *Game) Init() tea.Cmd {
	return g.updateGameState
}

// Update implements tea.Model by updating the model based on the given message.
func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.Type == tea.KeyCtrlC || msg.Type == tea.KeyEsc {
			return g, tea.Quit
		}
		// For all other key messages, don't return, since other components may need the key message.
	case tea.WindowSizeMsg:
		ui.UpdateTerminalSize(msg.Width, msg.Height)
		return g, nil
	case addUiMessageMsg:
		g.uiMessages = append(g.uiMessages, msg.uiMessage)
		return g, msg.uiMessage.Init()
	case ui.MessageResponseMsg:
		if len(msg.Response) > 0 {
			g.playerResponses = append(g.playerResponses, msg.Response)
		}
		return g, g.updateGameState
	case beginSummoningMsg:
		g.uiMessages = nil
		g.uiSummoningCircle = ui.NewSummoningCircle(g.messageProvider.GetMessage(messages.SummoningMessage))

		cmd := tea.Batch(
			g.uiSummoningCircle.Init(),
			g.generateCreatureDescription,
		)
		return g, cmd
	case exitGameMsg:
		return g, tea.Quit
	}

	var cmd tea.Cmd
	for i, uiMessage := range g.uiMessages {
		updatedUiMessage, uiMessageCmd := uiMessage.Update(msg)
		g.uiMessages[i] = updatedUiMessage.(ui.Message)
		cmd = tea.Batch(cmd, uiMessageCmd)
	}

	updatedSummoningCircle, summoningCircleCmd := g.uiSummoningCircle.Update(msg)
	g.uiSummoningCircle = updatedSummoningCircle.(ui.SummoningCircle)
	cmd = tea.Batch(cmd, summoningCircleCmd)

	return g, cmd
}

// View implements tea.Model by returning the model as a string to be rendered.
func (g *Game) View() string {
	if g.currentState == summoningState && len(g.uiMessages) == 0 {
		return g.uiSummoningCircle.View()
	}

	var view string
	for _, uiMessage := range g.uiMessages {
		view += uiMessage.View()
	}
	return ui.BaseStyle.Render(view)
}

// updateGameState advances the game state.
func (g *Game) updateGameState() tea.Msg {
	switch g.currentState {
	case introState:
		switch len(g.uiMessages) {
		case 0:
			return g.addNewUiMessage(g.messageProvider.GetMessage(messages.IntroMessage))
		case 1:
			return g.addNewUiMessage(g.messageProvider.GetMessage(messages.BeginRitualMessage))
		default:
			g.currentState = promptingState
			return g.addNewUiPrompt()
		}
	case promptingState:
		if len(g.playerResponses) < 5 {
			return g.addNewUiPrompt()
		} else {
			g.currentState = summoningState
			return beginSummoningMsg{}
		}
	case summoningState:
		if len(g.uiMessages) > 1 {
			return exitGameMsg{}
		}
		return g.addNewUiMessage(g.messageProvider.GetMessage(messages.EndingMessage))
	}

	return nil
}

// generateCreatureDescription generates a description of the creature being summoned.
func (g *Game) generateCreatureDescription() tea.Msg {
	description := g.creatureGenerator.GenerateDescription(g.playerResponses)
	return g.addNewUiMessage(description)
}

// addNewUiMessage adds a new message to the UI.
func (g *Game) addNewUiMessage(text string) tea.Msg {
	id := len(g.uiMessages)
	uiPlaceholder := ui.NewPlaceholder(g.messageProvider.GetMessage(messages.AwaitingAcknowledgementMessage))
	uiMessage := ui.NewMessage(id, text, uiPlaceholder)
	return addUiMessageMsg{uiMessage: uiMessage}
}

// addNewUiPrompt adds a new prompt to the UI.
func (g *Game) addNewUiPrompt() tea.Msg {
	id := len(g.uiMessages)
	uiInput := ui.NewInput(id)
	uiMessage := ui.NewMessage(id, g.messageProvider.GetPrompt(), uiInput)
	return addUiMessageMsg{uiMessage: uiMessage}
}
