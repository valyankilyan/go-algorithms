package sorts

import (
	"math"

	"github.com/valyankilyan/go-algorithms/constraints"
)

func max[T constraints.Real](a, b T) T {
	if a > b {
		return a
	} else {
		return b
	}
}

func CountingSort[T constraints.Real](arr []T) {
	var mx T
	for _, a := range arr {
		mx = max(a, mx)
	}
	// Мы не хотим быть дольше n logn
	if mx > T(len(arr))*T(math.Log2(float64(len(arr)))) {
		MergeSort(arr)
		return
	}

	counting_arr := make([]int, mx+1)
	for _, a := range arr {
		counting_arr[a]++
	}

	cur := 0
	for i, c := range counting_arr {
		if c != 0 {
			for j := 0; j < c; j++ {
				arr[cur] = T(i)
				cur++
			}
		}
	}
}
