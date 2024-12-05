package main

import (
	"log"
	"main/board"
	"main/game"

	"github.com/gdamore/tcell/v2"
)

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Failed to create screen: %v", err)
	}
	defer screen.Fini()

	if err := screen.Init(); err != nil {
		log.Fatalf("Failed to initialize screen: %v", err)
	}

	board.CreateBoard()

	Y, X := 0, 0
	var selectedY, selectedX int
	isSelected := false
	currentPlayer := board.TRIANGLE

	for {
		board.RenderBoard(screen, X, Y, currentPlayer)

		// Key strokes
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
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
					if board.MovePiece(selectedX, selectedY, X, Y, screen) {
						isSelected = false

						game.SwitchTurn(&currentPlayer)
					}

				} else {
					// Select the current piece (only if it belongs to the current player)
					if board.Board[X][Y] != board.EMPTY && board.Board[X][Y] == currentPlayer {
						selectedY, selectedX = Y, X
						isSelected = true

					}
				}

			}
		case *tcell.EventResize:
			screen.Sync()
		}

		//

		deathValues := game.DeathCheck()

		if deathValues[0] != [2]int{-1, -1} {
			board.RemovePiece(deathValues)
		}

	}
}
