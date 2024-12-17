package main

import (
	"log"
	"main/board"
	"main/game"
	"main/minimax"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
)

var fromY, fromX int
var isSelected = false

func playerInteraction(screen tcell.Screen) {
	ev := screen.PollEvent()
	switch ev := ev.(type) {
	case *tcell.EventKey:
		if ev.Key() == tcell.KeyEscape {
			screen.Fini()
			os.Exit(0) // Exit the program
		}

		switch ev.Key() {
		case tcell.KeyUp:
			board.X = (board.X - 1 + board.BOARD_SIZE) % board.BOARD_SIZE
			board.RenderBoard(screen, game.CurrentPlayer)
		case tcell.KeyDown:
			board.X = (board.X + 1) % board.BOARD_SIZE
			board.RenderBoard(screen, game.CurrentPlayer)
		case tcell.KeyLeft:
			board.Y = (board.Y - 1 + board.BOARD_SIZE) % board.BOARD_SIZE
			board.RenderBoard(screen, game.CurrentPlayer)
		case tcell.KeyRight:
			board.Y = (board.Y + 1) % board.BOARD_SIZE
			board.RenderBoard(screen, game.CurrentPlayer)
		case tcell.KeyEnter:
			if isSelected {
				game.MoveThePiece(fromX, fromY, board.X, board.Y, screen)

				if game.CurrentPlayer == board.CIRCLE {
					minimax.AgentAction(screen, board.CIRCLE)
				}

				isSelected = false

			} else if game.ValidSelectCheck(board.X, board.Y) {
				if game.CurrentPlayer == board.CIRCLE {
					return
				}

				fromY, fromX = board.Y, board.X
				isSelected = true
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

	game.CurrentPlayer = board.CIRCLE

	board.RenderBoard(screen, game.CurrentPlayer)

	minimax.AgentAction(screen, board.CIRCLE)

	for {
		if game.GameOver {
			break
		}

		playerInteraction(screen)

		time.Sleep(100 * time.Millisecond)
	}

	board.GameOverMessage(screen)
}
