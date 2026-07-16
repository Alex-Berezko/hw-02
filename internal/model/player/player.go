package player

import (
	"fmt"
)

type User struct {
	name       string
	pieceColor string
}

func (u *User) AddName(name string) {
	u.name = name
}

func (u *User) AddColor(color string) {
	u.pieceColor = color
}

func (u User) SayName() string {
	return u.name

}

func Rename(name *User) {

}

func SayPieceColor(color User) {
	fmt.Println(color.pieceColor)
}
