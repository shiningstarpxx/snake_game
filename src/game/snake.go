package game

import (
	"math/rand"
	"time"
)

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

type Snake struct {
	Body      []Position
	Direction Direction
}

type Position struct {
	X int
	Y int
}

func NewSnake(initialPosition Position) *Snake {
	return &Snake{
		Body:      []Position{initialPosition},
		Direction: Right,
	}
}

func (s *Snake) Move() {
	head := s.Body[0]
	var newHead Position

	switch s.Direction {
	case Up:
		newHead = Position{X: head.X, Y: head.Y - 1}
	case Down:
		newHead = Position{X: head.X, Y: head.Y + 1}
	case Left:
		newHead = Position{X: head.X - 1, Y: head.Y}
	case Right:
		newHead = Position{X: head.X + 1, Y: head.Y}
	}

	s.Body = append([]Position{newHead}, s.Body[:len(s.Body)-1]...)
}

func (s *Snake) Grow() {
	tail := s.Body[len(s.Body)-1]
	s.Body = append(s.Body, tail)
}

func (s *Snake) CheckCollision() bool {
	head := s.Body[0]

	// Check wall collision
	if head.X < 0 || head.Y < 0 || head.X >= boardWidth || head.Y >= boardHeight {
		return true
	}

	// Check self collision
	for i := 1; i < len(s.Body); i++ {
		if head == s.Body[i] {
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

func RandomPosition() Position {
	rand.Seed(time.Now().UnixNano())
	return Position{
		X: rand.Intn(boardWidth),
		Y: rand.Intn(boardHeight),
	}
}