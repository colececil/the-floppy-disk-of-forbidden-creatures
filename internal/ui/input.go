package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/audio"
	"unicode"
)

// Input is a UI component that accepts text input from the player. It implements tea.Model.
type Input struct {
	id           int
	backingInput textinput.Model
}

// NewInput creates a new Input with the given text input model.
func NewInput(id int) Input {
	backingInput := textinput.New()
	return Input{
		id:           id,
		backingInput: backingInput,
	}
}

// InputSetEnabledMsg is a tea.Msg used to indicate that the text input with the given ID should be enabled or disabled.
type InputSetEnabledMsg struct {
	Id      int
	Enabled bool
}

// Init implements tea.Model by calling Focus on the backing text input.
func (i Input) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model by calling Update on the backing text input.
func (i Input) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case InputSetEnabledMsg:
		if msg.Id == i.id {
			if msg.Enabled {
				cmd = i.backingInput.Focus()
			} else {
				i.backingInput.Blur()
			}
		}
	case tea.KeyMsg:
		if msg.Type == tea.KeyBackspace || msg.Type == tea.KeyDelete {
			_ = audio.Play(audio.QuietTapSoundEffect, nil, false)
		} else if msg.Type == tea.KeyRunes && unicode.IsGraphic(msg.Runes[0]) {
			_ = audio.Play(audio.TapSoundEffect, nil, true)
		}
	}

	var updateCmd tea.Cmd
	i.backingInput, updateCmd = i.backingInput.Update(msg)
	cmd = tea.Batch(cmd, updateCmd)

	return i, cmd
}

// View implements tea.Model by returning the backing text input's view.
func (i Input) View() string {
	return i.backingInput.View()
}

// Value returns the text input's value.
func (i Input) Value() string {
	return i.backingInput.Value()
}
