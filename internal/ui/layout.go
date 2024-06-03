package ui

import (
	"bytes"
	"github.com/charmbracelet/x/ansi"
	"github.com/mattn/go-runewidth"
	"strings"
)

// CenterVertically centers the text vertically within the given height.
func CenterVertically(height int, text string) string {
	textHeight := strings.Count(text, "\n") + 1
	if textHeight >= height {
		return text
	}

	offset := (height - textHeight) / 2
	var b strings.Builder
	for i := 0; i < offset; i++ {
		b.WriteString("\n")
	}
	b.WriteString(text)
	return b.String()
}

// PlaceOverlay places fg on top of bg. This function has been adapted from
// https://github.com/charmbracelet/lipgloss/pull/102/commits/a075bfc9317152e674d661a2cdfe58144306e77a, because Lip
// Gloss does not yet support overlays.
func PlaceOverlay(fg, bg string) string {
	fgLines, fgWidth := getLines(fg)
	bgLines, bgWidth := getLines(bg)
	bgHeight := len(bgLines)
	fgHeight := len(fgLines)

	if fgWidth >= bgWidth && fgHeight >= bgHeight {
		// FIXME: return fg or bg?
		return fg
	}

	var b strings.Builder
	for i, bgLine := range bgLines {
		if i > 0 {
			b.WriteByte('\n')
		}

		if i >= fgHeight {
			b.WriteString(bgLine)
			continue
		}

		fgLine := trimSpacesFromRight(fgLines[i])
		trimmedFgLine := ansi.Strip(fgLine)
		trimmedFgLine = strings.Trim(fgLine, " ")

		if len(trimmedFgLine) == 0 {
			b.WriteString(bgLine)
			continue
		}

		pos := 0

		b.WriteString(fgLine)
		pos += ansi.StringWidth(fgLine)

		right := cutLeft(bgLine, pos)
		bgWidth := ansi.StringWidth(bgLine)
		rightWidth := ansi.StringWidth(right)
		if rightWidth <= bgWidth-pos {
			b.WriteString(strings.Repeat(" ", bgWidth-rightWidth-pos))
		}

		b.WriteString(right)
	}

	return b.String()
}

// trimSpacesFromRight trims spaces from the right, leaving ANSI codes intact.
func trimSpacesFromRight(s string) string {
	strippedString := ansi.Strip(s)
	newLength := len(strings.TrimRight(strippedString, " "))
	return ansi.Truncate(s, newLength, "")
}

// cutLeft cuts printable characters from the left. This function has been copied from
// https://github.com/charmbracelet/lipgloss/pull/102/commits/a075bfc9317152e674d661a2cdfe58144306e77a.
func cutLeft(s string, cutWidth int) string {
	var (
		pos    int
		isAnsi bool
		ab     bytes.Buffer
		b      bytes.Buffer
	)
	for _, c := range s {
		var w int
		if c == ansi.ESC || isAnsi {
			isAnsi = true
			ab.WriteRune(c)
			if IsTerminator(c) {
				isAnsi = false
				if bytes.HasSuffix(ab.Bytes(), []byte("[0m")) {
					ab.Reset()
				}
			}
		} else {
			w = runewidth.RuneWidth(c)
		}

		if pos >= cutWidth {
			if b.Len() == 0 {
				if ab.Len() > 0 {
					b.Write(ab.Bytes())
				}
				if pos-cutWidth > 1 {
					b.WriteByte(' ')
					continue
				}
			}
			b.WriteRune(c)
		}
		pos += w
	}
	return b.String()
}

// This function has been copied from
// https://github.com/charmbracelet/lipgloss/pull/102/commits/a075bfc9317152e674d661a2cdfe58144306e77a.
func clamp(v, lower, upper int) int {
	return min(max(v, lower), upper)
}

// Split a string into lines, additionally returning the size of the widest line. This function has been copied from
// Lip Gloss.
func getLines(s string) (lines []string, widest int) {
	lines = strings.Split(s, "\n")

	for _, l := range lines {
		w := ansi.StringWidth(l)
		if widest < w {
			widest = w
		}
	}

	return lines, widest
}

// IsTerminator checks if the given rune is a terminator. This function has been copied from muesli/reflow.
func IsTerminator(c rune) bool {
	return (c >= 0x40 && c <= 0x5a) || (c >= 0x61 && c <= 0x7a)
}
