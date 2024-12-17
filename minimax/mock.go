package minimax

import (
	"main/board"
)

func MoveThePiece(fromX, fromY, X, Y int, boardState BoardState) BoardState {
	boardState.Board[X][Y] = boardState.Board[fromX][fromY]
	boardState.Board[fromX][fromY] = board.EMPTY

	boardState = DeathCheck(boardState)

	return boardState
}

func DeathCheck(boardState BoardState) BoardState {
	deathValues := deathCoordinates(boardState.Board)

	for _, v := range deathValues {
		if v == [2]int{-1, -1} {
			continue
		}

		if boardState.Board[v[0]][v[1]] == board.TRIANGLE {
			boardState.TriangleNum--
		} else {
			boardState.CircleNum--
		}

		boardState.Board[v[0]][v[1]] = board.EMPTY
	}

	return boardState
}

func deathCoordinates(MockBoard [7][7]board.Element) [][2]int {

	var result [][2]int // Initialize the result list to store coordinates

	result = append(result, [2]int{-1, -1})

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {

			//a piece in between the different pieces, horizontal
			if j < board.BOARD_SIZE-2 {
				if MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.TRIANGLE ||
					MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.CIRCLE {
					result = append(result, [2]int{i, j + 1})
				}
			}

			//a piece in between the different pieces, vertical
			if i < board.BOARD_SIZE-2 {
				if MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.CIRCLE {
					result = append(result, [2]int{i + 1, j})
				}
			}

			//two pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-3 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.TRIANGLE {
					result = append(result, [2]int{i, j + 1}, [2]int{i, j + 2})
				}
			}

			//two pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-3 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.TRIANGLE {
					result = append(result, [2]int{i + 1, j}, [2]int{i + 2, j})
				}
			}

			//three pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-4 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.TRIANGLE && MockBoard[i][j+4] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.CIRCLE && MockBoard[i][j+4] == board.TRIANGLE {
					result = append(result, [2]int{i, j + 1}, [2]int{i, j + 2}, [2]int{i, j + 3})
				}
			}

			//three pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-4 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.TRIANGLE && MockBoard[i+4][j] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.CIRCLE && MockBoard[i+4][j] == board.TRIANGLE {
					result = append(result, [2]int{i + 1, j}, [2]int{i + 2, j}, [2]int{i + 3, j})
				}
			}

			//four pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-5 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.TRIANGLE && MockBoard[i][j+4] == board.TRIANGLE && MockBoard[i][j+5] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.CIRCLE && MockBoard[i][j+4] == board.CIRCLE && MockBoard[i][j+5] == board.TRIANGLE {
					result = append(result, [2]int{i, j + 1}, [2]int{i, j + 2}, [2]int{i, j + 3}, [2]int{i, j + 4})
				}
			}

			//four pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-5 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.TRIANGLE && MockBoard[i+4][j] == board.TRIANGLE && MockBoard[i+5][j] == board.CIRCLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.CIRCLE && MockBoard[i+4][j] == board.CIRCLE && MockBoard[i+5][j] == board.TRIANGLE {
					result = append(result, [2]int{i + 1, j}, [2]int{i + 2, j}, [2]int{i + 3, j}, [2]int{i + 4, j})
				}
			}

			//********************************************************************************************************************

			//upper border four piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.CIRCLE && MockBoard[i+4][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.TRIANGLE && MockBoard[i+4][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i + 1, j}, [2]int{i + 2, j}, [2]int{i + 3, j})
				}
			}

			//lower border four piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.CIRCLE && MockBoard[i-2][j] == board.CIRCLE && MockBoard[i-3][j] == board.CIRCLE && MockBoard[i-4][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.TRIANGLE && MockBoard[i-2][j] == board.TRIANGLE && MockBoard[i-3][j] == board.TRIANGLE && MockBoard[i-4][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i - 1, j}, [2]int{i - 2, j}, [2]int{i - 3, j})
				}
			}

			// left border four piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.CIRCLE && MockBoard[i][j+4] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.TRIANGLE && MockBoard[i][j+4] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j + 1}, [2]int{i, j + 2}, [2]int{i, j + 3})
				}
			}

			// right border four piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.CIRCLE && MockBoard[i][j-2] == board.CIRCLE && MockBoard[i][j-3] == board.CIRCLE && MockBoard[i][j-4] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.TRIANGLE && MockBoard[i][j-2] == board.TRIANGLE && MockBoard[i][j-3] == board.TRIANGLE && MockBoard[i][j-4] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j - 1}, [2]int{i, j - 2}, [2]int{i, j - 3})
				}
			}

			//upper border three piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.CIRCLE && MockBoard[i+3][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.TRIANGLE && MockBoard[i+3][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i + 1, j}, [2]int{i + 2, j})
				}
			}

			//lower border three piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.CIRCLE && MockBoard[i-2][j] == board.CIRCLE && MockBoard[i-3][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.TRIANGLE && MockBoard[i-2][j] == board.TRIANGLE && MockBoard[i-3][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i - 1, j}, [2]int{i - 2, j})
				}
			}

			// left border three piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.CIRCLE && MockBoard[i][j+3] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.TRIANGLE && MockBoard[i][j+3] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j + 1}, [2]int{i, j + 2})
				}
			}

			// right border three piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.CIRCLE && MockBoard[i][j-2] == board.CIRCLE && MockBoard[i][j-3] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.TRIANGLE && MockBoard[i][j-2] == board.TRIANGLE && MockBoard[i][j-3] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j - 1}, [2]int{i, j - 2})
				}
			}

			// upper border two piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.CIRCLE && MockBoard[i+2][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.TRIANGLE && MockBoard[i+2][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i + 1, j})
				}
			}

			//lower border two piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.CIRCLE && MockBoard[i-2][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.TRIANGLE && MockBoard[i-2][j] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i - 1, j})
				}
			}

			// left border two piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.CIRCLE && MockBoard[i][j+2] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.TRIANGLE && MockBoard[i][j+2] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j + 1})
				}
			}

			// right border two piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.CIRCLE && MockBoard[i][j-2] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.TRIANGLE && MockBoard[i][j-2] == board.CIRCLE {
					result = append(result, [2]int{i, j}, [2]int{i, j - 1})
				}
			}

			// upper border one piece
			if i == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i+1][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i+1][j] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}

			//lower border one piece
			if i == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i-1][j] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i-1][j] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}

			// left border one piece
			if j == 0 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j+1] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j+1] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}

			// right border one piece
			if j == board.BOARD_SIZE-1 {
				if MockBoard[i][j] == board.CIRCLE && MockBoard[i][j-1] == board.TRIANGLE ||
					MockBoard[i][j] == board.TRIANGLE && MockBoard[i][j-1] == board.CIRCLE {
					result = append(result, [2]int{i, j})
				}
			}
		}
	}

	return result
}

type Action struct {
	FromX, FromY, ToX, ToY int
}

func GetPossibleActions(entity board.Element, MockBoard [7][7]board.Element) []Action {
	var actionList []Action

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			if MockBoard[i][j] == entity {

				if j < board.BOARD_SIZE-1 && MockBoard[i][j+1] == board.EMPTY {
					actionList = append(actionList, Action{i, j, i, j + 1})
				}

				if j > 0 && MockBoard[i][j-1] == board.EMPTY {
					actionList = append(actionList, Action{i, j, i, j - 1})
				}

				if i > 0 && MockBoard[i-1][j] == board.EMPTY {
					actionList = append(actionList, Action{i, j, i - 1, j})
				}

				if i < board.BOARD_SIZE-1 && MockBoard[i+1][j] == board.EMPTY {
					actionList = append(actionList, Action{i, j, i + 1, j})
				}
			}
		}
	}
	return actionList
}
