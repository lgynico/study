package main

import (
	"image/color"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 320
	screenHeight = 320
	cellSize     = 32
	gridWidth    = 10
	gridHeight   = 10
	mineCount    = 15
)

type CellState int

const (
	CellHidden CellState = iota
	CellRevealed
	CellFlagged
)

type Cell struct {
	isMine    bool
	state     CellState
	neighbors int
}

type Game struct {
	grid        [][]Cell
	gameOver    bool
	win         bool
	firstClick  bool
	fontFace    text.Face
	pressedKeys []ebiten.Key
}

func NewGame() *Game {
	grid := make([][]Cell, gridHeight)
	for i := range grid {
		grid[i] = make([]Cell, gridWidth)
		for j := range grid[i] {
			grid[i][j] = Cell{
				isMine:    false,
				state:     CellHidden,
				neighbors: 0,
			}
		}
	}

	// 创建字体
	fontFace := text.NewGoXFace(basicfont.Face7x13)

	return &Game{
		grid:       grid,
		gameOver:   false,
		win:        false,
		firstClick: true,
		fontFace:   fontFace,
	}
}

func (g *Game) placeMines(startX, startY int) {
	rand.Seed(time.Now().UnixNano())
	mines := 0

	for mines < mineCount {
		x := rand.Intn(gridWidth)
		y := rand.Intn(gridHeight)

		// 确保第一次点击的周围没有地雷
		if g.grid[y][x].isMine || (abs(x-startX) <= 1 && abs(y-startY) <= 1) {
			continue
		}

		g.grid[y][x].isMine = true
		mines++

		// 更新周围格子的地雷计数
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}

				nx, ny := x+dx, y+dy
				if nx >= 0 && nx < gridWidth && ny >= 0 && ny < gridHeight {
					g.grid[ny][nx].neighbors++
				}
			}
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (g *Game) revealCell(x, y int) {
	if x < 0 || x >= gridWidth || y < 0 || y >= gridHeight {
		return
	}

	cell := &g.grid[y][x]
	if cell.state != CellHidden {
		return
	}

	cell.state = CellRevealed

	if cell.isMine {
		g.gameOver = true
		return
	}

	// 如果是空白格子，自动揭开周围的格子
	if cell.neighbors == 0 {
		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				if dx == 0 && dy == 0 {
					continue
				}
				g.revealCell(x+dx, y+dy)
			}
		}
	}

	// 检查胜利条件
	g.checkWin()
}

func (g *Game) flagCell(x, y int) {
	if x < 0 || x >= gridWidth || y < 0 || y >= gridHeight {
		return
	}

	cell := &g.grid[y][x]
	if cell.state == CellHidden {
		cell.state = CellFlagged
	} else if cell.state == CellFlagged {
		cell.state = CellHidden
	}
}

func (g *Game) checkWin() {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			cell := g.grid[y][x]
			if !cell.isMine && cell.state != CellRevealed {
				return
			}
		}
	}
	g.win = true
}

func (g *Game) Update() error {
	// 处理键盘输入
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])
	for _, key := range g.pressedKeys {
		if key == ebiten.KeyR {
			*g = *NewGame()
			return nil
		}
	}

	if g.gameOver || g.win {
		return nil
	}

	// 处理鼠标输入
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		gridX, gridY := x/cellSize, y/cellSize

		if gridX >= 0 && gridX < gridWidth && gridY >= 0 && gridY < gridHeight {
			if g.firstClick {
				g.placeMines(gridX, gridY)
				g.firstClick = false
			}
			g.revealCell(gridX, gridY)
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		x, y := ebiten.CursorPosition()
		gridX, gridY := x/cellSize, y/cellSize

		if gridX >= 0 && gridX < gridWidth && gridY >= 0 && gridY < gridHeight {
			g.flagCell(gridX, gridY)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 绘制网格
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			cell := g.grid[y][x]
			rect := color.RGBA{200, 200, 200, 255}

			switch cell.state {
			case CellHidden:
				rect = color.RGBA{150, 150, 150, 255}
			case CellFlagged:
				rect = color.RGBA{255, 150, 150, 255}
			case CellRevealed:
				if cell.isMine {
					rect = color.RGBA{255, 0, 0, 255}
				} else {
					rect = color.RGBA{220, 220, 220, 255}
				}
			}

			ebitenutil.DrawRect(screen, float64(x*cellSize), float64(y*cellSize), cellSize-1, cellSize-1, rect)

			if cell.state == CellRevealed && !cell.isMine && cell.neighbors > 0 {
				numColor := color.RGBA{0, 0, 255, 255}
				if cell.neighbors == 2 {
					numColor = color.RGBA{0, 150, 0, 255}
				} else if cell.neighbors >= 3 {
					numColor = color.RGBA{255, 0, 0, 255}
				}

				// 使用新的 text.Draw API
				textOp := &text.DrawOptions{}
				// 设置位置
				textOp.GeoM.Translate(float64(x*cellSize+cellSize/2-5), float64(y*cellSize+cellSize/2+5))
				// 设置颜色
				textOp.ColorScale.SetR(float32(numColor.R) / 255)
				textOp.ColorScale.SetG(float32(numColor.G) / 255)
				textOp.ColorScale.SetB(float32(numColor.B) / 255)
				textOp.ColorScale.SetA(float32(numColor.A) / 255)

				text.Draw(
					screen,
					strconv.Itoa(cell.neighbors),
					g.fontFace,
					textOp,
				)
			}

			if cell.state == CellFlagged {
				// 使用新的 text.Draw API
				textOp := &text.DrawOptions{}
				// 设置位置
				textOp.GeoM.Translate(float64(x*cellSize+cellSize/2-5), float64(y*cellSize+cellSize/2+5))
				// 设置红色
				textOp.ColorScale.SetR(1.0)
				textOp.ColorScale.SetG(0.0)
				textOp.ColorScale.SetB(0.0)
				textOp.ColorScale.SetA(1.0)

				text.Draw(
					screen,
					"F",
					g.fontFace,
					textOp,
				)
			}
		}
	}

	// 显示游戏状态
	if g.gameOver {
		ebitenutil.DebugPrintAt(screen, "游戏结束! 按R键重新开始", 10, screenHeight-20)
		// 显示所有地雷
		for y := 0; y < gridHeight; y++ {
			for x := 0; x < gridWidth; x++ {
				if g.grid[y][x].isMine && g.grid[y][x].state != CellRevealed {
					g.grid[y][x].state = CellRevealed
				}
			}
		}
	} else if g.win {
		ebitenutil.DebugPrintAt(screen, "恭喜您赢了! 按R键重新开始", 10, screenHeight-20)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("扫雷游戏")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
