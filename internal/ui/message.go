package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/x/ansi"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/audio"
	"math/rand/v2"
	"time"
)

// Message is a UI component that displays a message to the player. It implements tea.Model.
type Message struct {
	id                       int
	text                     string
	responseComponent        tea.Model
	responseReceived         bool
	charactersRendered       int
	useBuzzForScrollingSound bool
}

// NewMessage creates a new Message.
func NewMessage(id int, text string, responseComponent tea.Model) Message {
	return Message{
		id:                       id,
		text:                     text,
		responseComponent:        responseComponent,
		useBuzzForScrollingSound: rand.IntN(2) == 0,
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
			m.charactersRendered = min(m.charactersRendered+1, len(m.text))
			m.playScrollingSoundEffect()

			var cmd tea.Cmd
			if m.charactersRendered < len(m.text) {
				// Not all characters have been rendered yet, so schedule another tick.
				cmd = tea.Tick(playInterval, func(t time.Time) tea.Msg {
					return nextCharMsg{id: m.id}
				})
			} else {
				if m.useBuzzForScrollingSound {
					audio.ResetLastSegmentPlayed(audio.DoubleBuzzSoundEffect)
				} else {
					audio.ResetLastSegmentPlayed(audio.RhythmicClicksAndBuzzesSoundEffect)
				}

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
		if msg.Type == tea.KeyEnter && !m.responseReceived {
			if m.charactersRendered != len(m.text) {
				// The message has not yet been fully rendered. Just skip to the end of the animation.
				m.charactersRendered = len(m.text)
				return m, func() tea.Msg { return nextCharMsg{id: m.id} }
			}

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
			_ = audio.Play(getRandomBeepSoundEffect(), nil, false)

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
	view := ansi.Wrap(visibleText, TerminalWidth, "")

	if m.responseReceived {
		if _, ok := m.responseComponent.(Input); ok {
			view = lipgloss.JoinVertical(lipgloss.Left, view, m.responseComponent.View())
		}
		view = InactiveTextStyle.Render(view)
	} else if m.charactersRendered == len(m.text) {
		response := ansi.Wrap(m.responseComponent.View(), TerminalWidth, "")
		response = SecondaryTextStyle.Render(response)
		view = lipgloss.JoinVertical(lipgloss.Left, view, response)
	}

	return PrimaryTextStyle.
		MarginBottom(1).
		Render(view)
}

// playScrollingSoundEffect plays the scrolling sound effect.
func (m Message) playScrollingSoundEffect() {
	if m.useBuzzForScrollingSound {
		audioSegmentIndex := getNextAudioSegmentIndexForScrolling(audio.DoubleBuzzSoundEffect, 2)
		_ = audio.Play(audio.DoubleBuzzSoundEffect, &audioSegmentIndex, false)
	} else {
		audioSegmentIndex := getNextAudioSegmentIndexForScrolling(audio.RhythmicClicksAndBuzzesSoundEffect, 8)
		_ = audio.Play(audio.RhythmicClicksAndBuzzesSoundEffect, &audioSegmentIndex, false)
	}
}

// getNextAudioSegmentIndexForScrolling returns the index of the next audio segment to play for the scrolling effect.
func getNextAudioSegmentIndexForScrolling(filename audio.SoundEffectFilename, numSegments int) int {
	lastIndex := audio.LastSegmentPlayed(filename)
	if lastIndex == -1 {
		return 1
	}
	return (lastIndex % numSegments) + 1
}

// getRandomBeepSoundEffect returns a random beep sound effect by its filename.
func getRandomBeepSoundEffect() audio.SoundEffectFilename {
	availableSoundEffects := []audio.SoundEffectFilename{
		audio.HighPitchedBeepSoundEffect,
		audio.LongLowPitchedBeepSoundEffect,
		audio.ShortLowPitchedBeepSoundEffect,
	}
	return availableSoundEffects[rand.IntN(len(availableSoundEffects))]
}
