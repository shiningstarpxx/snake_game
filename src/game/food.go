package game

import (
	"image/color"
	"math/rand"
	"strconv"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/shiningstarpxx/snake_game/src/utils"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type Food struct {
	Position Position
}

// NewFood creates a new food at a random position within the board dimensions
func NewFood(boardWidth, boardHeight int) *Food {
	return &Food{
		Position: Position{
			X: rand.Intn(boardWidth),
			Y: rand.Intn(boardHeight),
		},
	}
}

// Spawn moves the food to a new random position
func (f *Food) Spawn(boardWidth, boardHeight int) {
	f.Position = Position{
		X: rand.Intn(boardWidth),
		Y: rand.Intn(boardHeight),
	}
}

// Draw renders the food on the screen
func (f *Food) Draw(screen *ebiten.Image) {
	// Convert hex color to RGBA
	var r, g, b uint64
	colorStr := utils.FoodColor
	if colorStr[0] == '#' {
		colorStr = colorStr[1:]
	}
	r, _ = strconv.ParseUint(colorStr[0:2], 16, 8)
	g, _ = strconv.ParseUint(colorStr[2:4], 16, 8)
	b, _ = strconv.ParseUint(colorStr[4:6], 16, 8)
	foodColor := color.RGBA{uint8(r), uint8(g), uint8(b), 255}

	x := float64(f.Position.X * utils.SnakeSize)
	y := float64(f.Position.Y * utils.SnakeSize)

	ebitenutil.DrawRect(
		screen,
		x, y,
		float64(utils.FoodSize),
		float64(utils.FoodSize),
		foodColor,
	)
}
