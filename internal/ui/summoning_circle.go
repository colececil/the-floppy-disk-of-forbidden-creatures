package ui

import tea "github.com/charmbracelet/bubbletea"

// SummoningCircle is a UI component for displaying the summoning circle. It implements the tea.Model interface.
type SummoningCircle struct {
}

// NewSummoningCircle creates a new SummoningCircle.
func NewSummoningCircle() SummoningCircle {
	return SummoningCircle{}
}

// Init implements the tea.Model interface by returning nil.
func (c *SummoningCircle) Init() tea.Cmd {
	return nil
}

// Update implements the tea.Model interface by returning nil.
func (c *SummoningCircle) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return c, nil
}

// View implements the tea.Model interface by returning a string that displays the summoning circle.
func (c *SummoningCircle) View() string {
	return "...Summoning Circle..."
}
