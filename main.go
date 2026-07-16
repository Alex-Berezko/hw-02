package main

import (
	"bufio"
	"fmt"
	"hw-02/internal/model/board"
	"hw-02/internal/model/player"
	"os"
)

var whitePlayer = player.User{}
var blackPlayer = player.User{}
var playerBoard = board.Board{}

func main() {
	var a int

	fmt.Println("Введите количество ячеек")
	fmt.Scanln(&a)
	playerBoard.AddSizeBoard(a)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите имя игрока за белых")
	scanner.Scan()
	userWhiteName := scanner.Text()

	whitePlayer.AddName(userWhiteName)
	whitePlayer.AddColor("Белые")

	fmt.Println("Введите имя игрока за черных")
	scanner.Scan()
	userBlackName := scanner.Text()

	blackPlayer.AddName(userBlackName)
	blackPlayer.AddColor("Черные")

	printLetters(playerBoard.SaySize())

	printLine(playerBoard.SaySize())

	printBoard(playerBoard.SaySize())

	printLine(playerBoard.SaySize())

	printLetters(playerBoard.SaySize())

	fmt.Println()

}

func printLetters(a int) {
	abc := []rune{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'}
	fmt.Print("    ")
	for i := 0; i < a; i++ {
		switch {
		case i < len(abc):

			fmt.Printf(" %c ", abc[i])

		case i >= len(abc):

			number := i - len(abc)

			firstLetter := abc[number/len(abc)]
			secondLetter := abc[number%len(abc)]

			fmt.Printf(" %c%c", firstLetter, secondLetter)

		}
	}

	fmt.Println()
}

func printLine(a int) {
	for i := 0; i < a*3; i++ {
		if i == 0 {
			fmt.Print("   ")
		}
		fmt.Print("—")
	}
	fmt.Println()

}

func printCell(i, j int) {
	if (j+i)%2 == 0 {
		fmt.Print(" # ")
	} else {
		fmt.Print("   ")
	}

}

func printBoard(a int) {
	blackChess := []rune{'♜', '♞', '♝', '♛', '♚', '♝', '♞', '♜'}
	whiteChess := []rune{'♖', '♘', '♗', '♕', '♔', '♗', '♘', '♖'}
	blackPawn := '♟'
	whitePawn := '♙'

	for i := 0; i < a; i++ {

		fmt.Printf("%v |", a-i)

		for j := 0; j < a; j++ {
			switch {
			case i == 0:
				fmt.Printf(" %c ", blackChess[j%len(blackChess)])

				if j == a-1 {
					fmt.Printf("| %s \n", blackPlayer.SayName())
				}
			case a >= 4 && i == 1:
				fmt.Printf(" %c ", blackPawn)

			case a >= 4 && i == a-2:
				fmt.Printf(" %c ", whitePawn)

			case i == a-1:
				fmt.Printf(" %c ", whiteChess[j%len(whiteChess)])

				if j == a-1 {
					fmt.Printf("| %s \n", whitePlayer.SayName())
				}
			default:
				printCell(i, j)
			}
		}
		if i != 0 && i != a-1 {
			fmt.Printf("| \n")
		}
	}
}
