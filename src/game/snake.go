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

type Direction int

const (
	Up Direction = iota
	Right
	Down
	Left
)

type Position struct {
	X int
	Y int
}

type Snake struct {
	Body      []Position
	Direction Direction
	Growing   bool
}

// NewSnake creates a new snake with the given starting position
func NewSnake(startPos Position) *Snake {
	return &Snake{
		Body:      []Position{startPos},
		Direction: Right,
		Growing:   false,
	}
}

// Move updates the snake's position based on its direction
func (s *Snake) Move() {
	head := s.Body[0]
	newHead := Position{X: head.X, Y: head.Y}

	switch s.Direction {
	case Up:
		newHead.Y--
	case Right:
		newHead.X++
	case Down:
		newHead.Y++
	case Left:
		newHead.X--
	}

	// Add new head to the beginning of the body
	s.Body = append([]Position{newHead}, s.Body...)

	// Remove the tail if not growing
	if !s.Growing {
		s.Body = s.Body[:len(s.Body)-1]
	} else {
		s.Growing = false
	}
}

// CheckCollision checks if the snake has collided with walls or itself
func (s *Snake) CheckCollision() bool {
	head := s.Body[0]

	// Check wall collision
	if head.X < 0 || head.X >= utils.ScreenWidth/utils.SnakeSize ||
		head.Y < 0 || head.Y >= utils.ScreenHeight/utils.SnakeSize {
		return true
	}

	// Check self collision
	for i := 1; i < len(s.Body); i++ {
		if head.X == s.Body[i].X && head.Y == s.Body[i].Y {
			return true
		}
	}

	return false
}

func (s *Snake) ChangeDirection(newDirection Direction) {
	if (s.Direction == Up && newDirection == Down) || (s.Direction == Down && newDirection == Up) ||
		(s.Direction == Left && newDirection == Right) || (s.Direction == Right && newDirection == Left) {
		return
	}
	s.Direction = newDirection
}

// Eat checks if the snake has eaten food and grows the snake if so
func (s *Snake) Eat(food *Food) bool {
	head := s.Body[0]
	if head.X == food.Position.X && head.Y == food.Position.Y {
		s.Growing = true
		return true
	}
	return false
}

// Draw renders the snake on the screen
func (s *Snake) Draw(screen *ebiten.Image) {
	// Convert hex color to RGBA
	var r, g, b uint64
	colorStr := utils.SnakeColor
	if colorStr[0] == '#' {
		colorStr = colorStr[1:]
	}
	r, _ = strconv.ParseUint(colorStr[0:2], 16, 8)
	g, _ = strconv.ParseUint(colorStr[2:4], 16, 8)
	b, _ = strconv.ParseUint(colorStr[4:6], 16, 8)
	snakeColor := color.RGBA{uint8(r), uint8(g), uint8(b), 255}

	// Draw each segment
	for _, pos := range s.Body {
		x := float64(pos.X * utils.SnakeSize)
		y := float64(pos.Y * utils.SnakeSize)

		ebitenutil.DrawRect(
			screen,
			x, y,
			float64(utils.SnakeSize),
			float64(utils.SnakeSize),
			snakeColor,
		)
	}
}

func RandomPosition() Position {
	rand.Seed(time.Now().UnixNano())
	return Position{
		X: rand.Intn(utils.ScreenWidth / utils.SnakeSize),
		Y: rand.Intn(utils.ScreenHeight / utils.SnakeSize),
	}
}
