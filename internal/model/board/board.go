package board

type Board struct {
	size     int
	mapPiece string
}

func (b *Board) AddSizeBoard(size int) {
	b.size = size
}

// - сообщать размер
func (b Board) SaySize() int {
	return b.size
}
