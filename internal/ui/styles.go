package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	backgroundColor          = lipgloss.Color("#0F0114")
	backgroundAnimationColor = lipgloss.Color("#3A042B")
	textColor                = lipgloss.Color("#FFFFFF")
	secondaryTextColor       = lipgloss.Color("#FF2626")
	inactiveTextColor        = lipgloss.Color("#6A4D4D")
)

var TerminalWidth int
var TerminalHeight int

var BackgroundStyle = lipgloss.NewStyle().
	Background(backgroundColor).
	Foreground(backgroundAnimationColor)

var PrimaryTextStyle = BackgroundStyle.Foreground(textColor)

var SecondaryTextStyle = BackgroundStyle.Foreground(secondaryTextColor)

var InactiveTextStyle = BackgroundStyle.Foreground(inactiveTextColor)

var FullScreenStyle = BackgroundStyle.
	Width(TerminalWidth).
	Height(TerminalHeight)

// UpdateTerminalSize updates the terminal size used when rendering.
func UpdateTerminalSize(w, h int) {
	TerminalWidth = w
	TerminalHeight = h
	FullScreenStyle = FullScreenStyle.
		Width(w).
		Height(h)
}
