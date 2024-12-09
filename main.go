package main

import (
	"log"
	agent "main/Agent"
	"main/board"
	"main/game"
	"os"
	"sync"

	"github.com/gdamore/tcell/v2"
)

var Y, X = 0, 0
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

		// if game.CurrentPlayer == board.CIRCLE {
		// 	return
		// }

		switch ev.Key() {
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
				// Move the piece
				game.MoveThePiece(fromX, fromY, X, Y, screen)

				isSelected = false
			} else if game.ValidSelectCheck(X, Y) {
				// Select the piece if it belongs to the current player
				fromY, fromX = Y, X
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
	board.RenderBoard(screen, X, Y, game.CurrentPlayer)

	game.CurrentPlayer = board.CIRCLE

	var gameOver = false

	var wg sync.WaitGroup

	var gameStatus int

	wg.Add(4)

	go func() {
		defer wg.Done() // Mark this Goroutine as done when it finishes
		for {
			if gameOver {
				return
			}
			board.RenderBoard(screen, X, Y, game.CurrentPlayer)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			playerInteraction(screen)
			if gameOver {
				return
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if gameOver {
				return
			}
			game.DeathCheck()
		}

	}()

	go func() {
		defer wg.Done()
		for {
			gameStatus = game.GameOverCheck(screen)
			if gameStatus != 2 {
				gameOver = true
				return
			}
		}
	}()

	go func() {
		for {
			if game.CurrentPlayer == board.CIRCLE {
				agent.AgentAction(screen, board.CIRCLE)
			}
			// if game.CurrentPlayer == board.TRIANGLE {
			// 	agent.AgentAction(screen, board.TRIANGLE)
			// }
		}
	}()

	wg.Wait()

	board.GameOverMessage(screen, gameStatus)
}
