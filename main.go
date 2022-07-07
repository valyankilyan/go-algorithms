package main

import (
	"fmt"

	"github.com/valyankilyan/go-algorithms/sort/bubblesort"
)

func main() {
	test := [][]int{
		{6, 5, 4, 3, 2, 1},
		{5, 4, 3, 2, 1},
		{2, 5, 2, 5},
		{},
		{1},
		{36, 3342, 2534, 257, 32357, 23, 7, 2, 56, 3, 6},
	}

	for _, arr := range test {
		// mergesort.MergeSort(arr)
		// quicksort.QuickSort(arr)
		bubblesort.BubbleSort(arr)
	}
	fmt.Println(test)
}
