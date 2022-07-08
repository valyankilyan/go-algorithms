package main

import (
	"fmt"

	"github.com/valyankilyan/go-algorithms/sorts"
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

	testN := [][]uint{
		{5, 4, 3, 1, 2},
		{1, 24455, 13, 133, 53, 23},
		{10, 100, 1000, 10000},
	}

	testStr := [][]string{
		{"a", "c", "b", "n"},
		{"", "a", "aa", "aaa", "aaaa", "aaaaa"},
		{"aaaaa", "aaaa", "aaa", "aa", "a", ""},
		{"Apple", "juice", "is", "what", "I", "like"},
		{"go", "Go", "gopher", "Gopher"},
		{"Кот", "Cat", "Собака", "Dog"},
	}

	for _, arr := range test {
		// mergesort.MergeSort(arr)
		// quicksort.QuickSort(arr)
		// sorts.BubbleSort(arr)
		// sorts.SelectionSort(arr)
		// sorts.InsertionSort(arr)
		// sorts.HeapSort(arr)
		sorts.CountingSort(arr)
	}
	fmt.Println(test)

	for _, arr := range testN {
		sorts.RadixSort(arr)
	}
	fmt.Println(testN)

	for _, s := range testStr {
		// sort.Strings(s)
		sorts.RadixSortString(s)
	}
	fmt.Println(testStr)
}
