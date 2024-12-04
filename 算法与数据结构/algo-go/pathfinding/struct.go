package pathfinding

import "fmt"

type CellType int

const (
	CellGround CellType = iota
	CellWall
	CellPath
	CellOrigin
	CellDestination
)

type SceneMap struct {
	lengthX int
	lengthZ int
	cells   []CellType
}

func NewSceneMap(lengthX, lengthZ int) *SceneMap {
	return &SceneMap{lengthX: lengthX, lengthZ: lengthZ, cells: make([]CellType, lengthX*lengthZ)}
}

// func (p *SceneMap) Init(mapStr string) {

// }

func (p *SceneMap) CellType(x, z int) CellType {
	cell := p.combineToCell(x, z)
	return p.cells[cell]
}

func (p *SceneMap) Set(x, z int, cellType CellType) {
	cell := p.combineToCell(x, z)
	p.cells[cell] = cellType
}

func (p *SceneMap) Border() (lenX, lenZ int) {
	return p.lengthX, p.lengthZ
}

func (p *SceneMap) IsInside(x, z int) bool {
	return x >= 0 && x < p.lengthX && z >= 0 && z < p.lengthZ
}

func (p *SceneMap) combineToCell(x, z int) int {
	return z*p.lengthX + x
}

func (p *SceneMap) decombineToXZ(cell int) (x, z int) {
	x = cell % p.lengthX
	z = cell / p.lengthX
	return
}

func (p *SceneMap) Print() {
	for i, t := range p.cells {
		x, _ := p.decombineToXZ(i)
		switch t {
		case CellGround:
			fmt.Print(" ")
		case CellWall:
			fmt.Print("0")
		case CellPath:
			fmt.Print("*")
		case CellDestination:
			fmt.Print("$")
		case CellOrigin:
			fmt.Print("^")
		default:
			fmt.Print("?")
		}
		if x == p.lengthX-1 {
			fmt.Println()
		}
	}
	fmt.Println()
}
