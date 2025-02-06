package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	opt        Option
	chessboard Chessboard
}

func NewGame(opts ...OptionFunc) *Game {
	g := &Game{
		opt:        DefaultOption(),
		chessboard: NewChessboard(),
	}

	for _, optFunc := range opts {
		optFunc(&g.opt)
	}

	return g
}

func (p *Game) Update() error {
	return nil
}

func (p *Game) Draw(screen *ebiten.Image) {
	p.chessboard.Draw(screen)
}

func (p *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return p.opt.ScreenWidth, p.opt.ScreenHeight
}
