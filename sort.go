package main

import "fmt"

func main() {
	var arr = []int{2, 3, 1, 5, 4}
	fmt.Println()
	fmt.Println("Bubble Sort")
	arr = bubbleSort(arr)
	printArray(arr)
	fmt.Println()
	fmt.Println("Big O Sort")
	arr = sort([]int{2, 3, 1, 5, 4})
	printArray(arr)
	fmt.Println()
	fmt.Println("Selection Sort")
	printArray(selectionSort([]int{3, 2, 4, 5, 1}))
}
func sort(arr []int) []int {
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
func bubbleSort(arr []int) []int {
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

func selectionSort(arr []int) []int {
	var n = len(arr)
	for i := 0; i < n; i++ {
		var minIndex = i
		for j := i + 1; j < n; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		swap(arr, i, minIndex)
	}
	return arr
}

func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}

func swap(arr []int, i int, j int) {
	var temp = arr[i]
	arr[i] = arr[j]
	arr[j] = temp
}
