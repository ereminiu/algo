package fastmergesort

import "sync"

var limit = 2048

func merge(a *[]int, b *[]int) []int {
	n, m := len(*a), len(*b)
	i, j, k := 0, 0, 0
	ret := make([]int, n+m)
	for i < n && j < m {
		if (*a)[i] < (*b)[j] {
			ret[k] = (*a)[i]
			i++
		} else {
			ret[k] = (*b)[j]
			j++
		}
		k++
	}
	for ; i < n; i++ {
		ret[k] = (*a)[i]
		k++
	}
	for ; j < m; j++ {
		ret[k] = (*b)[j]
		k++
	}
	return ret
}

func devideV1(nums *[]int, l, r int) []int {
	if l == r-1 {
		return []int{(*nums)[l]}
	}
	mid := (l + r) / 2
	left := devideV1(nums, l, mid)
	right := devideV1(nums, mid, r)
	return merge(&left, &right)
}

func devideV2(a *[]int, l, r int) []int {
	if l == r-1 {
		return []int{(*a)[l]}
	}

	if r-l+1 < limit {
		return devideV1(a, l, r)
	}

	mid := (l + r) / 2

	var wg sync.WaitGroup
	wg.Add(1)

	var left []int
	var right []int

	go func() {
		defer wg.Done()
		left = devideV2(a, l, mid)
	}()

	right = devideV2(a, mid, r)

	wg.Wait()
	return merge(&left, &right)
}

//usage: a = Sort(&a), where a is []int
func Sort(nums *[]int) []int {
	return devideV2(nums, 0, len(*nums))
}
