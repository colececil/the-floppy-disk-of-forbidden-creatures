package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var asciiArt string

// init loads the summoning circle ASCII art from the assets directory.
func init() {
	var err error
	asciiArt, err = loadAsciiArt()
	if err != nil {
		panic(err)
	}
}

// SummoningCircle is a UI component for displaying the summoning circle. It implements the tea.Model interface.
type SummoningCircle struct {
	summoningMessage string
	animationFrame   int
}

// NewSummoningCircle creates a new SummoningCircle.
func NewSummoningCircle(summoningMessage string) SummoningCircle {
	return SummoningCircle{
		summoningMessage: summoningMessage,
	}
}

// summoningCircleAnimationInterval is the rate at which the message is animated.
const summoningCircleAnimationInterval = 500 * time.Millisecond

// animationMsg is a tea.Msg for playing the next animation frame.
type animationMsg struct{}

// Init implements the tea.Model interface by returning nil.
func (c SummoningCircle) Init() tea.Cmd {
	return tea.Tick(summoningCircleAnimationInterval, func(t time.Time) tea.Msg {
		return animationMsg{}
	})
}

// Update implements the tea.Model interface by returning nil.
func (c SummoningCircle) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(animationMsg); ok {
		c.animationFrame = (c.animationFrame + 1) % 6
		cmd := tea.Tick(summoningCircleAnimationInterval, func(t time.Time) tea.Msg {
			return animationMsg{}
		})
		return c, cmd
	}

	return c, nil
}

// View implements the tea.Model interface by returning a string that displays the summoning circle.
func (c SummoningCircle) View() string {
	view := lipgloss.PlaceHorizontal(TerminalWidth, lipgloss.Center, asciiArt)

	var dots string
	numDots := c.animationFrame
	if c.animationFrame > 3 {
		numDots = 6 - c.animationFrame
	}
	for i := 0; i < numDots; i++ {
		dots += "."
	}
	text := lipgloss.PlaceHorizontal(TerminalWidth, lipgloss.Center, c.summoningMessage)
	text = strings.TrimRight(text, " ") + dots
	text = lipgloss.NewStyle().
		Bold(true).
		MarginTop(1).
		Render(text)

	view = lipgloss.JoinVertical(lipgloss.Left, view, text)
	view = PrimaryTextStyle.Render(view)
	return CenterVertically(TerminalHeight, view)
}

// loadAsciiArt loads the summoning circle ASCII art from the assets directory.
func loadAsciiArt() (string, error) {
	pathToExecutable, err := os.Executable()
	if err != nil {
		return "", err
	}
	dirOfExecutable := filepath.Dir(pathToExecutable)

	path := filepath.Join(dirOfExecutable, "assets/ascii_art/summoning_circle.txt")
	bytes, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
