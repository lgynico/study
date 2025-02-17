package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Chessboard struct {
	tiles [TileCount][TileCount]int
}

func NewChessboard() Chessboard {
	return Chessboard{
		tiles: [TileCount][TileCount]int{},
	}
}

func (p *Chessboard) Draw(screen *ebiten.Image) {
	for i := 0; i < TileCount; i++ {
		var (
			y  = float32(i)*TileHeight + TileHeight/2
			x1 = float32(0)*TileWidth + TileWidth/2
			x2 = float32(14)*TileWidth + TileWidth/2
		)

		vector.StrokeLine(screen, x1, y, x2, y, 1, color.Black, false)
	}

	for i := 0; i < TileCount; i++ {
		var (
			x  = float32(i)*TileWidth + TileWidth/2
			y1 = float32(0)*TileHeight + TileHeight/2
			y2 = float32(14)*TileHeight + TileHeight/2
		)

		vector.StrokeLine(screen, x, y1, x, y2, 1, color.Black, false)
	}

	vector.DrawFilledCircle(
		screen,
		TileCount/2*TileWidth+TileWidth/2,
		TileCount/2*TileHeight+TileHeight/2,
		5,
		color.Black,
		true,
	)

}
