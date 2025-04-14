package asteroids

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lgynico/ebiten-demo/common/types"
)

var PlayerSpirte = mustLoadImage("assets/PNG/playerShip1_blue.png")

type Player struct {
	position types.Vector
	rotation float64
	sprite   *ebiten.Image

	halfW float64
	halfH float64
}

func NewPlayer() *Player {
	var (
		sprite = PlayerSpirte
		bounds = sprite.Bounds()
		halfW  = float64(bounds.Dx()) / 2
		halfH  = float64(bounds.Dy()) / 2
		pos    = types.Vector{
			X: ScreenWidth/2 - halfW,
			Y: ScreenHeight/2 - halfH,
		}
	)

	return &Player{
		position: pos,
		sprite:   sprite,
		halfW:    halfW,
		halfH:    halfH,
	}
}

func (p *Player) Udpate() {
	speed := math.Pi / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.rotation -= speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.rotation += speed
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(-p.halfW, -p.halfH)
	op.GeoM.Rotate(p.rotation)
	op.GeoM.Translate(p.halfW, p.halfH)

	op.GeoM.Translate(p.position.X, p.position.Y)

	screen.DrawImage(p.sprite, op)
}
