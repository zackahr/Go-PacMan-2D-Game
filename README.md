# PacMan Game

## Overview

This project is a 2D game inspired by the classic Pac-Man. The game is developed using Go (Golang) and SDL2, a powerful cross-platform library for handling multimedia, input, and graphics. In this game, players control Pac-Man, navigating through a maze to collect collectibles while avoiding enemies.

## Technologies Used

- **Go (Golang)**: A statically typed, compiled programming language designed for simplicity and performance. It is used for writing the game logic and handling the main game loop.
- **SDL2**: The Simple DirectMedia Layer (SDL) is a cross-platform library that provides low-level access to audio, keyboard, mouse, joystick, and graphics hardware via OpenGL and Direct3D. SDL2 is used to create the game window, render graphics, and manage input events.

## Game Description

PacMan is a classic arcade game where the player navigates through a maze, attempting to eat all the collectibles while avoiding enemies that try to catch them. The game includes the following features:

- **Player Control**: Move Pac-Man using the keyboard arrow keys.
- **Collectibles**: Collect items scattered throughout the maze to gain points.
- **Enemies**: Avoid enemies that patrol the maze. If an enemy catches Pac-Man, the game is over.
- **Grid-Based Movement**: The game world is a grid, with Pac-Man and the enemies moving from cell to cell.

### How to Play

1. **Start the Game**: Upon starting the game, an image is displayed. Press any key to begin the game.
2. **Control Pac-Man**: Use the `W`, `A`, `S`, and `D` keys to move Pac-Man up, left, down, and right, respectively.
3. **Collect Items**: Navigate Pac-Man through the maze to collect all the collectibles while avoiding enemies.
4. **Avoid Enemies**: If an enemy catches Pac-Man, the game is over.

## Installation and Setup

### Prerequisites

- **Go**: Make sure you have Go installed. You can download it from the [official site](https://golang.org/dl/).
- **SDL2**: Install SDL2 and the SDL2_image extension. You can find installation instructions on the [SDL2 website](https://www.libsdl.org/download-2.0.php).

### Steps

1. Clone the repository:

   ```bash
   git clone https://github.com/zackahr/Go-PacMan-2D-Game.git
   cd Go-PacMan-2D-Game
2. Build and run the game:
   ```bash
   make run
### Project Structure
  
- **main.go** : The entry point of the application.
- **game.go** : Contains the game logic, including initialization, event handling, and rendering.
- **assets**  : Directory containing game assets like images.
- **fonts**   : Directory containing the fonts used for text game.
- **Makefile**: File that help to run the project directly. 
   
