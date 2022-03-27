package mypackage

import "fmt"

type Pc struct {
	Ram   int
	Disk  int
	Brand string
}

// Ping retorna pong
func (myPc Pc) Ping() {
	fmt.Println(myPc.Brand, "Pong")
}

// DuplicateRAM retorna pong
func (myPc *Pc) DuplicateRAM() {
	myPc.Ram = myPc.Ram * 2
}
