package game

import (
	"math/rand"
	"time"
)

type Food struct {
	X int
	Y int
}

func NewFood(boardWidth int, boardHeight int) *Food {
	rand.Seed(time.Now().UnixNano())
	return &Food{
		X: rand.Intn(boardWidth),
		Y: rand.Intn(boardHeight),
	}
}

func (f *Food) Spawn(boardWidth int, boardHeight int) {
	f.X = rand.Intn(boardWidth)
	f.Y = rand.Intn(boardHeight)
}