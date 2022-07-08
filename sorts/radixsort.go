package sorts

import (
	"github.com/valyankilyan/go-algorithms/constraints"
)

func RadixSort[T constraints.Natural](arr []T) {
	var mx T
	for _, m := range arr {
		mx = max(m, mx)
	}

	var pow uint64 = 1
	for T(pow) < mx {
		pow *= 10
	}
	pow /= 10
	countSort(arr, pow)
}

func RadixSortString(arr []string) {
	byteArr := make([][]byte, len(arr))
	for i := range arr {
		byteArr[i] = []byte(arr[i])
	}

	countSortBytes(byteArr, 0)

	for i := range byteArr {
		arr[i] = string(byteArr[i])
	}
}

func countSort[T constraints.Natural](arr []T, pow uint64) {
	if pow < 1 {
		return
	}
	count := make([]int, 10)
	cop := make([]T, len(arr))
	copy(cop, arr)

	for _, a := range cop {
		count[(uint64(a)/pow)%10]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
		// defer countSort(arr[count[i-1]:count[i]], pow/10)
	}

	defer countSort(arr[:count[0]], pow/10)
	for i := 0; i < len(count)-1; i++ {
		defer countSort(arr[count[i]:count[i+1]], pow/10)
	}
	defer countSort(arr[count[9]:], pow/10)

	for i := len(cop) - 1; i >= 0; i-- {
		arr[count[(uint64(cop[i])/pow)%10]-1] = cop[i]
		count[(uint64(cop[i])/pow)%10]--
	}
}

func countSortBytes(arr [][]byte, ln int) {
	count := make([]int, 257)
	for _, c := range arr {
		if len(c) <= ln {
			count[0]++
		} else {
			count[c[ln]+1]++
		}
	}
	if count[0] == len(arr) {
		return
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}

	for i := 0; i < len(count)-1; i++ {
		defer countSortBytes(arr[count[i]:count[i+1]], ln+1)
	}
	defer countSortBytes(arr[count[len(count)-1]:], ln+1)

	cp := make([][]byte, len(arr))
	copy(cp, arr)

	for i := len(cp) - 1; i >= 0; i-- {
		var cur int
		if len(cp[i]) <= ln {
			cur = 0
		} else {
			cur = int(cp[i][ln] + 1)
		}
		arr[count[cur]-1] = cp[i]
		count[cur]--
	}
}
