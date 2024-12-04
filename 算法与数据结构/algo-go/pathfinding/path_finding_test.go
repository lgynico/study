package pathfinding

import (
	"strings"
	"testing"
)

var mapStr = `
000000000000000000100000000000
000000111111011000100000000000
000000000010000000111110000000
000000000010000000000000000000
000000000000000000000000000000
`

func TestAStar(t *testing.T) {
	var (
		lines    = strings.Fields(mapStr)
		lenX     = len(lines[0])
		lenZ     = len(lines)
		sceneMap = NewSceneMap(lenX, lenZ)
	)

	for z, line := range lines {
		for x, str := range line {
			switch str {
			case '1':
				sceneMap.Set(x, z, CellWall)
			default:
				sceneMap.Set(x, z, CellGround)
			}
		}
	}

	var (
		origX, origZ = 1, 4
		destX, destZ = 24, 0
	)

	if path := sceneMap.AStar(origX, origZ, destX, destZ); path != nil {
		for _, point := range path {
			x, z := sceneMap.decombineToXZ(point)
			sceneMap.Set(x, z, CellPath)
		}
	}
	sceneMap.Set(origX, origZ, CellOrigin)
	sceneMap.Set(destX, destZ, CellDestination)

	sceneMap.Print()
}

func TestAStar2(t *testing.T) {
	var (
		origX, origZ = 10, 4
		destX, destZ = 24, 0
		m            = Map{}
	)

	m.Init(mapStr)

	if path := m.AStar(origX, origZ, destX, destZ); path != nil {
		for _, point := range path {
			m.Set(point.x, point.z, CellPath)
		}
	}
	m.Set(origX, origZ, CellOrigin)
	m.Set(destX, destZ, CellDestination)

	m.Print()
}
