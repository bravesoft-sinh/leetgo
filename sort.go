package main

import "fmt"

func main() {
	var arr = [5]int{2, 3, 1, 5, 4}
	arr = bubleSort(arr)
	printArray(arr)
	arr = sort([5]int{2, 3, 1, 5, 4})
	printArray(arr)
}
func sort(arr [5]int) [5]int {
	var n int = len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				var temp = arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
func bubleSort(arr [5]int) [5]int {
	var n = len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				var temp = arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	return arr
}
func printArray(arr [5]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}
