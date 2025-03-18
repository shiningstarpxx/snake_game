package main

import (
    "fmt"
    "github.com/yourusername/snake-game/src/game"
    "github.com/yourusername/snake-game/src/ui"
    "github.com/yourusername/snake-game/src/utils"
    "time"
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