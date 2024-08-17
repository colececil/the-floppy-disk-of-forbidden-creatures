package ui

import (
	"github.com/charmbracelet/x/ansi"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/log"
	"regexp"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

var ansiControlSequenceIntroducer = string([]rune{rune(ansi.ESC), '['})
var ansiResetStyle = ansiControlSequenceIntroducer + "0m"
var ansiInverse = ansiControlSequenceIntroducer + "7m"
var ansiBackgroundStyle, _ = strings.CutSuffix(BackgroundStyle.Render(), ansiResetStyle)

// ansiOrSingleSpaceRegex matches a single ANSI control sequence.
var singleAnsiRegex, _ = regexp.Compile(string(rune(ansi.ESC)) + "\\[(\\d+;)*\\d*[a-zA-Z]")

// ansiOrSingleSpaceRegex matches either 1) a series of one or more ANSI control sequences, or 2) a space. The match
// also includes any surrounding whitespace.
var ansiOrSingleSpaceRegex, _ = regexp.Compile("\\s*((" + singleAnsiRegex.String() + ")+|\\s)\\s*")

// ansiOrSingleSpaceRegex matches either 1) a series of one or more ANSI control sequences, or 2) two spaces. The match
// also includes any surrounding whitespace.
var ansiOrDoubleSpaceRegex, _ = regexp.Compile("\\s*((" + singleAnsiRegex.String() + ")+|\\s{2})\\s*")

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

// PlaceOverlay places the given foreground on top of the given background, without messing up the ANSI styling. If any
// given foreground character is a space or nonexistent, then the corresponding background character is used in that
// location. Otherwise, the foreground character is used.
//
// The transparentSingleSpaces flag is used to indicate that single spaces in the foreground should be transparent. If
// this is set to false, then only two or more consecutive spaces in the foreground are considered transparent.
//
// Note: This function assumes there is no ANSI styling in the given background string. It applies a hard coded style to
// the background as the output is written.
func PlaceOverlay(foreground, background string, transparentSingleSpaces bool) string {
	foregroundLines := strings.Split(foreground, "\n")
	backgroundLines := strings.Split(background, "\n")

	if ansi.StringWidth(backgroundLines[0]) != TerminalWidth || len(backgroundLines) != TerminalHeight {
		log.Logger.Printf("func=\"ui.PlaceOverlay\", msg=\"Background dimensions do not match terminal dimensions.\", "+
			"backgroundWidth=\"%d\", backgroundHeight=\"%d\", terminalWidth=\"%d\", terminalHeight=\"%d\"",
			len(backgroundLines[0]), len(backgroundLines), TerminalWidth, TerminalHeight)
		return ""
	}

	var currentForegroundStyle = new(string)
	var lastUsedStyle = new(string)
	var stringBuilder strings.Builder

	foregroundChunkSeparatorRegex := ansiOrDoubleSpaceRegex
	if transparentSingleSpaces {
		foregroundChunkSeparatorRegex = ansiOrSingleSpaceRegex
	}

	for lineIndex := 0; lineIndex < len(backgroundLines); lineIndex++ {
		var backgroundLine, foregroundLine string
		backgroundLine = backgroundLines[lineIndex]
		if lineIndex < len(foregroundLines) {
			foregroundLine = foregroundLines[lineIndex]
		}

		outputLine := overlaySingleLine(backgroundLine, foregroundLine, currentForegroundStyle, lastUsedStyle,
			foregroundChunkSeparatorRegex)
		stringBuilder.WriteString(outputLine)

		if lineIndex < len(backgroundLines)-1 {
			stringBuilder.WriteString("\n")
		}
	}

	return stringBuilder.String()
}

// overlaySingleLine constructs a single line of the overlay.
func overlaySingleLine(backgroundLine string, foregroundLine string, currentForegroundStyle, lastUsedStyle *string,
	foregroundChunkSeparatorRegex *regexp.Regexp) string {

	if len(foregroundLine) == 0 {
		return textWithStyle(backgroundLine, ansiBackgroundStyle, lastUsedStyle)
	}

	foregroundChunkSeparators := foregroundChunkSeparatorRegex.FindAllStringIndex(foregroundLine, -1)

	if foregroundChunkSeparators == nil || len(foregroundChunkSeparators) == 0 {
		// The foreground has no ANSI control sequences or whitespace, so just write it to the output.
		return textWithStyle(foregroundLine, *currentForegroundStyle, lastUsedStyle)
	} else {
		var stringBuilder strings.Builder
		currentRuneIndex := new(int)

		// If the foreground's first separator is not at the start of the line, we need to write some of the foreground
		// at the start of the line.
		firstSeparatorStartIndex := foregroundChunkSeparators[0][0]
		if firstSeparatorStartIndex > 0 {
			styledText := textWithStyle(foregroundLine[0:firstSeparatorStartIndex], *currentForegroundStyle,
				lastUsedStyle)
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
				lastUsedStyle,
				currentRuneIndex,
			)
			stringBuilder.WriteString(outputChunk)
		}

		// If we are not at the end of the line, we need to write some of the background at the end of the line.
		remainingBackgroundRunes := []rune(backgroundLine)[*currentRuneIndex:]
		if len(remainingBackgroundRunes) > 0 {
			styledText := textWithStyle(string(remainingBackgroundRunes), ansiBackgroundStyle, lastUsedStyle)
			stringBuilder.WriteString(styledText)
		}

		return stringBuilder.String()
	}
}

// overlaySingleLineChunk constructs a single chunk of a single line of the overlay.
func overlaySingleLineChunk(separator, chunk, backgroundLine string, currentForegroundStyle, lastUsedStyle *string,
	currentRuneIndex *int) string {

	// Don't replace a single space with the background character if it's intended to be displayed with an inverse
	// style.
	if strings.Contains(separator, ansiInverse+" ") {
		separator = strings.ReplaceAll(separator, ansiInverse+" ", ansiInverse)
		chunk = " " + chunk
	}

	// If the separator contains spaces, each space needs to be replaced with the corresponding background character.
	spacesInSeparator := strings.Count(separator, " ")
	backgroundRunesToWrite := []rune(backgroundLine)[*currentRuneIndex : *currentRuneIndex+spacesInSeparator]

	// Update the current foreground style if the separator contains a series of one or more ANSI control sequences.
	ansiText := strings.TrimSpace(separator)
	if len(ansiText) > 0 {
		if strings.HasPrefix(ansiText, ansiResetStyle) {
			currentForegroundStyle = new(string)
			ansiText = ansiText[len(ansiResetStyle):]
		}
		ansiControlSequences := singleAnsiRegex.FindAllString(ansiText, -1)
		if ansiControlSequences != nil {
			// Remove any ANSI control sequence that is already in the current foreground style, to avoid duplicates.
			for _, ansiControlSequence := range ansiControlSequences {
				*currentForegroundStyle = strings.ReplaceAll(*currentForegroundStyle, ansiControlSequence, "")
			}
		}
		*currentForegroundStyle += ansiText
	}

	*currentRuneIndex += spacesInSeparator + utf8.RuneCountInString(chunk)

	return textWithStyle(string(backgroundRunesToWrite), ansiBackgroundStyle, lastUsedStyle) +
		textWithStyle(chunk, *currentForegroundStyle, lastUsedStyle)
}

// textWithStyle returns a string with the given text and style.
func textWithStyle(text, style string, lastUsedStyle *string) string {
	if *lastUsedStyle == style {
		return text
	}

	*lastUsedStyle = style
	return ansiResetStyle + style + text
}

// printInvisibleCharacters prints the given text, showing invisible characters as unicode escape sequences. This is
// useful for debugging.
func printInvisibleCharacters(text string) string {
	inputRunes := []rune(text)
	outputRunes := make([]rune, 0)
	for i := 0; i < len(inputRunes); i++ {
		if unicode.IsGraphic(inputRunes[i]) {
			outputRunes = append(outputRunes, inputRunes[i])
		} else {
			outputRunes = append(outputRunes, '\\', 'u')
			codePoint := []rune(strconv.Itoa(int(inputRunes[i])))
			outputRunes = append(outputRunes, codePoint...)
		}
	}
	return string(outputRunes)
}
