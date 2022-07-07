package quicksort

import (
	"github.com/valyankilyan/go-algorithms/constraints"
)

func genPivot(left, right int) int {
	return right
}

func QuickSort[T constraints.Ordered](arr []T) {
	quickSort(arr, 0, len(arr)-1)
}

func quickSort[T constraints.Ordered](arr []T, left, right int) {
	if left >= right {
		return
	}
	pl := partition(arr, left, right)
	quickSort(arr, left, pl-1)
	quickSort(arr, pl+1, right)
}

func partition[T constraints.Ordered](arr []T, left, right int) int {
	piv := genPivot(left, right)
	pivot := arr[piv]
	arr[piv], arr[right] = arr[right], arr[piv]
	pl := left
	for i := left; i < right; i++ {
		if arr[i] < pivot {
			arr[pl], arr[i] = arr[i], arr[pl]
			pl++
		}
	}
	arr[right], arr[pl] = arr[pl], arr[right]
	return pl
}
