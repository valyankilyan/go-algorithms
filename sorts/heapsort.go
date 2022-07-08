package sorts

import "github.com/valyankilyan/go-algorithms/constraints"

func heapify[T constraints.Ordered](arr []T, cur, n int) {
	largest := cur
	lc, rc := cur*2+1, cur*2+2
	if lc < n && arr[lc] > arr[largest] {
		largest = lc
	}
	if rc < n && arr[rc] > arr[largest] {
		largest = rc
	}

	if largest != cur {
		arr[cur], arr[largest] = arr[largest], arr[cur]
		heapify(arr, largest, n)
	}

}

func HeapSort[T constraints.Ordered](arr []T) {
	if len(arr) == 0 {
		return
	}
	for i := len(arr) / 2; i >= 0; i-- {
		heapify(arr, i, len(arr))
	}

	for i := len(arr) - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, 0, i)
	}
}
