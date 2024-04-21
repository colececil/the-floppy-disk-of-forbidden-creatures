package ui

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/reflow/wordwrap"
)

const (
	backgroundColor    = lipgloss.Color("#15011B")
	textColor          = lipgloss.Color("#FFFFFF")
	secondaryTextColor = lipgloss.Color("#FF8888")
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

// StyleWithCentering returns a new style based on the given style, with the content centered in the terminal. The given
// content width and height are used to calculate the padding.
func StyleWithCentering(style lipgloss.Style, contentWidth int, contentHeight int) lipgloss.Style {
	return style.Copy().
		PaddingTop((terminalHeight / 2) - (contentHeight / 2)).
		PaddingLeft((terminalWidth / 2) - (contentWidth / 2)).
		Inherit(BaseStyle)
}
