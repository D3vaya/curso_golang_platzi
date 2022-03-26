package main

import "fmt"

func main() {
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado: ", areaCuadrado)

	x := 10
	y := 50

	// SUMA
	result := x + y
	fmt.Println("Suma:", result)

	// RESTA
	result = y - x
	fmt.Println("Resta:", result)

	// MULTIPLICACION
	result = y * x
	fmt.Println("Multiplicacion:", result)

	// DIVISION
	result = y / x
	fmt.Println("Division:", result)

	// MODULO
	result = y % x
	fmt.Println("Cuosiente:", result)

	// INCREMENTAL
	x++
	fmt.Println("Incrementeal:", x)

	// DECREMENTAL
	x--
	fmt.Println("Incrementeal:", x)
}
