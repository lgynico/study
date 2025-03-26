package aoi

type (
	QuadTreeNode struct {
		Entities   []any
		Depth      int
		IsLeaf     bool
		Width      float64
		PosX, PosY float64
		Children   [4]*QuadTreeNode
		t          *QuadTree
	}

	QuadTree struct {
		maxDepth   int
		maxPayload int
		root       QuadTreeNode
	}
)

func NewQuadTree(depth, payload int, width float64) *QuadTree {
	t := &QuadTree{
		maxDepth:   depth,
		maxPayload: payload,
		root: QuadTreeNode{
			Depth:  0,
			IsLeaf: true,
			Width:  width,
			PosX:   0,
			PosY:   0,
		},
	}

	t.root.t = t
	return t
}

func (p *QuadTreeNode) Cutable() bool {
	return (p.PosX+p.Width)/2 > 0 && (p.PosY+p.Width)/2 > 0
}
