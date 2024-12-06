package agent

import (
	"main/board"
	"main/game"
	"time"

	"math/rand"

	"github.com/gdamore/tcell/v2"
)

var Sync = 0

func AgentAction(screen tcell.Screen, element board.Element) {
	time.Sleep(1000 * time.Millisecond)

	style := tcell.StyleDefault
	style = style.Background(tcell.ColorDarkGoldenrod)
	actions := game.GetPossibleActions(element)
	for _, a := range actions {
		// fmt.Print(a, " ")
		for k := 0; k < 2; k++ {
			screen.SetContent(a.ToY*2+k, a.ToX, ' ', nil, style)
		}
	}
	screen.Show()

	// ev := screen.PollEvent()
	// switch ev := ev.(type) {
	// case *tcell.EventKey:
	// 	if ev.Key() == tcell.KeyEscape {
	// 		screen.Fini()
	// 		os.Exit(0) // Exit the program
	// 	}
	// }

	// Shuffle the actions to add randomness
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(actions), func(i, j int) {
		actions[i], actions[j] = actions[j], actions[i]
	})

	// Attempt to perform two valid moves
	movesPerformed := 0
	for _, action := range actions {
		time.Sleep(1000 * time.Millisecond)

		if game.MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY, screen) {
			movesPerformed++
			if movesPerformed == 2 {
				break // Stop after two successful moves
			}

			if element == board.TRIANGLE {
				if board.TriangleNum == 1 {
					break
				}
			}

			if element == board.CIRCLE {
				if board.CircleNum == 1 {
					break
				}
			}
		}
	}
}
