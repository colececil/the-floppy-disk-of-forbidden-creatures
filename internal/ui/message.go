package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/wordwrap"
	"time"
)

// Message is a UI component that displays a message to the player. It implements tea.Model.
type Message struct {
	id                 int
	text               string
	responseComponent  tea.Model
	responseReceived   bool
	charactersRendered int
	maxWidth           int
}

// NewMessage creates a new Message.
func NewMessage(id int, text string, responseComponent tea.Model) Message {
	return Message{
		id:                id,
		text:              text,
		responseComponent: responseComponent,
	}
}

// playInterval is the rate at which the message plays.
const playInterval = 10 * time.Millisecond

// MessageResponseMsg is a tea.Msg used to indicate that the message has been acknowledged.
type MessageResponseMsg struct {
	response string
}

// nextCharMsg is a tea.Msg used to tell the model to play the next character.
type nextCharMsg struct {
	id int
}

// Init implements tea.Model by returning a tea.Cmd that schedules a nextCharMsg.
func (m Message) Init() tea.Cmd {
	tickCmd := tea.Tick(playInterval, func(t time.Time) tea.Msg {
		return nextCharMsg{id: m.id}
	})
	return tea.Batch(tickCmd, m.responseComponent.Init())
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
		if msg.Type == tea.KeyEnter && !m.responseReceived && m.charactersRendered == len(m.text) {
			m.responseReceived = true

			var response string
			// Todo: Get response if the responseComponent is a text input.

			return m, func() tea.Msg { return MessageResponseMsg{response: response} }
		}
	case tea.WindowSizeMsg:
		m.maxWidth = msg.Width
		return m, nil
	}

	var cmd tea.Cmd
	m.responseComponent, cmd = m.responseComponent.Update(msg)
	return m, cmd
}

// View implements tea.Model by returning the message as a string to be rendered.
func (m Message) View() string {
	runes := []rune(m.text)
	visibleText := string(runes[:m.charactersRendered])
	view := wordwrap.String(visibleText, m.maxWidth)

	if m.responseReceived {
		view += "\n\n"
	} else {
		if m.charactersRendered == len(m.text) {
			view += "\n" + m.responseComponent.View()
		}
	}

	return view
}
