package game

import "fmt"

type Game struct {
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Start() {
	fmt.Println(messages[introKey])
}
