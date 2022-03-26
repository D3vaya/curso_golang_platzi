package main

import "fmt"

func main() {

	// DEFER
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// CONTINUE & BREAK

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("es 2")
			continue
		}

		if i == 8 {
			fmt.Println("Break")
			break
		}

	}
}
