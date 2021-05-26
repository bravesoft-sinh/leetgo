package main

import (
	"fmt"
	"math/rand"
	"time"
)

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
	fmt.Println()
	fmt.Println("Insertion  Sort")
	printArray(insertionSort([]int{3, 2, 4, 5, 1}))

	fmt.Println()
	fmt.Println("Quick Sort")
	var a = []int{2, 3, 1, 5, 4}
	quickSort(a, 0, len(a)-1)
	printArray(a)
	var millionArray []int = randomArray(1000000)
	start := time.Now()
	// Code to measure
	quickSort(millionArray, 0, len(millionArray)-1)
	fmt.Println(len(millionArray))
	duration := time.Since(start)
	fmt.Println()
	fmt.Println("Execution time: ", duration)

}
func sort(arr []int) []int {
	var n int = len(arr)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i] > arr[j] {
				swap(arr, i, j)
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
				swap(arr, j, j+1)
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

func insertionSort(arr []int) []int {
	var n = len(arr)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if arr[j-1] > arr[j] {
				swap(arr, j-1, j)
			}
			j = j - 1
		}
	}
	return arr
}

func quickSort(arr []int, begin, end int) {
	if begin < end {
		partitionIndex := partition(arr, begin, end)
		quickSort(arr, begin, partitionIndex-1)
		quickSort(arr, partitionIndex+1, end)
	}
}

func partition(arr []int, begin, end int) int {
	pivot := arr[end]
	i := begin - 1

	for j := begin; j < end; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	swap(arr, i+1, end)
	//arr[i+1], arr[end] = arr[end], arr[i+1]

	return i + 1
}
func printArray(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Print(" ")
	}
}

func swap(arr []int, i int, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func randomArray(n int) []int {
	rand.Seed(time.Now().Unix())
	return rand.Perm(n)
}
