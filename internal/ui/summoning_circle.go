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

	view := lipgloss.PlaceHorizontal(terminalWidth, lipgloss.Center, asciiArt)
	text := lipgloss.PlaceHorizontal(terminalWidth, lipgloss.Center, c.summoningMessage)
	text = FocusedTextStyle.
		Width(terminalWidth).
		Height(1).
		MarginTop(1).
		Render(strings.TrimRight(text, " ") + dots)
	view = lipgloss.JoinVertical(lipgloss.Left, view, text)
	view = lipgloss.PlaceVertical(terminalHeight, lipgloss.Center, view,
		lipgloss.WithWhitespaceBackground(backgroundColor))
	return BaseStyle.Render(view)
}

const asciiArt = `                 @@@@@@@@@@@                 
             @@@@           @@@@@            
          @@@     @@@@@@@@@@@   @@           
        @@    @@@   A        @@@   @         
      @@   @@@     4^B.         @@   @@      
     @@   @       47  lb          @@   @@    
   @@   @@       d$    7b           @@   @   
  @   @@        12      Bb            @   @  
 @   @@        .8[   ____>L===oo888PPP7@   @ 
@@   @<=ooooood8999PPP***'<8\      ,0"  @  @ 
@@  @@ '&i_   /8'          '8.   ,0"    @  @@
@@  @@   '"b_.$'            ]&_,0"      @  @@
@@  @@     ''8b.           .,80[        @  @@
@@  @@      !9'"&._     .,8%'  lb      @   @@
@@   @@     8!    '"&>,8%'      7b     @   @ 
 @@  @@    !8       ,%"<9b._     Bb   @   @@ 
 @@   @@   1!    ,8%'     '"7&i,_ jb @@  @@  
  @@   @@ !8 .,8%'             '""liJ   @@   
   @@    \L4%"                   @@   @@     
    @@     @@                 @@     @@      
      @@@    @@@@@@@@@@@@ @@@@    @@@        
         @@@                   @@@           
            @@@@@@@@@@@@@@@@@@@              `
