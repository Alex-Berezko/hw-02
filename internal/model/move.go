package model

import (
	"fmt"
)

type Move struct {
	startPosition string
	endPosition   string
	piece         rune
	User
	numMove int
}

func SayStartPosition(move2 Move) {
	fmt.Printf("Начальная позиция %s", move2.startPosition)
}

func SayEndPositio(move2 Move) {
	fmt.Printf("Конечная позиция %s", move2.endPosition)
}

func SayPieceMove(move2 Move) {
	fmt.Printf("Ходила следующая фигура %v", move2.piece)
}

func MovePiece(move Move) {
	fmt.Printf("Игрок %s ходит из %s на %s", move.name, move.startPosition, move.endPosition)
}
