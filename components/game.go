package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"os"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Game Window"
	circleRadius = 8
	cellSize     = 20
	textHeight   = 20
)

type game struct {
	window       *sdl.Window
	renderer     *sdl.Renderer
	fontSize     int
	textToRender string
	circleX      int
	circleY      int
	playerX      int
	playerY      int
	initializing bool
	score		 int
}

// Create a new game instance
func newGame() *game {
	g := &game{}
	g.fontSize = 10
	g.textToRender = ""
	g.circleX = windowWidth / 2
	g.circleY = windowHeight / 2
	g.playerX = 1
	g.playerY = 1
	g.score   = 0

	return g
}

// Initialize the game window and renderer
func (g *game) init() error {
	var err error

	// Create the window
	if g.window, err = sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_CENTERED,
		sdl.WINDOWPOS_CENTERED, windowWidth, windowHeight, sdl.WINDOW_SHOWN); err != nil {
		return fmt.Errorf("Error creating Window: %v", err)
	}

	// Create the renderer
	if g.renderer, err = sdl.CreateRenderer(g.window, -1, sdl.RENDERER_ACCELERATED); err != nil {
		return fmt.Errorf("Error creating Renderer: %v", err)
	}

	return nil
}

// Close the game and clean up resources
func (g *game) close() {
	if g != nil {
		g.renderer.Destroy()
		g.renderer = nil
		g.window.Destroy()
		g.window = nil
	}
}

func (g *game) run() {
	running := true
	g.initializing = true
	var matrix [][]int
	var selectedMatrixIndex int

	// Render initialization image
	g.renderInitializationImage()

	// Delay for 5 seconds
	sdl.Delay(5000) // 5000 milliseconds = 5 seconds

	// Matrix selection phase
	for g.initializing {
		// Handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
				g.initializing = false // Exit initialization phase on quit
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_1:
						matrix = matrix1 // Assign matrix1 to be played
						selectedMatrixIndex = 1
						g.initializing = false // Exit initialization phase
					case sdl.K_2:
						matrix = matrix2 // Assign matrix2 to be played
						selectedMatrixIndex = 2
						g.initializing = false // Exit initialization phase
					case sdl.K_3:
						matrix = matrix3 // Assign matrix2 to be played
						selectedMatrixIndex = 3
						g.initializing = false // Exit initialization phase
					case sdl.K_4:
						matrix = matrix4 // Assign matrix2 to be played
						selectedMatrixIndex = 4
						g.initializing = false
					case sdl.K_5:
						matrix = matrix5 // Assign matrix2 to be played
						selectedMatrixIndex = 5
						g.initializing = false
					}
				}
			}
		}

		// Render matrix selection interface
		g.renderer.SetDrawColor(0, 0, 0, 255)
		g.renderer.Clear()

		// Example: Render matrix selection text
		if err := g.renderText("Choose Matrix:", 20, 20, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}
		if err := g.renderText("Press 1 for Matrix 1", 20, 50, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}
		if err := g.renderText("Press 2 for Matrix 2", 20, 80, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}
		if err := g.renderText("Press 3 for Matrix 3", 20, 110, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}
		if err := g.renderText("Press 4 for Matrix 4", 20, 140, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}
		if err := g.renderText("Press 5 for Matrix 5", 20, 170, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}
		g.renderer.Present() // Present the renderer

		sdl.Delay(16) // Delay to control frame rate
	}

	if !running {
		return
	}

	// Continue with the selected matrix
	g.initializePlayerPosition(matrix)

	initialTwosCount := countTwos(matrix) // Calculate initial count of 2s
	twosCount := initialTwosCount         // Current count of 2s

	fmt.Printf("Initial count of 2s in Matrix %d: %d\n", selectedMatrixIndex, initialTwosCount) // Print the initial count

	for running {
		// Handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.Type == sdl.KEYDOWN {
					if t.Keysym.Sym == sdl.K_ESCAPE {
						running = false
					}

					moved := false // Track if the player actually moved
					switch t.Keysym.Sym {
					case sdl.K_w:
						if g.playerY > 0 && matrix[g.playerY-1][g.playerX] != 1 {
							g.playerY--
							moved = true
						}
					case sdl.K_a:
						if g.playerX > 0 && matrix[g.playerY][g.playerX-1] != 1 {
							g.playerX--
							moved = true
						}
					case sdl.K_s:
						if g.playerY < len(matrix)-1 && matrix[g.playerY+1][g.playerX] != 1 {
							g.playerY++
							moved = true
						}
					case sdl.K_d:
						if g.playerX < len(matrix[0])-1 && matrix[g.playerY][g.playerX+1] != 1 {
							g.playerX++
							moved = true
						}
					}

					if moved && matrix[g.playerY][g.playerX] == 2 {
						matrix[g.playerY][g.playerX] = 0
						twosCount-- // Decrement the count of 2s

						// Exit the game loop if no 2s are left
						if twosCount == 0 {
							fmt.Printf("No more 2s remaining in Matrix %d, exiting the game.\n", selectedMatrixIndex)
							running = false
						}
					}
				}
			}
		}

		// Render the scene
		g.renderer.SetDrawColor(0, 0, 0, 255)
		g.renderer.Clear()

		g.drawGrid(matrix)

		// Calculate and render the player's score
		score := initialTwosCount - twosCount
		scoreText := fmt.Sprintf("Score: %d", score)
		if err := g.renderText(scoreText, 20, 20, 24); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
		}

		g.renderer.Present() // Present the renderer

		sdl.Delay(16) // Delay to control frame rate
	}
}