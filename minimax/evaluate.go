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

// func EvaluateBoard(player board.Element) int {
// 	// Variables to store the evaluation score
// 	var score int

// 	// Count pieces near the borders (edges of the board)
// 	var nearBorderCount int
// 	var centerControlCount int

// 	for i := 0; i < board.BOARD_SIZE; i++ {
// 		for j := 0; j < board.BOARD_SIZE; j++ {
// 			if MockBoard[i][j] == player {
// 				// Check if the piece is on the border
// 				if i == 0 || i == board.BOARD_SIZE-1 || j == 0 || j == board.BOARD_SIZE-1 {
// 					nearBorderCount++
// 					if (i == 0 && j == 0) || (i == 0 && j == board.BOARD_SIZE-1) || (i == board.BOARD_SIZE-1 && j == 0) || (i == board.BOARD_SIZE-1 && j == board.BOARD_SIZE-1) {
// 						// Corners give extra points
// 						score -= 2
// 					} else {
// 						// Edges give normal points
// 						score -= 1
// 					}
// 				}

// 				// Check if the piece is near the center (center control)
// 				if (i >= board.BOARD_SIZE/2-1 && i <= board.BOARD_SIZE/2+1) && (j >= board.BOARD_SIZE/2-1 && j <= board.BOARD_SIZE/2+1) {
// 					centerControlCount++
// 					score += 3 // Center pieces give more weight
// 				}
// 			}
// 		}
// 	}

// 	// Reward more for having more pieces near the border (but don't give too much weight)
// 	score += nearBorderCount * 2

// 	// Reward more for having more pieces in the center
// 	score += centerControlCount * 3

// 	// Finally, return the score based on the number of pieces the player has
// 	if player == board.TRIANGLE {
// 		return score + MockTriangleNum*5
// 	}
// 	return score + MockCircleNum*5
// }

// func evaluateBoard(player board.Element) int {
// 	// Variables to store the evaluation score
// 	var score int
// 	var nearBorderCount, centerControlCount int

// 	// Store positions of all circles and triangles
// 	var circlePositions, trianglePositions [][2]int

// 	for i := 0; i < board.BOARD_SIZE; i++ {
// 		for j := 0; j < board.BOARD_SIZE; j++ {
// 			if MockBoard[i][j] == board.CIRCLE {
// 				circlePositions = append(circlePositions, [2]int{i, j})
// 			} else if MockBoard[i][j] == board.TRIANGLE {
// 				trianglePositions = append(trianglePositions, [2]int{i, j})
// 			}

// 			// Evaluation logic for border and center control
// 			if MockBoard[i][j] == player {
// 				if i == 0 || i == board.BOARD_SIZE-1 || j == 0 || j == board.BOARD_SIZE-1 {
// 					nearBorderCount++
// 					if (i == 0 && j == 0) || (i == 0 && j == board.BOARD_SIZE-1) || (i == board.BOARD_SIZE-1 && j == 0) || (i == board.BOARD_SIZE-1 && j == board.BOARD_SIZE-1) {
// 						score -= 2
// 					} else {
// 						score -= 1
// 					}
// 				}
// 				if (i >= board.BOARD_SIZE/2-1 && i <= board.BOARD_SIZE/2+1) && (j >= board.BOARD_SIZE/2-1 && j <= board.BOARD_SIZE/2+1) {
// 					centerControlCount++
// 					score += 3
// 				}
// 			}
// 		}
// 	}

// 	// Check for adjacent circle and triangle
// 	for _, circlePos := range circlePositions {
// 		for _, trianglePos := range trianglePositions {
// 			if manhattanDistance(circlePos, trianglePos) == 1 {
// 				// Calculate total Manhattan distance for each type
// 				circleTotalDistance := calculateTotalDistance(circlePos, circlePositions)
// 				triangleTotalDistance := calculateTotalDistance(trianglePos, trianglePositions)

// 				// Decide based on distances
// 				if circleTotalDistance < triangleTotalDistance {
// 					if player == board.CIRCLE {
// 						score += 3 // Favor circles
// 					} else {
// 						score -= 3 // Disfavor triangles
// 					}
// 				}
// 			}
// 		}
// 	}

// 	// Reward calculations for near-border and center
// 	score += nearBorderCount * 2
// 	score += centerControlCount * 3

// 	// Finally, return the score based on the number of pieces the player has
// 	if player == board.TRIANGLE {
// 		return score + MockTriangleNum*5
// 	}
// 	return score + MockCircleNum*5
// }
