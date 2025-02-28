package pathfinding

import (
	"math"

	"github.com/emirpasic/gods/maps/treemap"
)

func (p *SceneMap) AStar(srcX, srcZ, dstX, dstZ int) []int {
	if !p.IsInside(srcX, srcZ) || !p.IsInside(dstX, dstZ) {
		return nil
	}

	var (
		visited = map[int]struct{}{}
		fronts  = map[int]int{}
		gcosts  = map[int]int{}
		fcosts  = map[int]int{}
		found   = false
		src     = p.combineToCell(srcX, srcZ)
		dst     = p.combineToCell(dstX, dstZ)
		queue   = treemap.NewWith(func(a, b any) int {
			var (
				costA = fcosts[a.(int)]
				costB = fcosts[b.(int)]
			)
			if costA > costB {
				return 1
			}
			if costB > costA {
				return -1
			}
			return 0
		})
	)

	queue.Put(src, 0)

outer:
	for !queue.Empty() {
		var (
			key, _  = queue.Min()
			curCell = key.(int)
			cells   = p.neighbors(curCell)
		)

		queue.Remove(key)

		for _, cell := range cells {
			if _, ok := visited[cell]; ok {
				continue
			}

			var x, z = p.decombineToXZ(cell)
			if !p.IsInside(x, z) || p.CellType(x, z) == CellWall {
				continue
			}

			if cell == dst {
				found = true
				fronts[cell] = curCell
				break outer
			}

			gcosts[cell] = gcosts[curCell] + 1
			fcosts[cell] = gcosts[cell] + p.heuristic(cell, dst)
			visited[cell] = struct{}{}
			fronts[cell] = curCell

			queue.Put(cell, 0)
		}

	}

	var path []int
	if found {
		for i := dst; i != src; i = fronts[i] {
			path = append(path, i)
		}
	}

	return path
}

func (p *SceneMap) neighbors(cell int) []int {
	x, z := p.decombineToXZ(cell)
	cells := make([]int, 0, 4)

	if x-1 >= 0 {
		cells = append(cells, p.combineToCell(x-1, z))
	}

	if x+1 < p.lengthX {
		cells = append(cells, p.combineToCell(x+1, z))
	}

	if z-1 >= 0 {
		cells = append(cells, p.combineToCell(x, z-1))
	}

	if z+1 < p.lengthZ {
		cells = append(cells, p.combineToCell(x, z+1))
	}

	return cells
}

func (p *SceneMap) heuristic(src, dst int) int {
	var (
		srcX, srcZ = p.decombineToXZ(src)
		dstX, dstZ = p.decombineToXZ(dst)
	)
	return int(math.Abs(float64(dstX-srcX)) + math.Abs(float64(dstZ-srcZ)))
}

func (p *Map) AStar(origX, origZ, destX, destZ int) []*point {
	if !(p.IsInside(origX, origZ) && p.IsInside(destX, destZ)) {
		return nil
	}

	var (
		openList  = treemap.NewWith(p.compareCell)
		closeList = map[*cell]struct{}{}
	)

	openList.Put(p.Cell(origX, origZ), struct{}{})
	closeList[p.Cell(origX, origZ)] = struct{}{}

outer:
	for !openList.Empty() {
		var (
			key, _  = openList.Min()
			curCell = key.(*cell)
		)

		openList.Remove(key)

		for _, cell := range p.neighbors(curCell.x, curCell.z) {
			if !cell.isWalkable() {
				continue
			}

			if _, ok := closeList[cell]; ok {
				continue
			}

			cell.front = curCell

			if cell.x == destX && cell.z == destZ {
				break outer
			}

			var (
				H = p.heuristic(cell.x, cell.z, destX, destZ)
				G = 10
			)

			if int(math.Abs(float64(cell.x-cell.front.x))+math.Abs(float64(cell.z-cell.front.z))) == 2 {
				G += 4
			}

			if _, ok := openList.Get(cell); ok {
				if cell.costG+cell.costH < H+G {
					cell.costH = H
					cell.costG = curCell.costG + G
				}
			} else {
				cell.costH = H
				cell.costG = curCell.costG + G
				openList.Put(cell, struct{}{})
			}

			closeList[cell] = struct{}{}
		}
	}

	var (
		path = make([]*point, 0, 1024)
		orig = p.Cell(origX, origZ)
		dest = p.Cell(destX, destZ)
	)

	for cell := dest; cell != orig; cell = cell.front {
		path = append(path, &point{x: cell.x, z: cell.z})
	}

	return path
}

func (p *Map) compareCell(a, b any) int {
	var (
		cellA = a.(*cell)
		cellB = b.(*cell)
		costA = cellA.costG + cellA.costH
		costB = cellB.costG + cellB.costH
	)

	if costA > costB {
		return 1
	}
	if costB > costA {
		return -1
	}
	return 0
}
