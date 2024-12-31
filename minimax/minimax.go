package minimax

import (
	"fmt"
	"log"
	"main/board"
	"main/game"
	"math"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"golang.org/x/exp/rand"
)

type BoardState struct {
	Board       [board.BOARD_SIZE][board.BOARD_SIZE]board.Element
	CircleNum   int
	TriangleNum int
	MoveCounter int
}

func AgentAction(screen tcell.Screen, element board.Element) {
	const depth = 4

	//board state from main game.
	boardState := BoardState{
		Board:       board.Board,
		CircleNum:   board.CircleNum,
		TriangleNum: board.TriangleNum,
		MoveCounter: board.MoveCounter,
	}

	actions := GetPossibleActions(element, board.Board)
	if len(actions) == 0 {
		return
	}

	var bestActions [2]Action
	bestEval := math.MinInt32

	var bestFound = false

	var validActions [][2]Action

	if element == board.CIRCLE && boardState.CircleNum == 1 || element == board.TRIANGLE && boardState.TriangleNum == 1 {
		for i := 0; i < len(actions); i++ {
			action := actions[i]

			copyBoardState := CopyBoardState(boardState)

			copyBoardState = MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY, copyBoardState)

			eval := Minimax(depth-1, math.MinInt32, math.MaxInt32, false, copyBoardState, getOpponent(element))

			if eval > bestEval {
				bestFound = true
				bestEval = eval
				bestActions = [2]Action{action, action}
			}

			validActions = append(validActions, [2]Action{action, action})
		}
	} else {
		for i := 0; i < len(actions); i++ {
			for j := i + 1; j < len(actions); j++ {
				action1 := actions[i]
				action2 := actions[j]

				if (action1.FromX == action2.FromX && action1.FromY == action2.FromY) ||
					(action1.ToX == action2.ToX && action1.ToY == action2.ToY) ||
					(action1.ToX == action2.FromX && action1.ToY == action2.FromY) ||
					(action1.FromX == action2.ToX && action1.FromY == action2.ToY) {
					continue
				}

				copyBoardState := CopyBoardState(boardState)

				copyBoardState = MoveThePiece(action1.FromX, action1.FromY, action1.ToX, action1.ToY, copyBoardState)

				copyBoardState = MoveThePiece(action2.FromX, action2.FromY, action2.ToX, action2.ToY, copyBoardState)

				eval := Minimax(depth-1, math.MinInt32, math.MaxInt32, false, copyBoardState, getOpponent(element))

				if eval > bestEval {

					bestFound = true
					bestEval = eval
					bestActions = [2]Action{action1, action2}
				}

				validActions = append(validActions, [2]Action{action1, action2})
			}
		}
	}

	if !bestFound {
		LogError("No valid actions found\n", "./log/valid_combinations.log")
		randIndex := rand.Intn(len(validActions))
		bestActions = validActions[randIndex]
	}

	for _, action := range bestActions {
		if bestFound {
			LogError("\n\nBest action: "+fmt.Sprintf("%v", action)+"\n", "./log/valid_combinations.log")
		}
		game.MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY, screen)
		time.Sleep(350 * time.Millisecond)
	}
}

func Minimax(depth int, alpha, beta int, isMaximizing bool, boardState BoardState, player board.Element) int {
	if depth == 0 {
		return EvaluateBoard(player, boardState)
	}

	actions := GetPossibleActions(player, boardState.Board)
	if len(actions) == 0 {
		return EvaluateBoard(player, boardState)
	}

	if isMaximizing {
		return evaluateActionCombinations(depth, alpha, beta, boardState, actions, true, player, Max)
	} else {
		return evaluateActionCombinations(depth, alpha, beta, boardState, actions, false, player, Min)
	}
}

func evaluateActionCombinations(depth, alpha, beta int, boardState BoardState, actions []Action, isMaximizing bool, player board.Element, compare func(a, b int) int) int {
	bestEval := func() int {
		if isMaximizing {
			return math.MinInt32
		}
		return math.MaxInt32
	}()

	if player == board.CIRCLE && boardState.CircleNum == 1 || player == board.TRIANGLE && boardState.TriangleNum == 1 {
		for i := 0; i < len(actions); i++ {
			action := actions[i]

			copyBoardState := CopyBoardState(boardState)

			copyBoardState = MoveThePiece(action.FromX, action.FromY, action.ToX, action.ToY, copyBoardState)

			eval := Minimax(depth-1, alpha, beta, !isMaximizing, copyBoardState, getOpponent(player))

			bestEval = compare(bestEval, eval)

			if isMaximizing {
				alpha = Max(alpha, bestEval)
				if beta <= alpha {
					break
				}
			} else {
				beta = Min(beta, bestEval)
				if beta <= alpha {
					break
				}
			}
		}
		return bestEval
	}

	for i := 0; i < len(actions); i++ {
		for j := i + 1; j < len(actions); j++ {
			action1 := actions[i]
			action2 := actions[j]

			if action1.FromX == action2.FromX && action1.FromY == action2.FromY {
				continue
			}

			copyBoardState := CopyBoardState(boardState)

			copyBoardState = MoveThePiece(action1.FromX, action1.FromY, action1.ToX, action1.ToY, copyBoardState)
			copyBoardState = MoveThePiece(action2.FromX, action2.FromY, action2.ToX, action2.ToY, copyBoardState)

			eval := Minimax(depth-1, alpha, beta, !isMaximizing, copyBoardState, getOpponent(player))

			bestEval = compare(bestEval, eval)

			if isMaximizing {
				alpha = Max(alpha, bestEval)
				if beta <= alpha {
					break
				}
			} else {
				beta = Min(beta, bestEval)
				if beta <= alpha {
					break
				}
			}
		}
	}

	return bestEval
}

func getOpponent(player board.Element) board.Element {
	if player == board.TRIANGLE {
		return board.CIRCLE
	}
	return board.TRIANGLE
}

func LogError(content, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Error opening file: %v", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(content); err != nil {
		log.Printf("Error writing to file: %v", err)
		return
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CopyBoardState(boardState BoardState) BoardState {
	var copyBoard [board.BOARD_SIZE][board.BOARD_SIZE]board.Element
	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			copyBoard[i][j] = boardState.Board[i][j]
		}
	}
	return BoardState{
		Board:       copyBoard,
		CircleNum:   boardState.CircleNum,
		TriangleNum: boardState.TriangleNum,
		MoveCounter: boardState.MoveCounter,
	}
}
