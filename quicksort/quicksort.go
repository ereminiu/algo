package quicksort

func Sort(a []int) {
	quicksort(0, len(a)-1, a)
}

func quicksort(l, r int, a []int) {
	if l >= r {
		return
	}
	q := partition(l, r, a)
	quicksort(l, q, a)
	quicksort(q+1, r, a)
}

func partition(l, r int, a []int) int {
	v := (a)[(l+r)/2]
	i, j := l, r
	for i <= j {
		for a[i] < v {
			i++
		}
		for a[j] > v {
			j--
		}
		if i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	return j
}
