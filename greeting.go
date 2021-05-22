package main

import "fmt"

func main() {
	fmt.Println("1 + 2 = ", add(1, 2))
	var arr = [5]int{2, 3, 1, 5, 4}
	sort(arr)
}

func add(a int, b int) int {
	return a + b
}

func sort(arr [5]int) {
	var n int = len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n-1; j++ {
			if arr[i] > arr[j] {
				var temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Print(arr[i])
	}
}
