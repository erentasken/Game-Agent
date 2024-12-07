package board

import (
	"os"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

const (
	BOARD_SIZE = 7
)

var CircleNum = 4
var TriangleNum = 4
var RoundCounter = 0

type Element int

const (
	EMPTY Element = iota
	TRIANGLE
	CIRCLE
)

var element = map[Element]string{
	EMPTY:    "  ",
	TRIANGLE: "\u25B3", // Triangle
	CIRCLE:   "\u25CB", // Circle
}

var Board [BOARD_SIZE][BOARD_SIZE]Element

// Initialize the board with predefined positions
func CreateBoard() {
	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			if (i == 0 && j == 0) || (i == 2 && j == 0) || (i == 4 && j == 6) || (i == 6 && j == 6) {
				Board[i][j] = TRIANGLE
			} else if (i == 0 && j == 6) || (i == 2 && j == 6) || (i == 4 && j == 0) || (i == 6 && j == 0) {
				Board[i][j] = CIRCLE
			} else {
				Board[i][j] = EMPTY
			}
		}
	}
}

func RenderBoard(screen tcell.Screen, X, Y int, currentPlayer Element) {
	screen.Clear()
	cellWidth := 2 // Fixed width for each cell to maintain alignment

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			ch := element[Board[i][j]]
			style := tcell.StyleDefault

			// Alternate background color for the checkerboard pattern
			if (i+j)%2 == 0 {
				style = style.Background(tcell.ColorDarkGreen)
			} else {
				style = style.Background(tcell.ColorDarkBlue)
			}

			// Highlight the cursor's square
			if i == X && j == Y {
				style = style.Background(tcell.ColorGray)
			}

			// if Board[i][j] == currentPlayer {
			// 	style = style.Foreground(tcell.ColorYellow)
			// }

			if Board[i][j] == CIRCLE || Board[i][j] == TRIANGLE {
				style = style.Foreground(tcell.ColorYellow)
			}

			// Set cell content with proper padding
			runes := []rune(ch)
			screen.SetContent(j*cellWidth, i, runes[0], nil, style.Foreground(tcell.ColorYellow))
			for k := 1; k < cellWidth; k++ {
				screen.SetContent(j*cellWidth+k, i, ' ', nil, style)
			}
		}
	}

	// Display the current player's turn
	info := "Turn: " + element[currentPlayer]
	for i, r := range info {
		screen.SetContent(i, BOARD_SIZE+1, r, nil, tcell.StyleDefault)
	}

	// Display the move counter
	info = "Move Counter: " + strconv.Itoa(RoundCounter)
	for i, r := range info {
		screen.SetContent(i, BOARD_SIZE+2, r, nil, tcell.StyleDefault)
	}

	screen.Show()
}

// Move a piece based on cursor position and input
func MovePiece(fromX, fromY, X, Y int, screen tcell.Screen) bool {
	Board[X][Y] = Board[fromX][fromY]
	Board[fromX][fromY] = EMPTY

	return true
}

func RemovePiece(deathValues [][2]int) {
	for _, v := range deathValues {
		if Board[v[0]][v[1]] == TRIANGLE {
			TriangleNum--
		} else {
			CircleNum--
		}

		Board[v[0]][v[1]] = EMPTY
	}
}

func EndGameDisplay(info string, screen tcell.Screen) {
	switch info {
	case "Triangle":
		info := "TRIANGLE WINS!, GAME OVER"
		for i, r := range info {
			screen.SetContent(i, BOARD_SIZE+3, r, nil, tcell.StyleDefault)
		}
		screen.Show()

		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return // Exit on ESC
				}
			}
		}
	case "Circle":
		info := "CIRCLE WINS!, GAME OVER"
		for i, r := range info {
			screen.SetContent(i, BOARD_SIZE+3, r, nil, tcell.StyleDefault)
		}
		screen.Show()

		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return // Exit on ESC
				}
			}
		}
	case "Draw":
		info := "It's a Draw!, GAME OVER"
		for i, r := range info {
			screen.SetContent(i, BOARD_SIZE+3, r, nil, tcell.StyleDefault)
		}
		screen.Show()

		os.Exit(0)

		for {
			ev := screen.PollEvent()
			switch ev := ev.(type) {
			case *tcell.EventKey:
				switch ev.Key() {
				case tcell.KeyEscape:
					return // Exit on ESC
				}
			}
		}
	}
}
