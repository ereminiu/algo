package quicksortv2

import "sync"

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

func quicksort(l, r int, a *[]int) {
	if l >= r {
		return
	}
	if r-l+1 < limit {
		f(l, r, a)
		return
	}

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
