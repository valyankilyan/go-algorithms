package sorts

import (
	"github.com/valyankilyan/go-algorithms/constraints"
)

func InsertionSort[T constraints.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		cur, place := arr[i], i
		for place > 0 && cur < arr[place-1] {
			place--
		}
		for j := i - 1; j >= place; j-- {
			arr[j+1] = arr[j]
		}
		arr[place] = cur
	}
}
