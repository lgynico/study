package asteroids

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	player           *Player
	meteors          []*Meteor
	meteorSpawnTimer *Timer
}

func NewGame() *Game {
	return &Game{
		player:           NewPlayer(),
		meteors:          make([]*Meteor, 0),
		meteorSpawnTimer: NewTimer(time.Second * 5),
	}
}

func (g *Game) Update() error {
	g.player.Udpate()

	g.meteorSpawnTimer.Update()
	if g.meteorSpawnTimer.IsReady() {
		g.meteorSpawnTimer.Reset()
		g.meteors = append(g.meteors, NewMeteor())
	}

	for _, m := range g.meteors {
		m.Update()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.player.Draw(screen)

	for _, m := range g.meteors {
		m.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth int, outsideHeight int) (screenWidth int, screenHeight int) {
	return ScreenWidth, ScreenHeight
}
