package sorts

import "github.com/valyankilyan/go-algorithms/constraints"

func MergeSort[T constraints.Ordered](a []T) {
	if len(a) <= 1 {
		return
	}

	MergeSort(a[len(a)/2:])
	MergeSort(a[:len(a)/2])

	c := make([]T, len(a))
	copy(c, a)

	l, r := 0, len(c)/2

	for i := range a {
		if l == len(c)/2 {
			a[i] = c[r]
			r++
			continue
		}
		if r == len(c) {
			a[i] = c[l]
			l++
			continue
		}

		if c[l] < c[r] {
			a[i] = c[l]
			l++
		} else {
			a[i] = c[r]
			r++
		}
	}
}
