package main

import (
	"fmt"
	"time"

	"github.com/shiningstarpxx/snake_game/src/game"
	"github.com/shiningstarpxx/snake_game/src/ui"
	"github.com/shiningstarpxx/snake_game/src/utils"
)

func main() {
	board := game.NewBoard()
	snake := game.NewSnake()
	food := game.NewFood()

	ui.Initialize()

	for {
		ui.Render(board, snake, food)
		snake.Move()

		if snake.CheckCollision(board) {
			fmt.Println("Game Over!")
			break
		}

		if snake.Eat(food) {
			food.Spawn(board)
		}

		time.Sleep(time.Millisecond * utils.GameSpeed)
	}
}
