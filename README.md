# Snake Game

This is a simple implementation of the classic Snake game written in Go. The game features a snake that the player controls to eat food and grow in length while avoiding collisions with itself and the walls.

## Project Structure

```
snake-game
├── src
│   ├── main.go          # Entry point of the application
│   ├── game
│   │   ├── board.go     # Game board logic
│   │   ├── snake.go     # Snake logic
│   │   └── food.go      # Food logic
│   ├── ui
│   │   └── renderer.go   # Rendering logic
│   └── utils
│       └── constants.go  # Constants used throughout the game
├── assets
│   └── fonts
│       └── README.md     # Documentation for fonts used
├── go.mod                # Module definition
├── go.sum                # Dependency checksums
└── README.md             # Project documentation
```

## Requirements

- Go 1.16 or later

## How to Run the Game

1. Clone the repository:
   ```
   git clone <repository-url>
   cd snake-game
   ```

2. Install dependencies:
   ```
   go mod tidy
   ```

3. Run the game:
   ```
   go run src/main.go
   ```

## Controls

- Use the arrow keys to control the direction of the snake.
- The objective is to eat the food that appears on the board to grow the snake.

## Contributing

Feel free to submit issues or pull requests for improvements or bug fixes.