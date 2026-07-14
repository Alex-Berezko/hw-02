package model

import (
	"fmt"
)

type game struct {
	User
	Board
	Move
	histori   []string
	status    string
	whoseMove string
}

func MoveGame(game2 game) {
	fmt.Println("откуда ходим")
	fmt.Scan(&game2.startPosition)
	fmt.Println("куда ходим")
	fmt.Scan(&game2.endPosition)
	newMove := game2.startPosition + game2.endPosition
	game2.histori = append(newMove)
}
