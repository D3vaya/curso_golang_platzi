package main

import (
	"fmt"
)

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPc pc) String() string {
	return fmt.Sprintf("Tengo %d GB de RAM, %d GB de Disco y es una %s", myPc.ram, myPc.disk, myPc.brand)
}

func main() {
	myPc := pc{ram: 16, brand: "ms1", disk: 100}
	fmt.Println(myPc)
}
