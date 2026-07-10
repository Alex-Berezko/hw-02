package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var a int

	fmt.Println("Введите количество ячеек")
	fmt.Scanln(&a)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите имя игрока за белых")
	scanner.Scan()
	userWhite := scanner.Text()

	fmt.Println("Введите имя игрока за черных")
	scanner.Scan()
	userBlack := scanner.Text()

	printLetters(a)

	printLine(a)

	blackChess := []rune{'♜', '♞', '♝', '♛', '♚', '♝', '♞', '♜'}
	whiteChess := []rune{'♖', '♘', '♗', '♕', '♔', '♗', '♘', '♖'}
	blackPawn := '♟'
	whitePawn := '♙'

	for i := 0; i < a; i++ {

		fmt.Printf("%v |", a-i)

		for j := 0; j < a; j++ {
			switch {
			case i == 0:
				if j < 8 {
					fmt.Printf(" %c ", blackChess[j])
				} else {
					printCell(i, j)
				}

				if j == a-1 {
					fmt.Printf("| %s \n", userBlack)
				}
			case a >= 4 && i == 1:
				if j < 8 {
					fmt.Printf(" %c ", blackPawn)
				} else {
					printCell(i, j)
				}
			case a >= 4 && i == a-2:
				if j < 8 {
					fmt.Printf(" %c ", whitePawn)
				} else {
					printCell(i, j)
				}

			case i == a-1:
				if j < 8 {
					fmt.Printf(" %c ", whiteChess[j])
				} else {
					printCell(i, j)
				}
				if j == a-1 {
					fmt.Printf("| %s \n", userWhite)
				}
			default:
				printCell(i, j)
			}
		}
		if i != 0 && i != a-1 {
			fmt.Printf("| \n")
		}
	}

	printLine(a)

	printLetters(a)

	fmt.Println()
}

func printLetters(a int) {
	for i := 0; i < a; i++ {
		if i == 0 {
			fmt.Print("    ")
		}
		fmt.Printf("%c  ", rune('A'+i))
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
