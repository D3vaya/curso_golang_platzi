package main

import "fmt"

func say(text string, c chan<- string) {
	c <- text
}

func main() {
	c := make(chan string)

	fmt.Println("Hello")

	go say("Bye", c)
	fmt.Println(<-c)
}
