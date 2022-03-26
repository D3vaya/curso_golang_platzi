package main

import "fmt"

func normalFunction(saludo string) {
	fmt.Println(saludo)
}

func muchosArgumentos(a, b int, c string) {
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int) {
	return a, a * 2
}

func main() {
	normalFunction("Hola Mundo !")
	muchosArgumentos(1, 10, "hola")
	value := returnValue(2)
	fmt.Println("Value: ", value)
	value1, value2 := dobleReturn(2)
	fmt.Println("Value1 y Value2", value1, value2)
	valueOne, _ := dobleReturn(2)
	fmt.Println("valueOne ", valueOne)

}
