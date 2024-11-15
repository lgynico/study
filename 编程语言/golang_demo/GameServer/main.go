package main

import (
	"flag"

	"github.com/lgynico/gameserver/game"
	"github.com/lgynico/gameserver/gate"
)

var node string

func init() {
	flag.StringVar(&node, "node", "", "run node type")
}

func main() {
	flag.Parse()

	switch node {
	case "game":
		game.Run()
	case "gate":
		gate.Run()
	default:
		panic("unknown node type: " + node)
	}
}
