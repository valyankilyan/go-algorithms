package mergesort

import "github.com/valyankilyan/go-algorithms/constraints"

func MergeSort[T constraints.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}

	MergeSort(a[len(a)/2:])
	MergeSort(a[:len(a)/2])

	copy := make([]T, len(a))
	for i := range a {
		copy[i] = a[i]
	}
	l, r := 0, len(copy)/2

	for i := range a {
		if l == len(copy)/2 {
			a[i] = copy[r]
			r++
			continue
		}
		if r == len(copy) {
			a[i] = copy[l]
			l++
			continue
		}

		if copy[l] < copy[r] {
			a[i] = copy[l]
			l++
		} else {
			a[i] = copy[r]
			r++
		}
	}
}
