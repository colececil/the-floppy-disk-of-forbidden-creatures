package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
	"time"
)

// SummoningCircle is a UI component for displaying the summoning circle. It implements the tea.Model interface.
type SummoningCircle struct {
	summoningMessage string
	animationFrame   int
}

// NewSummoningCircle creates a new SummoningCircle.
func NewSummoningCircle(summoningMessage string) SummoningCircle {
	return SummoningCircle{
		summoningMessage: summoningMessage,
	}
}

// animationInterval is the rate at which the message is animated.
const animationInterval = 250 * time.Millisecond

// animationMsg is a tea.Msg for playing the next animation frame.
type animationMsg struct{}

// Init implements the tea.Model interface by returning nil.
func (c SummoningCircle) Init() tea.Cmd {
	return tea.Tick(animationInterval, func(t time.Time) tea.Msg {
		return animationMsg{}
	})
}

// Update implements the tea.Model interface by returning nil.
func (c SummoningCircle) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	if _, ok := msg.(animationMsg); ok {
		c.animationFrame = (c.animationFrame + 1) % 4
		cmd := tea.Tick(animationInterval, func(t time.Time) tea.Msg {
			return animationMsg{}
		})
		return c, cmd
	}

	return c, nil
}

// View implements the tea.Model interface by returning a string that displays the summoning circle.
func (c SummoningCircle) View() string {
	var dots string
	for i := 0; i < c.animationFrame; i++ {
		dots += "."
	}

	// Get width and height of summoning circle ascii art
	asciiArtLines := strings.Split(asciiArt, "\n")
	asciiArtWidth := len([]rune(asciiArtLines[0]))
	asciiArtHeight := len(asciiArtLines)
	text := lipgloss.PlaceHorizontal(asciiArtWidth-2, lipgloss.Center, c.summoningMessage)
	text = strings.TrimRight(text, " ") + dots
	return StyleWithCentering(FocusedTextStyle, asciiArtWidth, asciiArtHeight+4).
		Render(asciiArt + "\n\n" + text)
}

const asciiArt = `                 @@@@@@@@@@@                 
             @@@@           @@@@@            
          @@@     @@@@@@@@@@@   @@           
        @@    @@@   @        @@@   @         
      @@   @@@     @@@          @@   @@      
     @@   @       @@  @@          @@   @@    
   @@   @@        @    @@           @@  @@   
  @   @@         @      @@            @  @@  
 @   @@         @@      @@@@@@@@@@@@@@@@  @@ 
@@   @     @@@@@@@@@@@@@@@@@@@      @@ @  @@ 
@@  @@@@@@@@   @@          @@@    @@   @   @@
@@  @@  @@@@@ @@            @@  @@     @@  @@
@@  @@      @@@@@           @@@@       @@  @@
@@  @@       @@@@@@       @@@  @@      @   @@
@@   @@     @@    @@@@ @@@      @@    @@   @ 
 @@  @@     @       @@@@@@@      @@   @   @@ 
 @@   @@   @@    @@@       @@@@@   @ @@  @@  
  @@   @@@@@ @@@@              @@@@@@   @@   
   @@    @@@@                   @@@   @@     
    @@     @@@                @@     @@      
      @@@     @@@@@@@@@@@ @@@@    @@@        
         @@@                   @@@           
            @@@@@@@@@@@@@@@@@@@              
`
