package main

import (
	"fmt"
	"os"
)

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

	g.run(matrix2)
}
