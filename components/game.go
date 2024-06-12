package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "Game Window"
	circleRadius = 3
	cellSize     = 20
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

// Main game loop
func (g *game) run(matrix [][]int) {
	running := true
	for running {
		// Handle events
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				running = false
			case *sdl.KeyboardEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					running = false
				}
				if t.Type == sdl.KEYDOWN {
					switch t.Keysym.Sym {
					case sdl.K_w:
						if g.playerY > 0 && matrix[g.playerY-1][g.playerX] != 1 {
							g.playerY--
						}
					case sdl.K_a:
						if g.playerX > 0 && matrix[g.playerY][g.playerX-1] != 1 {
							g.playerX--
						}
					case sdl.K_s:
						if g.playerY < len(matrix)-1 && matrix[g.playerY+1][g.playerX] != 1 {
							g.playerY++
						}
					case sdl.K_d:
						if g.playerX < len(matrix[0])-1 && matrix[g.playerY][g.playerX+1] != 1 {
							g.playerX++
						}
					}
				}
			}
		}

		// Render the scene
		g.renderer.SetDrawColor(0, 0, 0, 255)
		g.renderer.Clear()

		g.drawGrid(matrix)

		g.renderer.Present() // Present the renderer

		sdl.Delay(16) // Delay to control frame rate
	}
}
