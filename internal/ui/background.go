package ui

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/colececil/the-floppy-disk-of-forbidden-creatures/internal/log"
	"math/rand/v2"
	"strings"
	"time"
)

var availableCharacterTypes = [][]rune{
	{'.', ',', '`'},
	{'?', '¿'},
	{'o'},
	{'Æ', 'À', 'Á'},
	{'#'},
}

// Background is a UI component for displaying the background. It implements the tea.Model interface.
type Background struct {
	characterTypes     [][]int
	characters         [][]rune
	currentAnimationId int
	currentWidth       int
	currentHeight      int
}

// NewBackground returns a new Background.
func NewBackground() Background {
	return Background{}
}

// backgroundAnimationInterval is the rate at which the background is animated.
const backgroundAnimationInterval = 100 * time.Millisecond

// characterTypesUpdateMsg is a tea.Msg that updates the characterTypes of the background.
type characterTypesUpdateMsg struct {
	animationId    int
	characterTypes [][]int
	characters     [][]rune
}

// Init implements tea.Model by returning a tea.Cmd that initializes the characterTypes slice.
func (b Background) Init() tea.Cmd {
	b.currentWidth = TerminalWidth
	b.currentHeight = TerminalHeight
	return func() tea.Msg {
		return initializeCharacterTypes(b.currentAnimationId, TerminalWidth, TerminalHeight)
	}
}

// Update implements tea.Model by updating the model based on the given message.
func (b Background) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case characterTypesUpdateMsg:
		if msg.animationId == b.currentAnimationId {
			b.characterTypes = msg.characterTypes
			b.characters = msg.characters
			cmd = tea.Tick(backgroundAnimationInterval, func(t time.Time) tea.Msg {
				return generateNextStep(msg.animationId, msg.characterTypes, msg.characters)
			})
		}
	case tea.WindowSizeMsg:
		if msg.Width != b.currentWidth || msg.Height != b.currentHeight {
			log.Logger.Print(fmt.Sprintf("func=\"ui.Background.Update\", msg=\"Window size updated.\", width=\"%d\", "+
				"height=\"%d\"", msg.Width, msg.Height))
			b.currentAnimationId++
			b.currentWidth = msg.Width
			b.currentHeight = msg.Height
			cmd = func() tea.Msg {
				return initializeCharacterTypes(b.currentAnimationId, msg.Width, msg.Height)
			}
		}
	}

	return b, cmd
}

// View implements tea.Model by returning a string that displays the background.
func (b Background) View() string {
	var stringBuilder strings.Builder
	for i, line := range b.characters {
		stringBuilder.WriteString(string(line))
		if i < len(b.characters)-1 {
			stringBuilder.WriteString("\n")
		}
	}
	return stringBuilder.String()
}

// initializeCharacterTypes initializes the character types to show in the background and returns them in a
// characterTypesUpdateMsg.
func initializeCharacterTypes(animationId, width, height int) tea.Msg {
	log.Logger.Print(fmt.Sprintf("func=\"ui.Background.initializeCharacterTypes\", msg=\"Function called.\", "+
		"animationId=\"%d\", width=\"%d\", height=\"%d\"", animationId, width, height))

	characterTypes := make([][]int, height)
	for i := range characterTypes {
		characterTypes[i] = make([]int, width)
		for j := range characterTypes[i] {
			characterTypes[i][j] = randomCharacterType()
		}
	}
	return characterTypesUpdateMsg{
		animationId:    animationId,
		characterTypes: characterTypes,
		characters:     generateCharactersForTypes(characterTypes, nil, nil),
	}
}

// generateNextStep generates the next step of the cellular automaton.
func generateNextStep(animationId int, characterTypes [][]int, characters [][]rune) tea.Msg {
	nextCharacterTypes := make([][]int, len(characterTypes))
	for i := range characterTypes {
		nextCharacterTypes[i] = make([]int, len(characterTypes[i]))
		for j := range characterTypes[i] {
			nextCharacterTypes[i][j] = nextCharacterType(characterTypes[i][j], getNeighbors(i, j, characterTypes))
		}
	}
	return characterTypesUpdateMsg{
		animationId:    animationId,
		characterTypes: nextCharacterTypes,
		characters:     generateCharactersForTypes(nextCharacterTypes, characterTypes, characters),
	}
}

// nextCharacterType returns the next character type by randomly selecting either the given character type or one of
// its neighbors.
func nextCharacterType(characterType int, neighbors []int) int {
	neighbors = append(neighbors, characterType)
	return neighbors[rand.IntN(len(neighbors))]
}

// getNeighbors returns the horizontal, vertical, and diagonal neighbors of the given cell.
func getNeighbors(i, j int, characterTypes [][]int) []int {
	var neighbors []int
	for deltaI := -1; deltaI <= 1; deltaI++ {
		for deltaJ := -1; deltaJ <= 1; deltaJ++ {
			if deltaI == 0 && deltaJ == 0 {
				continue
			}
			neighborI := i + deltaI
			neighborJ := j + deltaJ
			if neighborI >= 0 && neighborI < len(characterTypes) &&
				neighborJ >= 0 && neighborJ < len(characterTypes[neighborI]) {

				neighbors = append(neighbors, characterTypes[neighborI][neighborJ])
			}
		}
	}
	return neighbors
}

// generateCharactersForTypes generates characters based on the given character types. If the character type in a given
// cell hasn't changed since the previous step, the character in that cell remains the same. Otherwise, a random
// character of that type is chosen.
func generateCharactersForTypes(characterTypes [][]int, previousCharacterTypes [][]int,
	previousCharacters [][]rune) [][]rune {

	characters := make([][]rune, len(characterTypes))
	for i, line := range characterTypes {
		characters[i] = make([]rune, len(line))
		for j, characterType := range line {
			if previousCharacterTypes != nil && previousCharacterTypes[i][j] == characterType {
				characters[i][j] = previousCharacters[i][j]
			} else {
				characters[i][j] = randomCharacterOfType(characterType)
			}
		}
	}
	return characters
}

// randomCharacterType returns a random character type.
func randomCharacterType() int {
	return rand.IntN(len(availableCharacterTypes))
}

// randomCharacterOfType returns a random character of the given type.
func randomCharacterOfType(characterType int) rune {
	numCharactersOfType := len(availableCharacterTypes[characterType])
	return availableCharacterTypes[characterType][rand.IntN(numCharactersOfType)]
}
