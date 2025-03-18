package main

import (
	"fmt"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/shiningstarpxx/snake_game/src/game"
	"github.com/shiningstarpxx/snake_game/src/ui"
	"github.com/shiningstarpxx/snake_game/src/utils"
)

type Game struct {
	board      *game.Board
	snake      *game.Snake
	food       *game.Food
	renderer   *ui.Renderer
	gameOver   bool
	score      int
	keyPressed bool // Track if a key was pressed this frame
}

func NewGame() *Game {
	board := game.NewBoard()
	snake := game.NewSnake(game.Position{X: 10, Y: 10})
	food := game.NewFood(board.Width, board.Height)
	renderer := ui.NewRenderer(board, snake, food)

	return &Game{
		board:      board,
		snake:      snake,
		food:       food,
		renderer:   renderer,
		gameOver:   false,
		score:      0,
		keyPressed: false,
	}
}

func (g *Game) Update() error {
	g.keyPressed = false

	// Only handle input if the game is not over
	if !g.gameOver {
		// Only move the snake when a key is pressed
		if inpututil.IsKeyJustPressed(ebiten.KeyUp) && g.snake.Direction != game.Down {
			g.snake.Direction = game.Up
			g.keyPressed = true
		} else if inpututil.IsKeyJustPressed(ebiten.KeyRight) && g.snake.Direction != game.Left {
			g.snake.Direction = game.Right
			g.keyPressed = true
		} else if inpututil.IsKeyJustPressed(ebiten.KeyDown) && g.snake.Direction != game.Up {
			g.snake.Direction = game.Down
			g.keyPressed = true
		} else if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && g.snake.Direction != game.Right {
			g.snake.Direction = game.Left
			g.keyPressed = true
		}

		// Only move the snake if a key was pressed
		if g.keyPressed {
			// Move the snake
			g.snake.Move()

			// Check for collision with walls or self
			if g.snake.CheckCollision() {
				g.gameOver = true
				return nil
			}

			// Check if snake ate food
			if g.snake.Eat(g.food) {
				g.food.Spawn(g.board.Width, g.board.Height)
				g.score++
			}
		}
	} else {
		// Allow restarting the game with Enter key when game over
		if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
			g.snake = game.NewSnake(game.Position{X: 10, Y: 10})
			g.food = game.NewFood(g.board.Width, g.board.Height)
			g.gameOver = false
			g.score = 0
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the game elements
	g.renderer.Draw(screen)

	// Draw the score
	scoreText := fmt.Sprintf("Score: %d", g.score)
	ebitenutil.DebugPrint(screen, scoreText)

	// Draw game over message
	if g.gameOver {
		gameOverText := "Game Over! Press Enter to restart"
		textWidth := len(gameOverText) * 6 // Approximate width of text
		x := (utils.ScreenWidth - textWidth) / 2
		ebitenutil.DebugPrintAt(screen, gameOverText, x, utils.ScreenHeight/2)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return utils.ScreenWidth, utils.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(utils.ScreenWidth, utils.ScreenHeight)
	ebiten.SetWindowTitle("Snake Game")

	if err := ebiten.RunGame(NewGame()); err != nil {
		log.Fatal(err)
	}
}
