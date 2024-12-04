package pathfinding

import (
	"fmt"
	"math"
	"strings"
)

type point struct {
	x, z int
}

type Map struct {
	lengthX, lengthZ int
	cells            []*cell
}

func (p *Map) Init(mapStr string) {
	lines := strings.Fields(mapStr)
	p.lengthZ = len(lines)

	if p.lengthZ > 0 {
		p.lengthX = len(lines[0])
	}

	p.cells = make([]*cell, p.lengthX*p.lengthZ)

	for z, line := range lines {
		for x, str := range line {
			cell := cell{x: x, z: z}
			switch str {
			case '0':
				cell.cellType = CellGround
			case '1':
				cell.cellType = CellWall
			}
			p.cells[p.toIndex(x, z)] = &cell
		}
	}
}

func (p *Map) Set(x, z int, cellType CellType) {
	if !p.IsInside(x, z) {
		return
	}

	p.Cell(x, z).cellType = cellType
}

func (p *Map) IsInside(x, z int) bool {
	return x >= 0 && x < p.lengthX && z >= 0 && z < p.lengthZ
}

func (p *Map) Cell(x int, z int) *cell {
	if !p.IsInside(x, z) {
		return nil
	}

	return p.cells[p.toIndex(x, z)]
}

func (p *Map) Print() {
	for i, cell := range p.cells {
		switch cell.cellType {
		case CellGround:
			// fmt.Print("口")
			fmt.Print("░░")
			// fmt.Printf("%v ", "\u2591")
		case CellWall:
			fmt.Print("墙")
		case CellPath:
			// fmt.Print("走")
			fmt.Print("口")
			// fmt.Printf("%v ", "\u0A0A")
		case CellDestination:
			fmt.Print("终")
		case CellOrigin:
			fmt.Print("起")
		default:
			fmt.Print("？")
		}
		if (i+1)%p.lengthX == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (p *Map) toIndex(x, z int) int {
	return z*p.lengthX + x
}

func (p *Map) neighbors(x, z int) []*cell {
	cells := make([]*cell, 0, 4)
	if p.IsInside(x+1, z) {
		cells = append(cells, p.Cell(x+1, z))
	}
	if p.IsInside(x-1, z) {
		cells = append(cells, p.Cell(x-1, z))
	}
	if p.IsInside(x, z+1) {
		cells = append(cells, p.Cell(x, z+1))
	}
	if p.IsInside(x, z-1) {
		cells = append(cells, p.Cell(x, z-1))
	}

	if p.IsInside(x+1, z+1) {
		cells = append(cells, p.Cell(x+1, z+1))
	}
	if p.IsInside(x-1, z-1) {
		cells = append(cells, p.Cell(x-1, z-1))
	}
	if p.IsInside(x-1, z+1) {
		cells = append(cells, p.Cell(x-1, z+1))
	}
	if p.IsInside(x+1, z-1) {
		cells = append(cells, p.Cell(x+1, z-1))
	}
	return cells
}

func (p *Map) heuristic(origX, origZ, destX, destZ int) int {
	return int(math.Abs(float64(destX-origX))+math.Abs(float64(destZ-origZ))) * 10
}

// func (p *SceneMap) decombineToXZ(cell int) (x, z int) {
// 	x = cell % p.lengthX
// 	z = cell / p.lengthX
// 	return
// }

type cell struct {
	x, z         int
	cellType     CellType
	costH, costG int
	front        *cell
}

func (p *cell) isWalkable() bool {
	return p.cellType != CellWall
}
