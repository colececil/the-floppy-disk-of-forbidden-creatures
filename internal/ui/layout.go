package ui

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/log"
	"regexp"
	"strings"
	"unicode/utf8"
)

var ansiControlSequenceIntroducer = string([]rune{rune(ansi.ESC), '['})
var ansiResetStyle = ansiControlSequenceIntroducer + "0m"
var ansiBackgroundStyle, _ = strings.CutSuffix(BackgroundStyle.Render(), ansiResetStyle)

var ansiAndWhitespaceRegex, _ = regexp.Compile("\\s*(" + string(rune(ansi.ESC)) + "\\[(\\d+;)*\\d*[a-zA-Z]|\\s+)\\s*")

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

// PlaceOverlay places the foreground on top of the background, without messing up the ANSI styling. If any given
// foreground character is a space or nonexistent, then the corresponding background character is used in that location.
// Otherwise, the foreground character is used.
//
// Note: This function assumes there is no ANSI styling in the given background string. It applies a hard coded style to
// the background as the output is written.
func PlaceOverlay(foreground, background string) string {
	foregroundLines := strings.Split(foreground, "\n")
	backgroundLines := strings.Split(background, "\n")

	if ansi.StringWidth(backgroundLines[0]) != TerminalWidth || len(backgroundLines) != TerminalHeight {
		log.Logger.Printf("func=\"ui.PlaceOverlay\", msg=\"Background dimensions do not match terminal dimensions.\", "+
			"backgroundWidth=\"%d\", backgroundHeight=\"%d\", terminalWidth=\"%d\", terminalHeight=\"%d\"",
			len(backgroundLines[0]), len(backgroundLines), TerminalWidth, TerminalHeight)
		return ""
	}

	var currentForegroundStyle = new(string)
	var stringBuilder strings.Builder

	for lineIndex := 0; lineIndex < len(backgroundLines); lineIndex++ {
		var backgroundLine, foregroundLine string
		backgroundLine = backgroundLines[lineIndex]
		if lineIndex < len(foregroundLines) {
			foregroundLine = foregroundLines[lineIndex]
		}

		outputLine := overlaySingleLine(backgroundLine, foregroundLine, currentForegroundStyle)
		stringBuilder.WriteString(outputLine)

		if lineIndex < len(backgroundLines)-1 {
			stringBuilder.WriteString("\n")
		}
	}

	return stringBuilder.String()
}

// overlaySingleLine constructs a single line of the overlay.
func overlaySingleLine(backgroundLine string, foregroundLine string, currentForegroundStyle *string) string {
	if len(foregroundLine) == 0 {
		return textWithStyle(backgroundLine, ansiBackgroundStyle)
	}

	foregroundChunkSeparators := ansiAndWhitespaceRegex.FindAllStringIndex(foregroundLine, -1)

	if foregroundChunkSeparators == nil || len(foregroundChunkSeparators) == 0 {
		// The foreground has no ANSI control sequences or whitespace, so just write it to the output.
		return textWithStyle(foregroundLine, *currentForegroundStyle)
	} else {
		var stringBuilder strings.Builder
		currentRuneIndex := new(int)

		// If the foreground's first separator is not at the start of the line, we need to write some of the foreground
		// at the start of the line.
		firstSeparatorStartIndex := foregroundChunkSeparators[0][0]
		if firstSeparatorStartIndex > 0 {
			styledText := textWithStyle(foregroundLine[0:firstSeparatorStartIndex], *currentForegroundStyle)
			stringBuilder.WriteString(styledText)
			*currentRuneIndex = utf8.RuneCountInString(foregroundLine[0:firstSeparatorStartIndex])
		}

		for separatorResultIndex := 0; separatorResultIndex < len(foregroundChunkSeparators); separatorResultIndex++ {
			separatorStartIndex := foregroundChunkSeparators[separatorResultIndex][0]
			separatorEndIndex := foregroundChunkSeparators[separatorResultIndex][1]
			chunkEndIndex := len(foregroundLine)
			if separatorResultIndex+1 < len(foregroundChunkSeparators) {
				chunkEndIndex = foregroundChunkSeparators[separatorResultIndex+1][0]
			}

			outputChunk := overlaySingleLineChunk(
				foregroundLine[separatorStartIndex:separatorEndIndex],
				foregroundLine[separatorEndIndex:chunkEndIndex],
				backgroundLine,
				currentForegroundStyle,
				currentRuneIndex,
			)
			stringBuilder.WriteString(outputChunk)
		}

		// If we are not at the end of the line, we need to write some of the background at the end of the line.
		remainingBackgroundRunes := []rune(backgroundLine)[*currentRuneIndex:]
		if len(remainingBackgroundRunes) > 0 {
			styledText := textWithStyle(string(remainingBackgroundRunes), ansiBackgroundStyle)
			stringBuilder.WriteString(styledText)
		}

		return stringBuilder.String()
	}
}

// overlaySingleLineChunk constructs a single chunk of a single line of the overlay.
func overlaySingleLineChunk(separator, chunk, backgroundLine string, currentForegroundStyle *string,
	currentRuneIndex *int) string {

	// If the separator contains spaces, each space needs to be replaced with the corresponding background character.
	spacesInSeparator := strings.Count(separator, " ")
	backgroundRunesToWrite := []rune(backgroundLine)[*currentRuneIndex : *currentRuneIndex+spacesInSeparator]

	// Update the current foreground style if the separator contains an ANSI control sequence.
	ansiControlSequence := strings.TrimSpace(separator)
	if len(ansiControlSequence) > 0 {
		if strings.HasPrefix(ansiControlSequence, ansiResetStyle) {
			currentForegroundStyle = new(string)
			ansiControlSequence = ansiControlSequence[len(ansiResetStyle):]
		}
		*currentForegroundStyle = strings.ReplaceAll(*currentForegroundStyle, ansiControlSequence, "")
		*currentForegroundStyle += ansiControlSequence
	}

	*currentRuneIndex += spacesInSeparator + utf8.RuneCountInString(chunk)

	return textWithStyle(string(backgroundRunesToWrite), ansiBackgroundStyle) +
		textWithStyle(chunk, *currentForegroundStyle)
}

// textWithStyle returns a string with the given text and style.
func textWithStyle(text, style string) string {
	return ansiResetStyle + style + text
}
