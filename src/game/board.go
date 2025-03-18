package game

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	boardWidth  = 20
	boardHeight = 20
)

type Board struct {
	grid [][]int
}

func NewBoard() *Board {
	board := &Board{
		grid: make([][]int, boardHeight),
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