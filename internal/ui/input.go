package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// Input is a UI component that accepts text input from the player. It implements tea.Model.
type Input struct {
	backingInput textinput.Model
}

// NewInput creates a new Input with the given text input model.
func NewInput() Input {
	backingInput := textinput.New()
	backingInput.Focus()
	return Input{backingInput: backingInput}
}

// Init implements tea.Model by calling Focus on the backing text input.
func (i Input) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model by calling Update on the backing text input.
func (i Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	i.backingInput, cmd = i.backingInput.Update(msg)
	return i, cmd
}

// View implements tea.Model by returning the backing text input's view.
func (i Input) View() string {
	return i.backingInput.View()
}

// Disable disables the component.
func (i Input) Disable() {
	// Todo: Fix this. I might need to call it from Update instead?
	i.backingInput.Blur()
}

// Value returns the text input's value.
func (i Input) Value() string {
	return i.backingInput.Value()
}
