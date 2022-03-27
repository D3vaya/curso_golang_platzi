package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["jose"] = 14
	m["pepito"] = 20

	fmt.Println(m)

	// Recorrer un map

	for i, v := range m {
		fmt.Println(i, v)
	}

	// encontrar un valor en un map
	value, ok := m["jose"]
	fmt.Println(value, ok)
}
