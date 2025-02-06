package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	if err := ebiten.RunGame(NewGame()); err != nil {
		panic(err)
	}
}
