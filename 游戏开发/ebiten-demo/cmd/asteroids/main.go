package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lgynico/ebiten-demo/app/asteroids"
)

func main() {
	g := asteroids.NewGame()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
