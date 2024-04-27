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

const asciiArt = "            ,,ooo000000000000ooo,,            " +
	"        ,oP''                    ''^o.        " +
	"      ,d\"     ,,ooo00000000ooo,,     \"b.      " +
	"    .&'    ,dP'    A           '^b.    '&.    " +
	"   .&    dP\"      d^V,            \"^b    &.   " +
	"  d&   .&'       ,8  V,             '&.   &b  " +
	" ,P   .&         d!   !b              &.   ^. " +
	" 8   .&         ,8     'V,             &.   8 " +
	".8   8          d!   ____V,===oo888PP87 8   8." +
	"8   .8<=ooooood8999PPP***'V,       ,0'  8.   8" +
	"8   8' '&i_    d!          !b   ,0\"'    '8   8" +
	"8   8    '\"b_ ,8            'V~7'        8   8" +
	"8   8      ''8!!           ,0\"0,         8   8" +
	"8   8.       ,8\"&._      ,0\"   V,       .8   8" +
	"8   `8       d!   '\"&>,~7'      !b      8'   8" +
	"`8   8      ,8      ,0\"<9b._     'V,    8   8'" +
	" 8   `8     d!    ,0\"     '\"7&i,_  V,  d'   8 " +
	" `8   `b   ,8  ,~7'            '\"\"liJ,d'   8' " +
	"  `b   `b. !',0\"                    .d'   d'  " +
	"   `b    *bL%'                    .d*    d'   " +
	"    `*.    `*b.                .d*'    .*'    " +
	"      `*b.    ``***00000000***''    .d*'      " +
	"         `*b..                  ..d*'         " +
	"             `***000000000000***'             "
