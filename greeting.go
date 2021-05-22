package main

import "fmt"

func main() {
	fmt.Println("1 + 2 = ", add(1, 2))
}

func add(a int, b int) int {
	return a + b
}


