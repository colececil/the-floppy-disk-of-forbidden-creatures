package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

// Message is a UI component that displays a message to the player. It implements tea.Model.
type Message struct {
	id                 int
	text               string
	responseComponent  tea.Model
	responseReceived   bool
	charactersRendered int
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
	Response string
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
				// Not all characters have been rendered yet, so schedule another tick.
				cmd = tea.Tick(playInterval, func(t time.Time) tea.Msg {
					return nextCharMsg{id: m.id}
				})
			} else {
				// The message has been fully rendered, so if the response component is an input component, allow it to
				// begin receiving input.
				if _, ok := m.responseComponent.(Input); ok {
					cmd = func() tea.Msg { return InputSetEnabledMsg{Id: m.id, Enabled: true} }
				}
			}
			return m, cmd
		}
	case tea.KeyMsg:
		// If the presses Enter after the message has been fully rendered, send the response.
		if msg.Type == tea.KeyEnter && !m.responseReceived && m.charactersRendered == len(m.text) {
			var cmd tea.Cmd
			var response string

			// If the response component is an input component, get the response and disable it to stop receiving input.
			if input, ok := m.responseComponent.(Input); ok {
				response = input.Value()
				if len(response) == 0 {
					// Don't allow an empty response.
					return m, nil
				}
				cmd = func() tea.Msg { return InputSetEnabledMsg{Id: m.id, Enabled: false} }
			}

			m.responseReceived = true

			cmd = tea.Batch(
				cmd,
				func() tea.Msg { return MessageResponseMsg{Response: response} },
			)
			return m, cmd
		}
	}

	var cmd tea.Cmd
	m.responseComponent, cmd = m.responseComponent.Update(msg)
	return m, cmd
}

// View implements tea.Model by returning the message as a string to be rendered.
func (m Message) View() string {
	runes := []rune(m.text)
	visibleText := string(runes[:m.charactersRendered])
	view := WrapText(visibleText)

	if m.responseReceived {
		view = InactiveTextStyle.Render(view)
		if _, ok := m.responseComponent.(Input); ok {
			response := InactiveTextStyle.Render(m.responseComponent.View())
			view += "\n" + response
		}
		view += "\n\n"
	} else {
		if m.charactersRendered == len(m.text) {
			response := SecondaryTextStyle.Render(m.responseComponent.View())
			view += "\n" + response
		}
	}

	return view
}
