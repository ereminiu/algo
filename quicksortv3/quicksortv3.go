package quicksortv3

import (
	"sync"
)

var limit = 1 << 11

func Sort(a *[]int) {
	quicksort(0, len(*a)-1, a)
}

func f(l, r int, a *[]int) {
	if l >= r {
		return
	}
	q := partition(l, r, a)
	f(l, q, a)
	f(q+1, r, a)
}

func insertion(l, r int, a *[]int) {
	n := r - l + 1
	for i := 0; i < n; i++ {
		idx := i
		for j := i + 1; j < n; j++ {
			if (*a)[j] < (*a)[idx] {
				idx = j
			}
		}
		(*a)[i], (*a)[idx] = (*a)[idx], (*a)[i]
	}
}

func quicksort(l, r int, a *[]int) {
	if l >= r {
		return
	}
	if r-l+1 < 11 {
		insertion(l, r, a)
		return
	}
	if r-l+1 < limit {
		f(l, r, a)
		return
	}

	med := median(l, (l+r)/2, r, a)
	(*a)[med], (*a)[(l+r)/2] = (*a)[(l+r)/2], (*a)[med]

	q := partition(l, r, a)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		quicksort(l, q, a)
	}()

	wg.Wait()
	quicksort(q+1, r, a)
}

func partition(l, r int, a *[]int) int {
	v := (*a)[(l+r)/2]
	i, j := l, r
	for i <= j {
		for (*a)[i] < v {
			i++
		}
		for (*a)[j] > v {
			j--
		}
		if i >= j {
			break
		}
		(*a)[i], (*a)[j] = (*a)[j], (*a)[i]
		i++
		j--
	}
	return j
}

func boolToInt(f bool) int {
	if f {
		return 1
	}
	return 0
}

func median(x, y, z int, a *[]int) int {
	if boolToInt((*a)[x] > (*a)[y])^boolToInt((*a)[x] > (*a)[z]) == 1 {
		return x
	} else if boolToInt((*a)[y] < (*a)[x])^boolToInt((*a)[y] < (*a)[z]) == 1 {
		return y
	}
	return z
}
