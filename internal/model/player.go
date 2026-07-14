package model

import (
	"bufio"
	"fmt"
	"os"
)

type User struct {
	name       string
	pieceColor string
	/*pieceColor map[string]string(
	"White": "",
	"Black": "",
	)*/
}

func SayName(name User) {
	fmt.Println(name.name)
}

func Rename(name *User) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите новок имя игрока")
	scanner.Scan()
	newName := scanner.Text()
	name.name = newName
}

func SayPieceColor(color User) {
	fmt.Println(color.pieceColor)
}
