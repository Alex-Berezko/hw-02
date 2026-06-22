package main

import "fmt"

func main() {
	var a int
	fmt.Println("Введите количество ячеек а")
	fmt.Scan(&a)

	for i := 0; i < a; i++ {
		if i%2 == 0 {
			for j := 0; j < a; j++ {
				if j%2 == 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}

			}
			fmt.Print("\n")
		} else {
			for j := 0; j < a; j++ {

				if j%2 != 0 {
					fmt.Print("#")
				} else {
					fmt.Print(" ")
				}

			}
			fmt.Print("\n")
		}
	}
}
