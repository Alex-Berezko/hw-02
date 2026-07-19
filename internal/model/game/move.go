package game

import (
	"fmt"
)

type Move struct {
	startPosition string
	endPosition   string
	piece         rune
}

func (m *Move) SetStartPosition(start string) {
	m.startPosition = start
}
func (m *Move) SetEndPosition(end string) {
	m.endPosition = end
}
func (m *Move) SetPiece(piece rune) {
	m.piece = piece
}

func (m Move) SayMove() string {
	return fmt.Sprintf(" %c %s → %s", m.piece, m.startPosition, m.endPosition)

}

func (m Move) SayEndPosition() {
	fmt.Printf("Конечная позиция %s", m.endPosition)
}

func (m Move) SayPieceMove() {
	fmt.Printf("Ходила следующая фигура %v", m.piece)
}
func (m Move) SayStartPosition() {
	fmt.Printf("Начальная позиция %s", m.startPosition)
}
