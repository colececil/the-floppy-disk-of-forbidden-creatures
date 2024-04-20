package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"time"
)

// SummoningCircle is a UI component for displaying the summoning circle. It implements the tea.Model interface.
type SummoningCircle struct {
	summoningMessage string
	terminalWidth    int
	terminalHeight   int
	animationFrame   int
}

// NewSummoningCircle creates a new SummoningCircle.
func NewSummoningCircle(summoningMessage string, terminalWidth int, terminalHeight int) SummoningCircle {
	return SummoningCircle{
		summoningMessage: summoningMessage,
		terminalWidth:    terminalWidth,
		terminalHeight:   terminalHeight,
	}
}

// animationInterval is the rate at which the message is animated.
const animationInterval = 250 * time.Millisecond

// animationMsg is a tea.Msg for playing the next animation frame.
type animationMsg struct{}

// Init implements the tea.Model interface by returning nil.
func (c SummoningCircle) Init() tea.Cmd {
	return tea.Tick(animationInterval, func(t time.Time) tea.Msg {
		return animationMsg{}
	})
}

// Update implements the tea.Model interface by returning nil.
func (c SummoningCircle) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case animationMsg:
		c.animationFrame = (c.animationFrame + 1) % 4
		cmd := tea.Tick(animationInterval, func(t time.Time) tea.Msg {
			return animationMsg{}
		})
		return c, cmd
	case tea.WindowSizeMsg:
		c.terminalWidth = msg.Width
		c.terminalHeight = msg.Height
	}

	return c, nil
}

// View implements the tea.Model interface by returning a string that displays the summoning circle.
func (c SummoningCircle) View() string {
	var dots string
	for i := 0; i < c.animationFrame; i++ {
		dots += "."
	}

	messageLength := len([]rune(c.summoningMessage))
	style := lipgloss.NewStyle().
		Bold(true).
		PaddingTop(c.terminalHeight / 2).
		PaddingLeft((c.terminalWidth / 2) - (messageLength / 2) - 2)
	return style.Render(c.summoningMessage + dots)
}
