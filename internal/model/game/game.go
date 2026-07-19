package game

import (
	"fmt"
	"hw-02/internal/model/board"
	"hw-02/internal/model/player"
)

type Game struct {
	whitePlayer player.User
	blackPlayer player.User
	board       board.Board
	move        Move
	moveHistory []Move
	status      string
	whoseMove   string
}

func (g *Game) SetUserWhite(name string) {
	g.whitePlayer.SetName(name)
	g.whitePlayer.SetColor("Белые")

}

func (g Game) SayNameUserWhite() string {
	return g.whitePlayer.SayName()
}

func (g *Game) SetUserBlack(name string) {
	g.blackPlayer.SetName(name)
	g.blackPlayer.SetColor("Черные")
}

func (g Game) SayNameUserBlack() string {
	return g.blackPlayer.SayName()
}

func (g Game) InfoOfTheGamePlayer() {
	fmt.Printf("Размер доски %v на %v \n", g.board.SaySize(), g.board.SaySize())
	fmt.Printf("Белые: %s \n", g.SayNameUserWhite())
	fmt.Printf("Цвет фигур: %s \n", g.whitePlayer.SayPieceColor())

	fmt.Printf("Чёрные: %s \n", g.SayNameUserBlack())
	fmt.Printf("Цвет фигур: %s \n", g.blackPlayer.SayPieceColor())
}

func (g *Game) SetSizeBoard(size int) {
	g.board.AddSizeBoard(size)
}
func (g Game) SaySizeBoard() int {
	return g.board.SaySize()
}

func (g *Game) Move(piece rune, start, end string) {
	g.move.SetStartPosition(start)
	g.move.SetEndPosition(end)
	g.move.SetPiece(piece)

	g.saveMoveToHistori()
}

func (g *Game) saveMoveToHistori() {
	g.moveHistory = append(g.moveHistory, g.move)
}
func (g Game) SayMoveHistori() {
	fmt.Println("История ходов:")

	for i, move := range g.moveHistory {
		fmt.Printf(" %d %s\n", i+1, move.SayMove())
	}

}
