package main

import (
	"log"
	agent "main/Agent"
	"main/board"
	"main/game"
	"os"
	"sync"
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
				// Move the piece
				game.MoveThePiece(fromX, fromY, board.X, board.Y, screen)

				isSelected = false
			} else if game.ValidSelectCheck(board.X, board.Y) {
				if game.CurrentPlayer == board.CIRCLE {
					return
				}
				// Select the piece if it belongs to the current player
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

	var gameOver = false

	var wg sync.WaitGroup

	var gameStatus int

	wg.Add(5)

	go func() {
		defer wg.Done()
		for {
			if gameOver {
				return
			}

			playerInteraction(screen)

			time.Sleep(100 * time.Millisecond)
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
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		for {
			if gameOver {
				return
			}
			if game.CurrentPlayer == board.CIRCLE {
				agent.AgentAction(screen, board.CIRCLE)
			}

			// if game.CurrentPlayer == board.TRIANGLE {
			// 	agent.AgentAction(screen, board.TRIANGLE)
			// }

			time.Sleep(100 * time.Millisecond)
		}

	}()

	go func() {
		defer wg.Done()
		for {
			if gameOver {
				return
			}

			game.DeathCheck(screen)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			if gameOver {
				return
			}

			board.RenderBoard(screen, game.CurrentPlayer)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()

	board.GameOverMessage(screen, gameStatus)
}
