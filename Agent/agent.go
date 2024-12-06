package agent

import (
	"main/board"
	"main/game"

	"github.com/gdamore/tcell/v2"
)

func AgentAction(screen tcell.Screen) {
	if game.ValidMoveCheck(4, 0, 4, 1) {
		board.MovePiece(4, 0, 4, 1, screen)
	}
}
