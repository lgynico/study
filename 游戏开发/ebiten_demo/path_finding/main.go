package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("Path Finding")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
