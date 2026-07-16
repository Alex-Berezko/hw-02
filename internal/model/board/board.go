package board

//Поле должно хранить как минимум:

//Что хранит:
//- размер поля
//- расположение фигур на поле
//
//Что умеет делать:

//- получить содержимое клетки
//- поставить фигуру в клетку
//- переместить фигуру

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

func ShowPieceOnBoard() {

}

// - получить содержимое клетки
func MapBoard() {

}
