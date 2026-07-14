package model

import "fmt"

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

// - сообщать размер
func SaySize(board Board) {
	fmt.Printf("Размер поля %v на %v", board.size, board.size)
}

func ShowPieceOnBoard() {

}

// - получить содержимое клетки
func MapBoard() {

}
