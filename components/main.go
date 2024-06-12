package main

import (
	"fmt"
	"os"
)

func (g *game) renderInitialImage(imagePath string) error {
	imageTexture, err := g.loadImage(imagePath)
	if err != nil {
		return fmt.Errorf("failed to load image: %v", err)
	}
	defer imageTexture.Destroy()

	// Clear the renderer
	g.renderer.Clear()

	// Draw the image texture
	g.renderer.Copy(imageTexture, nil, nil)

	// Present the renderer
	g.renderer.Present()

	return nil
}

func main() {
	var err error

	defer closeSDL() // Ensure SDL is closed on exit
	if err = initializeSDL(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err) // Print error to stderr
		return
	}

	g := newGame() // Create a new game instance
	defer g.close() // Ensure game resources are cleaned up on exit
	if err = g.init(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err) // Print error to stderr
		return
	}
	if err := g.renderInitialImage("assets/PacMan.jpg"); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render initial image: %v\n", err)
		return
	}

	g.run(matrix2)
}
