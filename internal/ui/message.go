package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"time"
)

// Message is a UI component that displays a message to the player. It implements tea.Model.
type Message struct {
	id                          int
	text                        string
	awaitingAcknowledgementText string
	isAcknowledged              bool
	charactersRendered          int
	maxWidth                    int
}

// NewMessage creates a new Message.
func NewMessage(id int, text string, awaitingAcknowledgementText string) Message {
	return Message{
		id:                          id,
		text:                        text,
		awaitingAcknowledgementText: awaitingAcknowledgementText,
	}
}

// playInterval is the rate at which the message plays.
const playInterval = 10 * time.Millisecond

// AcknowledgeMsg is a tea.Msg used to indicate that the message has been acknowledged.
type AcknowledgeMsg struct{}

// nextCharMsg is a tea.Msg used to tell the model to play the next character.
type nextCharMsg struct {
	id int
}

// Init implements tea.Model by returning a tea.Cmd that schedules a nextCharMsg.
func (m Message) Init() tea.Cmd {
	return tea.Tick(playInterval, func(t time.Time) tea.Msg {
		return nextCharMsg{id: m.id}
	})
}

// Update implements tea.Model by updating the model based on the given message.
func (m Message) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case nextCharMsg:
		if msg.id == m.id {
			m.charactersRendered++
			var cmd tea.Cmd
			if m.charactersRendered < len(m.text) {
				cmd = tea.Tick(playInterval, func(t time.Time) tea.Msg {
					return nextCharMsg{id: m.id}
				})
			}
			return m, cmd
		}
	case tea.KeyMsg:
		if msg.Type == tea.KeyEnter && !m.isAcknowledged && m.charactersRendered == len(m.text) {
			m.isAcknowledged = true
			return m, func() tea.Msg { return AcknowledgeMsg{} }
		}
	case tea.WindowSizeMsg:
		m.maxWidth = msg.Width
		return m, nil
	}

	return m, nil
}

// View implements tea.Model by returning the message as a string to be rendered.
func (m Message) View() string {
	runes := []rune(m.text)
	visibleText := string(runes[:m.charactersRendered])
	view := wordwrap.String(visibleText, m.maxWidth)

	if m.isAcknowledged {
		view += "\n\n"
	} else {
		if m.charactersRendered == len(m.text) {
			view += "\n" + m.awaitingAcknowledgementText
		}
	}

	return view
}
