package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	panic("never called")
}

func (g *Game) LayoutF(outsideWidth float64, outsideHeight float64) (screenWidth float64, screenHeight float64) {
	return 1920, 1080
}
