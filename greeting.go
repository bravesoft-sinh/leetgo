package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("1 + 2 = ", add(1, 2))
	go say("world")
	say("hello")
}

func add(a int, b int) int {
	return a + b
}
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}


