package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/shiningstarpxx/snake_game/src/utils"
)

const (
	boardWidth  = 20
	boardHeight = 20
)

type Board struct {
	grid   [][]int
	Width  int
	Height int
}

func NewBoard() *Board {
	board := &Board{
		grid:   make([][]int, boardHeight),
		Width:  utils.ScreenWidth / utils.SnakeSize,
		Height: utils.ScreenHeight / utils.SnakeSize,
	}
	for i := range board.grid {
		board.grid[i] = make([]int, boardWidth)
	}
	return board
}

func (b *Board) Render() {
	for _, row := range b.grid {
		for _, cell := range row {
			if cell == 1 {
				fmt.Print("O ") // Snake
			} else if cell == 2 {
				fmt.Print("X ") // Food
			} else {
				fmt.Print(". ") // Empty
			}
		}
		fmt.Println()
	}
}

func (b *Board) CheckCollision(x, y int) bool {
	if x < 0 || x >= boardWidth || y < 0 || y >= boardHeight {
		return true // Collision with wall
	}
	return b.grid[y][x] == 1 // Collision with snake
}

func (b *Board) SpawnFood() {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(boardWidth)
	y := rand.Intn(boardHeight)

	for b.grid[y][x] != 0 { // Ensure food does not spawn on snake
		x = rand.Intn(boardWidth)
		y = rand.Intn(boardHeight)
	}

	b.grid[y][x] = 2 // Place food
}

func (b *Board) Draw(screen *ebiten.Image) {
	// Draw the background
	bgColor, _ := ParseHexColor(utils.BackgroundColor)
	screen.Fill(bgColor)

	// Draw grid lines (optional)
	gridColor := color.RGBA{40, 40, 40, 255} // Dark grey

	// Draw vertical grid lines
	for x := 0; x < b.Width; x++ {
		xPos := float64(x * utils.SnakeSize)
		ebitenutil.DrawLine(screen, xPos, 0, xPos, float64(utils.ScreenHeight), gridColor)
	}

	// Draw horizontal grid lines
	for y := 0; y < b.Height; y++ {
		yPos := float64(y * utils.SnakeSize)
		ebitenutil.DrawLine(screen, 0, yPos, float64(utils.ScreenWidth), yPos, gridColor)
	}
}
