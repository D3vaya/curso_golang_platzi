package main

import "fmt"

func main() {
	//PrintLn
	helloMessage := "Hello"
	worldMessage := "world"

	fmt.Println(helloMessage, worldMessage)

	//Printf
	nombre := "platzi"
	cursos := 500
	fmt.Printf("%s tiene mas de %d\n", nombre, cursos)
	fmt.Printf("%v tiene mas de %v\n", nombre, cursos)

	//Sprintf
	message := fmt.Sprintf("%s tiene mas de %d", nombre, cursos)
	fmt.Sprintln(message)

	// Tipo de datos
	fmt.Printf("helloMessage; %T \n", helloMessage)
	fmt.Printf("cursos; %T \n", cursos)
}
