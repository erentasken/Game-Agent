package minimax

import "main/board"

const (
	borderPenalty           = -1
	cornerPenalty           = -2
	centerReward            = 3
	adjacentPairBonus       = 3
	nearBorderMultiplier    = 2
	centerControlMultiplier = 3
	pieceCountMultiplier    = 5
	killBonus               = 5
)

func EvaluateBoard(player board.Element) int {
	// Scoring constants for flexibility

	// Variables to store the evaluation score
	var score int
	var nearBorderCount, centerControlCount int

	// Store positions of all circles and triangles
	var circlePositions, trianglePositions [][2]int

	// Analyze the board
	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {
			switch MockBoard[i][j] {
			case board.CIRCLE:
				circlePositions = append(circlePositions, [2]int{i, j})
			case board.TRIANGLE:
				trianglePositions = append(trianglePositions, [2]int{i, j})
			}

			// Update scoring for the current player's pieces
			if MockBoard[i][j] == player {
				// Border and corner penalties
				if isOnBorder(i, j) {
					nearBorderCount++
					if isOnCorner(i, j) {
						score += cornerPenalty
					} else {
						score += borderPenalty
					}
				}

				// Center control rewards
				if isNearCenter(i, j) {
					centerControlCount++
					score += centerReward
				}
			}
		}
	}

	// Check adjacent circle-triangle pairs and adjust score
	for _, circlePos := range circlePositions {
		for _, trianglePos := range trianglePositions {
			if manhattanDistance(circlePos, trianglePos) == 1 {
				evaluateAdjacency(&score, player, circlePos, trianglePos, circlePositions, trianglePositions)
			}
		}
	}

	// Apply multipliers for near-border and center control counts
	score += nearBorderCount * nearBorderMultiplier
	score += centerControlCount * centerControlMultiplier

	// Add score based on the number of pieces the player has
	if player == board.TRIANGLE {
		score += MockTriangleNum * pieceCountMultiplier
	} else {
		score += MockCircleNum * pieceCountMultiplier
	}

	deathValues := deathCoordinates()
	if deathValues[0][0] != -1 {
		for _, v := range deathValues {
			if MockBoard[v[0]][v[1]] == board.TRIANGLE {
				if player == board.CIRCLE {
					score += killBonus
				}
			} else {
				if player == board.TRIANGLE {
					score += killBonus
				}
			}
		}
	}

	return score
}

// Helper to check if a position is on the border
func isOnBorder(i, j int) bool {
	return i == 0 || i == board.BOARD_SIZE-1 || j == 0 || j == board.BOARD_SIZE-1
}

// Helper to check if a position is on a corner
func isOnCorner(i, j int) bool {
	return (i == 0 || i == board.BOARD_SIZE-1) && (j == 0 || j == board.BOARD_SIZE-1)
}

// Helper to check if a position is near the center
func isNearCenter(i, j int) bool {
	center := board.BOARD_SIZE / 2
	return abs(i-center) <= 1 && abs(j-center) <= 1
}

// Helper to evaluate adjacency logic
func evaluateAdjacency(score *int, player board.Element, circlePos, trianglePos [2]int, circlePositions, trianglePositions [][2]int) {
	// Calculate total Manhattan distances
	circleTotalDistance := calculateTotalDistance(circlePos, circlePositions)
	triangleTotalDistance := calculateTotalDistance(trianglePos, trianglePositions)

	// Adjust score based on which type's pieces are closer
	if circleTotalDistance < triangleTotalDistance {
		if player == board.CIRCLE {
			*score += adjacentPairBonus // Favor circles
		} else {
			*score -= adjacentPairBonus // Disfavor triangles
		}
	} else {
		if player == board.TRIANGLE {
			*score += adjacentPairBonus // Favor triangles
		} else {
			*score -= adjacentPairBonus // Disfavor circles
		}
	}
}

// Helper to calculate Manhattan distance
func manhattanDistance(pos1, pos2 [2]int) int {
	return abs(pos1[0]-pos2[0]) + abs(pos1[1]-pos2[1])
}

// Helper to calculate total Manhattan distance for all positions of a type
func calculateTotalDistance(refPos [2]int, positions [][2]int) int {
	totalDistance := 0
	for _, pos := range positions {
		totalDistance += manhattanDistance(refPos, pos)
	}
	return totalDistance
}

// Helper for absolute value
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func deathCoordinates() [][2]int {

	for i := 0; i < board.BOARD_SIZE; i++ {
		for j := 0; j < board.BOARD_SIZE; j++ {

			//a piece in between the different pieces, horizontal
			if j < board.BOARD_SIZE-2 {
				if board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.TRIANGLE ||
					board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.CIRCLE {
					return [][2]int{{i, j + 1}}
				}
			}

			//a piece in between the different pieces, vertical
			if i < board.BOARD_SIZE-2 {
				if board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.TRIANGLE ||
					board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.CIRCLE {
					return [][2]int{{i + 1, j}}
				}
			}

			//two pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-3 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.TRIANGLE {
					return [][2]int{{i, j + 1}, {i, j + 2}}
				}
			}

			//two pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-3 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.TRIANGLE {
					return [][2]int{{i + 1, j}, {i + 2, j}}
				}
			}

			//three pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-4 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.TRIANGLE && board.Board[i][j+4] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.CIRCLE && board.Board[i][j+4] == board.TRIANGLE {
					return [][2]int{{i, j + 1}, {i, j + 2}, {i, j + 3}}
				}
			}

			//three pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-4 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.TRIANGLE && board.Board[i+4][j] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.CIRCLE && board.Board[i+4][j] == board.TRIANGLE {
					return [][2]int{{i + 1, j}, {i + 2, j}, {i + 3, j}}
				}
			}

			//four pieces in between the different pieces, horizontal
			if j < board.BOARD_SIZE-5 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.TRIANGLE && board.Board[i][j+4] == board.TRIANGLE && board.Board[i][j+5] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.CIRCLE && board.Board[i][j+4] == board.CIRCLE && board.Board[i][j+5] == board.TRIANGLE {
					return [][2]int{{i, j + 1}, {i, j + 2}, {i, j + 3}, {i, j + 4}}
				}
			}

			//four pieces in between the different pieces, vertical
			if i < board.BOARD_SIZE-5 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.TRIANGLE && board.Board[i+4][j] == board.TRIANGLE && board.Board[i+5][j] == board.CIRCLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.CIRCLE && board.Board[i+4][j] == board.CIRCLE && board.Board[i+5][j] == board.TRIANGLE {
					return [][2]int{{i + 1, j}, {i + 2, j}, {i + 3, j}, {i + 4, j}}
				}
			}

			//********************************************************************************************************************

			//upper border four piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.CIRCLE && board.Board[i+4][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.TRIANGLE && board.Board[i+4][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i + 1, j}, {i + 2, j}, {i + 3, j}}
				}
			}

			//lower border four piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.CIRCLE && board.Board[i-2][j] == board.CIRCLE && board.Board[i-3][j] == board.CIRCLE && board.Board[i-4][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.TRIANGLE && board.Board[i-2][j] == board.TRIANGLE && board.Board[i-3][j] == board.TRIANGLE && board.Board[i-4][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i - 1, j}, {i - 2, j}, {i - 3, j}}
				}
			}

			// left border four piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.CIRCLE && board.Board[i][j+4] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.TRIANGLE && board.Board[i][j+4] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j + 1}, {i, j + 2}, {i, j + 3}}
				}
			}

			// right border four piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.CIRCLE && board.Board[i][j-2] == board.CIRCLE && board.Board[i][j-3] == board.CIRCLE && board.Board[i][j-4] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.TRIANGLE && board.Board[i][j-2] == board.TRIANGLE && board.Board[i][j-3] == board.TRIANGLE && board.Board[i][j-4] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j - 1}, {i, j - 2}, {i, j - 3}}
				}
			}

			//upper border three piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.CIRCLE && board.Board[i+3][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.TRIANGLE && board.Board[i+3][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i + 1, j}, {i + 2, j}}
				}
			}

			//lower border three piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.CIRCLE && board.Board[i-2][j] == board.CIRCLE && board.Board[i-3][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.TRIANGLE && board.Board[i-2][j] == board.TRIANGLE && board.Board[i-3][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i - 1, j}, {i - 2, j}}
				}
			}

			// left border three piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.CIRCLE && board.Board[i][j+3] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.TRIANGLE && board.Board[i][j+3] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j + 1}, {i, j + 2}}
				}
			}

			// right border three piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.CIRCLE && board.Board[i][j-2] == board.CIRCLE && board.Board[i][j-3] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.TRIANGLE && board.Board[i][j-2] == board.TRIANGLE && board.Board[i][j-3] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j - 1}, {i, j - 2}}
				}
			}

			// upper border two piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.CIRCLE && board.Board[i+2][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.TRIANGLE && board.Board[i+2][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i + 1, j}}
				}
			}

			//lower border two piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.CIRCLE && board.Board[i-2][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.TRIANGLE && board.Board[i-2][j] == board.CIRCLE {
					return [][2]int{{i, j}, {i - 1, j}}
				}
			}

			// left border two piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.CIRCLE && board.Board[i][j+2] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.TRIANGLE && board.Board[i][j+2] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j + 1}}
				}
			}

			// right border two piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.CIRCLE && board.Board[i][j-2] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.TRIANGLE && board.Board[i][j-2] == board.CIRCLE {
					return [][2]int{{i, j}, {i, j - 1}}
				}
			}

			// upper border one piece
			if i == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i+1][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i+1][j] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}

			//lower border one piece
			if i == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i-1][j] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i-1][j] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}

			// left border one piece
			if j == 0 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j+1] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j+1] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}

			// right border one piece
			if j == board.BOARD_SIZE-1 {
				if board.Board[i][j] == board.CIRCLE && board.Board[i][j-1] == board.TRIANGLE ||
					board.Board[i][j] == board.TRIANGLE && board.Board[i][j-1] == board.CIRCLE {
					return [][2]int{{i, j}}
				}
			}
		}
	}

	return [][2]int{{-1, -1}}
}
