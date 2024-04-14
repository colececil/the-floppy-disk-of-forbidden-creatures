package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"time"
)

// Message is a UI component that displays a message to the player.
type Message struct {
	text               string
	isAcknowledged     bool
	charactersRendered int
	maxWidth           int
}

// NewMessage creates a new Message.
func NewMessage(text string) *Message {
	return &Message{
		text: text,
	}
}

// playInterval is the rate at which the message plays.
const playInterval = 10 * time.Millisecond

// nextCharMsg is a tea.Msg used to tell the model to play the next character.
type nextCharMsg struct{}

// Init implements tea.Model by returning a tea.Cmd that schedules a nextCharMsg.
func (m *Message) Init() tea.Cmd {
	return tea.Tick(playInterval, nextChar)
}

// Update implements tea.Model by updating the model based on the given message.
func (m *Message) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg.(type) {
	case nextCharMsg:
		m.charactersRendered++
		var cmd tea.Cmd
		if m.charactersRendered < len(m.text) {
			cmd = tea.Tick(playInterval, nextChar)
		}
		return m, cmd
	case tea.WindowSizeMsg:
		m.maxWidth = msg.(tea.WindowSizeMsg).Width
		return m, nil
	}

	return m, nil
}

// View implements tea.Model by returning the message as a string to be rendered.
func (m *Message) View() string {
	runes := []rune(m.text)
	visibleText := string(runes[:m.charactersRendered])
	return wordwrap.String(visibleText, m.maxWidth)
}

// nextChar is a tea.Cmd that tells the model to play the next character.
func nextChar(_ time.Time) tea.Msg {
	return nextCharMsg{}
}
