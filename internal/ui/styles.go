package ui

import (
	"github.com/charmbracelet/lipgloss"
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
	Foreground(textColor).
	Width(terminalWidth)

var SecondaryTextStyle = lipgloss.NewStyle().
	Foreground(secondaryTextColor).
	Inherit(BaseStyle)

var InactiveTextStyle = lipgloss.NewStyle().
	Foreground(inactiveTextColor).
	Inherit(BaseStyle)

var FocusedTextStyle = lipgloss.NewStyle().
	Bold(true).
	Inherit(BaseStyle)

var WrappedTextStyle = lipgloss.NewStyle().
	Width(terminalWidth)

var MarginBottomStyle = lipgloss.NewStyle().
	MarginBottom(1).
	Inherit(BaseStyle)

var FullScreenStyle = lipgloss.NewStyle().
	Height(terminalHeight).
	Inherit(BaseStyle)

// UpdateTerminalSize updates the terminal size used when rendering.
func UpdateTerminalSize(w, h int) {
	terminalWidth = w
	terminalHeight = h
	BaseStyle = BaseStyle.Width(w)
	WrappedTextStyle = WrappedTextStyle.Width(w)
	FullScreenStyle = FullScreenStyle.
		Width(w).
		Height(h)
}
