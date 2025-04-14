package asteroids

import (
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/lgynico/ebiten-demo/common/types"
)

var MeteorSprites = mustLoadImages("assets/PNG/Meteors/*.png")

type Meteor struct {
	position types.Vector
	sprite   *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := MeteorSprites[rand.IntN(len(MeteorSprites))]
	return &Meteor{
		position: types.Vector{},
		sprite:   sprite,
	}
}

func (m *Meteor) Update() {
}

func (m *Meteor) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(m.position.X, m.position.Y)
	screen.DrawImage(m.sprite, op)
}
