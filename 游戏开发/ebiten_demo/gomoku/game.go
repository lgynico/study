package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	chessboard Chessboard
}

func NewGame() *Game {
	return &Game{
		chessboard: NewChessboard(),
	}
}

func (p *Game) Update() error {
	return nil
}

func (p *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.White)
	p.chessboard.Draw(screen)
}

func (p *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
