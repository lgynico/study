package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type HelloWorld struct {
	display bool
}

func (h *HelloWorld) Update() error {

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		h.display = false
		fmt.Println("Pressed")
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		h.display = true
		fmt.Println("Released")
	}

	return nil
}

func (h *HelloWorld) Draw(screen *ebiten.Image) {
	if h.display {
		f, err := os.Open("assets/image/gopher.jpg")
		if err != nil {
			log.Fatal(err)
		}

		img, err := jpeg.Decode(f)
		if err != nil {
			log.Fatal(err)
		}

		image := ebiten.NewImageFromImage(img)
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(100, 100)
		screen.DrawImage(image, opts)
	} else {
		screen.Clear()
	}

	ebitenutil.DebugPrint(screen, "Hello World!")
}

func (h *HelloWorld) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	screenWidth = 640
	screenHeight = 480
	return
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello World")

	app := &HelloWorld{
		display: true,
	}

	if err := ebiten.RunGame(app); err != nil {
		log.Fatal(err)
	}
}
