package ui

import (
    "github.com/hajimehoshi/ebiten/v2"
    "github.com/hajimehoshi/ebiten/v2/ebitenutil"
    "snake-game/src/game"
)

type Renderer struct {
    board  *game.Board
    snake  *game.Snake
    food   *game.Food
}

func NewRenderer(board *game.Board, snake *game.Snake, food *game.Food) *Renderer {
    return &Renderer{
        board: board,
        snake: snake,
        food:  food,
    }
}

func (r *Renderer) Draw(screen *ebiten.Image) {
    r.board.Draw(screen)
    r.snake.Draw(screen)
    r.food.Draw(screen)
}

func (r *Renderer) Layout(outsideWidth, outsideHeight int) (int, int) {
    return outsideWidth, outsideHeight
}