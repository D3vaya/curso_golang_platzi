package main

import (
	pk "curso_golang_platzi/src/mypackage"
	"fmt"
)

func main() {
	a := 50
	b := &a
	fmt.Println(a)
	fmt.Println(*b)

	*b = 100
	fmt.Println(a)

	myPc := pk.Pc{Ram: 16, Disk: 200, Brand: "msi"}
	fmt.Println(myPc)
	myPc.Ping()

	fmt.Println(myPc)

	fmt.Println(myPc)
	myPc.DuplicateRAM()

	fmt.Println(myPc)
}
