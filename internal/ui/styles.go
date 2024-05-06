package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

const (
	backgroundColor    = lipgloss.Color("#0F0114")
	textColor          = lipgloss.Color("#FFFFFF")
	secondaryTextColor = lipgloss.Color("#FF2626")
	inactiveTextColor  = lipgloss.Color("#6A4D4D")
)

var terminalWidth, terminalHeight int

var BaseStyle = lipgloss.NewStyle().
	Background(backgroundColor).
	Foreground(textColor)

var SecondaryTextStyle = lipgloss.NewStyle().
	Foreground(secondaryTextColor).
	Inherit(BaseStyle)

var InactiveTextStyle = lipgloss.NewStyle().
	Foreground(inactiveTextColor).
	Inherit(BaseStyle)

var FocusedTextStyle = lipgloss.NewStyle().
	Foreground(textColor).
	Bold(true).
	Inherit(BaseStyle)

// UpdateTerminalSize updates the terminal size used when rendering.
func UpdateTerminalSize(w, h int) {
	terminalWidth = w
	terminalHeight = h
	BaseStyle.
		Width(w).
		Height(h)
}

// WrapText wraps the given text based on the terminal width.
func WrapText(text string) string {
	return wordwrap.String(text, terminalWidth)
}
