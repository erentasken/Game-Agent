package agent

import (
	"main/board"
	"main/game"
	"main/minimax"
	"math"
	"math/rand"
	"time"

	"github.com/gdamore/tcell/v2"
)

func AgentAction(screen tcell.Screen, element board.Element) {
	const depth = 10 ^ 17

	actions := game.GetPossibleActions(element)
	if len(actions) == 0 {
		return
	}

	var bestAction game.Action
	var bestActionList []game.Action
	bestEval := math.MinInt32

	for _, action := range actions {
		deepCopyBoard()

		minimax.MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY)

		eval := minimax.Minimax(depth-1, math.MinInt32, math.MaxInt32, false, element)

		if eval > bestEval {
			bestEval = eval
			bestActionList = []game.Action{action}
		} else if eval == bestEval {
			bestActionList = append(bestActionList, action)
		}

	}

	bestAction = bestActionList[rand.Intn(len(bestActionList))]

	time.Sleep(300 * time.Millisecond)

	game.MoveThePiece(bestAction.FromX, bestAction.FromY, bestAction.ToX, bestAction.ToY, screen)

	AgentAction(screen, element)
}

func deepCopyBoard() {
	minimax.MockTurnCounter = game.TurnCounter
	minimax.MockMoveCounter = board.MoveCounter
	minimax.MockCurrentPlayer = game.CurrentPlayer
	minimax.MockCircleNum = board.CircleNum
	minimax.MockTriangleNum = board.TriangleNum

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			minimax.MockBoard[i][j] = board.Board[i][j]
		}
	}
}
