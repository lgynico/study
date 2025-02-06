package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

type Chessboard struct {
	tiles [15][15]int
}

func NewChessboard() Chessboard {
	return Chessboard{
		tiles: [15][15]int{},
	}
}

func (p *Chessboard) Draw(screen *ebiten.Image) {
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			fmt.Print(p.tiles[i][j])
		}
		fmt.Println()
	}
}
