package ui

import (
	"github.com/charmbracelet/x/ansi"
	"strings"
	"unicode/utf8"
)

const ansiResetStyle = "0m"

var ansiControlSequenceIntroducer = string([]rune{rune(ansi.ESC), '['})

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

// PlaceOverlay places the foreground on top of the background, without messing up the ANSI styling. It does
// this by comparing each line character-by-character. If the foreground character is a space or nonexistent, then the
// background character is used. Otherwise, the foreground character is used. Whenever switching between foreground and
// background characters, the ANSI styling is reset and the current styling of the layer being switched to is applied.
func PlaceOverlay(foreground, background string) string {
	foregroundLines := strings.Split(foreground, "\n")
	backgroundLines := strings.Split(background, "\n")
	width := len(backgroundLines[0])
	height := len(backgroundLines)

	var currentStyle = new(string)
	var currentForegroundStyle = new(string)
	var currentBackgroundStyle = new(string)

	var stringBuilder strings.Builder

	for lineIndex := 0; lineIndex < height; lineIndex++ {
		if lineIndex >= len(foregroundLines) {
			if *currentStyle != *currentBackgroundStyle {
				stringBuilder.WriteString(resetToStyle(currentBackgroundStyle))
			}
			stringBuilder.WriteString(backgroundLines[lineIndex] + "\n")
			continue
		}

		foregroundIndex := 0
		backgroundIndex := 0
		for charIndex := 0; charIndex < width; charIndex++ {
			var foregroundCharacter, backgroundCharacter rune
			foregroundCharacter, foregroundIndex =
				getNextPrintableCharacter(foregroundLines[lineIndex], foregroundIndex, currentForegroundStyle)
			backgroundCharacter, backgroundIndex =
				getNextPrintableCharacter(backgroundLines[lineIndex], backgroundIndex, currentBackgroundStyle)

			if foregroundCharacter != ' ' { // Write foreground character
				if *currentStyle != *currentForegroundStyle {
					stringBuilder.WriteString(resetToStyle(currentForegroundStyle))
				}
				stringBuilder.WriteString(string(foregroundCharacter))
			} else { // Write background character
				if *currentStyle != *currentBackgroundStyle {
					stringBuilder.WriteString(resetToStyle(currentBackgroundStyle))
				}
				stringBuilder.WriteString(string(backgroundCharacter))
			}

			foregroundIndex = foregroundIndex + utf8.RuneLen(foregroundCharacter)
			backgroundIndex = backgroundIndex + utf8.RuneLen(backgroundCharacter)
		}
		stringBuilder.WriteString("\n")
	}
	return stringBuilder.String()
}

// getNextPrintableCharacter returns the next printable character, parsing any ANSI escape sequences along the way and
// modifying the current style as necessary. It returns the next printable character and the index of that character. If
// the given index is out of range, it returns the space character and the given index.
func getNextPrintableCharacter(s string, startIndex int, currentStyle *string) (rune, int) {
	if startIndex >= len(s) {
		return ' ', startIndex
	}

	i := startIndex
	for strings.HasPrefix(s[i:], ansiControlSequenceIntroducer) {
		controlSequence := ansiControlSequenceIntroducer
		i += len(ansiControlSequenceIntroducer)
		for {
			nextRune, runeSize := utf8.DecodeRuneInString(s[i:])
			controlSequence += string(nextRune)
			i += runeSize
			if (nextRune >= 'a' && nextRune <= 'z') || (nextRune >= 'A' && nextRune <= 'Z') {
				if controlSequence == ansiControlSequenceIntroducer+ansiResetStyle {
					*currentStyle = ""
				} else {
					*currentStyle += controlSequence
				}
				break
			}
		}
	}

	if i >= len(s) {
		return ' ', i
	}

	nextRune, _ := utf8.DecodeRuneInString(s[i:])
	return nextRune, i
}

// resetToStyle returns a string that resets the ANSI styling to the given style.
func resetToStyle(style *string) string {
	var b strings.Builder
	b.WriteString(ansiControlSequenceIntroducer + ansiResetStyle)
	if style != nil {
		b.WriteString(*style)
	}
	return b.String()
}
