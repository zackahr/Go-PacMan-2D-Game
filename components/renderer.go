package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
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

func (g *game) drawGrid(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			x := col * cellSize
			y := row * cellSize
			if matrix[row][col] == 1 {
				g.drawLine(x, y, x+cellSize, y)         
				g.drawLine(x, y, x, y+cellSize)             
				g.drawLine(x+cellSize, y, x+cellSize, y+cellSize) 
				g.drawLine(x, y+cellSize, x+cellSize, y+cellSize) 			
			}
		}
	}

	// Draw the player
	playerX := g.playerX * cellSize
	playerY := g.playerY * cellSize
	g.renderer.SetDrawColor(255, 255, 0, 255) // Yellow color
	g.drawCircle(playerX+cellSize/2, playerY+cellSize/2, circleRadius)
}


// Function to render text
func (g *game) renderText(text string, x, y int) error {
	font, err := ttf.OpenFont("fonts/freesansbold.ttf", g.fontSize)
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
