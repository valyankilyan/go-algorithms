package sorts

import "github.com/valyankilyan/go-algorithms/constraints"

func SelectionSort[T constraints.Ordered](arr []T) {
	for i := range arr {
		min := i
		for j := i; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
