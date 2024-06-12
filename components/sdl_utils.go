package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Initialize SDL2 and SDL_ttf
func initializeSDL() error {
	var err error
	var sdlFlags uint32 = sdl.INIT_EVERYTHING

	// Initialize SDL with all subsystems
	if err = sdl.Init(sdlFlags); err != nil {
		return fmt.Errorf("Error initializing SDL2: %v", err)
	}

	// Initialize SDL_ttf
	if err = ttf.Init(); err != nil {
		return fmt.Errorf("Error initializing SDL_ttf: %v", err)
	}

	return nil
}

// Quit SDL2 and SDL_ttf
func closeSDL() {
	ttf.Quit()
	sdl.Quit()
}
