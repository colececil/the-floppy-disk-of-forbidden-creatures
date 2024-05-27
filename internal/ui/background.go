package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"math/rand/v2"
	"time"
)

var availableCharacterTypes = [][]rune{
	{'$'},
	{'Y'},
	{'f'},
	{'{'},
	{'0'},
}

// Background is a UI component for displaying the background. It implements the tea.Model interface.
type Background struct {
	characterTypes [][]int
	characters     string
	currentWidth   int
	currentHeight  int
}

// NewBackground returns a new Background.
func NewBackground() Background {
	return Background{}
}

// backgroundAnimationInterval is the rate at which the background is animated.
const backgroundAnimationInterval = 100 * time.Millisecond

// characterTypesUpdateMsg is a tea.Msg that updates the characterTypes of the background.
type characterTypesUpdateMsg struct {
	characterTypes [][]int
	characters     string
}

// Init implements tea.Model by returning a tea.Cmd that initializes the characterTypes slice.
func (b Background) Init() tea.Cmd {
	return func() tea.Msg { return initializeCharacterTypes(terminalWidth, terminalHeight) }
}

// Update implements tea.Model by updating the model based on the given message.
func (b Background) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case characterTypesUpdateMsg:
		if len(msg.characterTypes) == b.currentHeight &&
			len(msg.characterTypes) >= 1 &&
			len(msg.characterTypes[0]) == b.currentWidth {

			b.characterTypes = msg.characterTypes
			b.characters = msg.characters
			cmd = tea.Tick(backgroundAnimationInterval, func(t time.Time) tea.Msg {
				return generateNextStep(msg.characterTypes)
			})
		}
	case tea.WindowSizeMsg:
		b.currentWidth = msg.Width
		b.currentHeight = msg.Height
		cmd = func() tea.Msg { return initializeCharacterTypes(msg.Width, msg.Height) }
	}

	return b, cmd
}

// View implements tea.Model by returning a string that displays the background.
func (b Background) View() string {
	return BackgroundStyle.Render(b.characters)
}

// initializeCharacterTypes initializes the character types to show in the background and returns them in a
// characterTypesUpdateMsg.
func initializeCharacterTypes(width int, height int) tea.Msg {
	characterTypes := make([][]int, height)
	for i := range characterTypes {
		characterTypes[i] = make([]int, width)
		for j := range characterTypes[i] {
			characterTypes[i][j] = randomCharacterType()
		}
	}
	return characterTypesUpdateMsg{
		characterTypes: characterTypes,
		characters:     generateCharactersForTypes(characterTypes),
	}
}

// generateNextStep generates the next step of the cellular automaton.
func generateNextStep(characterTypes [][]int) tea.Msg {
	nextCharacterTypes := make([][]int, len(characterTypes))
	for i := range characterTypes {
		nextCharacterTypes[i] = make([]int, len(characterTypes[i]))
		for j := range characterTypes[i] {
			nextCharacterTypes[i][j] = nextCharacterType(characterTypes[i][j], getNeighbors(i, j, characterTypes))
		}
	}
	return characterTypesUpdateMsg{
		characterTypes: nextCharacterTypes,
		characters:     generateCharactersForTypes(nextCharacterTypes),
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

// generateCharactersForTypes generates a random string of characters based on the given character types.
func generateCharactersForTypes(characterTypes [][]int) string {
	var characters string
	for i, line := range characterTypes {
		for _, characterType := range line {
			characters += string(randomCharacterOfType(characterType))
		}
		if i < len(characterTypes)-1 {
			characters += "\n"
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
