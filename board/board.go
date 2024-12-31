package board

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

const (
	BOARD_SIZE = 7
)

var CircleNum = 4
var TriangleNum = 4
var MoveCounter = 0

var GameStatus = 0

var X, Y = 0, 0

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

func RenderBoard(screen tcell.Screen, currentPlayer Element) {
	screen.Clear()
	cellWidth := 2

	for i := 0; i < BOARD_SIZE; i++ {
		for j := 0; j < BOARD_SIZE; j++ {
			ch := element[Board[i][j]]
			style := tcell.StyleDefault

			if (i+j)%2 == 0 {
				style = style.Background(tcell.ColorDarkGreen)
			} else {
				style = style.Background(tcell.ColorDarkBlue)
			}

			if i == X && j == Y {
				style = style.Background(tcell.ColorGray)
			}

			if Board[i][j] == currentPlayer {
				style = style.Foreground(tcell.ColorYellow)
			}

			if Board[i][j] == CIRCLE || Board[i][j] == TRIANGLE {
				style = style.Foreground(tcell.ColorYellow)
			}

			runes := []rune(ch)
			screen.SetContent(j*cellWidth, i, runes[0], nil, style.Foreground(tcell.ColorYellow))
			for k := 1; k < cellWidth; k++ {
				screen.SetContent(j*cellWidth+k, i, ' ', nil, style)
			}
		}
	}

	info := "Turn: " + element[currentPlayer]
	for i, r := range info {
		screen.SetContent(i, BOARD_SIZE+1, r, nil, tcell.StyleDefault)
	}

	info = "Move Counter: " + strconv.Itoa(MoveCounter)
	for i, r := range info {
		screen.SetContent(i, BOARD_SIZE+2, r, nil, tcell.StyleDefault)
	}

	info = "Press ESC for exit."
	for i, r := range info {
		screen.SetContent(i, BOARD_SIZE+3, r, nil, tcell.StyleDefault)
	}

	screen.Show()
}

func MovePiece(fromX, fromY, X, Y int, screen tcell.Screen, CurrentPlayer Element) bool {
	Board[X][Y] = Board[fromX][fromY]
	Board[fromX][fromY] = EMPTY

	RenderBoard(screen, CurrentPlayer)
	return true
}

func RemovePiece(deathValues [][2]int, screen tcell.Screen, CurrentPlayer Element) {
	for _, v := range deathValues {

		if v == [2]int{-1, -1} {
			continue
		}

		if Board[v[0]][v[1]] == TRIANGLE {
			TriangleNum--
		} else {
			CircleNum--
		}

		Board[v[0]][v[1]] = EMPTY
	}
	RenderBoard(screen, CurrentPlayer)
}

func GameOverMessage(screen tcell.Screen) {
	screen.Clear()
	gameOverMsg := "Game Over"
	for i, ch := range gameOverMsg {
		screen.SetContent(i, 0, ch, nil, tcell.StyleDefault)
	}

	if GameStatus == -1 {
		gameOverMsg = "It's a Draw"
	} else if GameStatus == 0 {
		gameOverMsg = "Triangle Wins"
	} else if GameStatus == 1 {
		gameOverMsg = "Circle Wins"
	}

	for i, ch := range gameOverMsg {
		screen.SetContent(i, 1, ch, nil, tcell.StyleDefault)
	}

	gameOverMsg = "Press ESC to exit"

	for i, ch := range gameOverMsg {
		screen.SetContent(i, 2, ch, nil, tcell.StyleDefault)
	}

	screen.Show()

	for {
		ev := screen.PollEvent()
		switch ev := ev.(type) {
		case *tcell.EventKey:
			switch ev.Key() {
			case tcell.KeyEscape:
				return
			}
		}
	}

}
