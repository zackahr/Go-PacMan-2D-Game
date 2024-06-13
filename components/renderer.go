package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"github.com/veandco/go-sdl2/img"
	"os"
)

func (g *game) drawLine(x1, y1, x2, y2 int) {
	g.renderer.SetDrawColor(0, 255, 255, 255) // White color for the lines
	g.renderer.DrawLine(int32(x1), int32(y1), int32(x2), int32(y2))
}

// Function to draw a circle
func (g *game) drawCircle(x, y, radius int) {
	for w := 0; w < radius*2; w++ {
		for h := 0; h < radius*2; h++ {
			dx := radius - w
			dy := radius - h
			if (dx*dx + dy*dy) <= (radius * radius) {
				g.renderer.DrawPoint(int32(x+dx), int32(y+dy))
			}
		}
	}
}

func (g *game) loadImage(filePath string) (*sdl.Texture, error) {
	imgSurface, err := img.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %v", err)
	}
	defer imgSurface.Free()

	texture, err := g.renderer.CreateTextureFromSurface(imgSurface)
	if err != nil {
		return nil, fmt.Errorf("failed to create texture: %v", err)
	}

	return texture, nil
}

func (g *game) renderInitializationImage() {
	imagePath := "assets/PacMan.jpg"
	imageTexture, err := g.loadImage(imagePath)
	if err != nil {
		fmt.Printf("Failed to load image: %v\n", err)
		return
	}

	// Define the size and position of the image
	destRect := sdl.Rect{
		W: 160,  // Width of the image
		H: 160,  // Height of the image
		X: windowWidth/2 - 80,  // X position on the screen
		Y: windowHeight/2 - 80,  // Y position on the screen
	}

	// Render the image
	g.renderer.Clear()
	g.renderer.Copy(imageTexture, nil, &destRect)

	// Render text under the image
	text := "Pac Man Gooo!!!!"

	if err := g.renderText(text, windowWidth / 2 - 100, windowHeight / 2 + 100, 24); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to render text: %v\n", err)
	}

	// Present the renderer
	g.renderer.Present()

	// Clean up
	imageTexture.Destroy()
}

func (g *game) renderGameGrid(matrix [][]int, startX, startY int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x := startX + col*cellSize
			y := startY + row*cellSize
			if matrix[row][col] == 1 {
				g.drawLine(x, y, x+cellSize, y)
				g.drawLine(x, y, x, y+cellSize)
				g.drawLine(x+cellSize, y, x+cellSize, y+cellSize)
				g.drawLine(x, y+cellSize, x+cellSize, y+cellSize)
			} else if matrix[row][col] == 2 {
				g.renderer.SetDrawColor(255, 25, 0, 255)
				g.drawCircle(x+cellSize/2, y+cellSize/2, circleRadius/2)
			}
		}
	}
}


func (g *game) initializePlayerPosition(matrix [][]int) {
	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 'P' {
				g.playerX = col
				g.playerY = row
				matrix[row][col] = 0 // Replace 'P' with 0 after setting the position
				return
			}
		}
	}
}

func (g *game) loadPlayerImage(filePath string) (*sdl.Texture, error) {
	imgSurface, err := img.Load(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load image: %v", err)
	}
	defer imgSurface.Free()

	texture, err := g.renderer.CreateTextureFromSurface(imgSurface)
	if err != nil {
		return nil, fmt.Errorf("failed to create texture: %v", err)
	}

	return texture, nil
}

func (g *game) renderPlayer(startX, startY int) {
	playerX := startX + g.playerX*cellSize
	playerY := startY + g.playerY*cellSize
	g.renderer.SetDrawColor(255, 255, 0, 255) // Yellow color
	g.drawCircle(playerX+cellSize/2, playerY+cellSize/2, circleRadius)
}

func (g *game) drawGrid(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])
	startX := (windowWidth - cols*cellSize) / 2
	startY := (windowHeight - rows*cellSize) / 2

	if g.initializing {
		g.renderInitializationImage()
	} else {
		g.renderGameGrid(matrix, startX, startY)
		g.renderPlayer(startX, startY)
	}
}

// Function to render text
func (g *game) renderText(text string, x, y, fontSize int) error {
	font, err := ttf.OpenFont("fonts/freesansbold.ttf", fontSize)
	if err != nil {
		return fmt.Errorf("Failed to open font: %v", err)
	}
	defer font.Close()

	color := sdl.Color{R: 255, G: 255, B: 255, A: 255} // White color

	surface, err := font.RenderUTF8Blended(text, color)
	if err != nil {
		return fmt.Errorf("Failed to render text: %v", err)
	}
	defer surface.Free()

	texture, err := g.renderer.CreateTextureFromSurface(surface)
	if err != nil {
		return fmt.Errorf("Failed to create texture: %v", err)
	}
	defer texture.Destroy()

	rect := sdl.Rect{X: int32(x), Y: int32(y), W: surface.W, H: surface.H}
	if err := g.renderer.Copy(texture, nil, &rect); err != nil {
		return fmt.Errorf("Failed to copy texture: %v", err)
	}

	return nil
}