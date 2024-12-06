package main

import (
	"log"
	"main/board"
	"main/game"
	"os"

	"github.com/gdamore/tcell/v2"
)

var Y, X = 0, 0
var fromY, fromX int
var isSelected = false

func playerInteraction(screen tcell.Screen) {
	ev := screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		switch ev.Key() {
		case tcell.KeyEscape:
			screen.Fini()
			os.Exit(0)
			return // Exit on ESC
		case tcell.KeyUp:
			X = (X - 1 + board.BOARD_SIZE) % board.BOARD_SIZE
		case tcell.KeyDown:
			X = (X + 1) % board.BOARD_SIZE
		case tcell.KeyLeft:
			Y = (Y - 1 + board.BOARD_SIZE) % board.BOARD_SIZE
		case tcell.KeyRight:
			Y = (Y + 1) % board.BOARD_SIZE
		case tcell.KeyEnter:
			if isSelected {
				game.MoveThePiece(fromX, fromY, X, Y, screen)

				isSelected = false

			} else {
				// Select the current piece (only if it belongs to the current player)
				if game.ValidSelectCheck(X, Y) {
					fromY, fromX = Y, X
					isSelected = true
				}
			}

		}
	case *tcell.EventResize:
		screen.Sync()
	}
}

func main() {
	var err error
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}

	board.CreateBoard()

	for {
		board.RenderBoard(screen, X, Y, game.CurrentPlayer)

		playerInteraction(screen)

		deathValues := game.DeathCheck()
		if deathValues[0] != [2]int{-1, -1} {
			board.RemovePiece(deathValues)
		}

		game.GameOverCheck(screen)

	}
}
