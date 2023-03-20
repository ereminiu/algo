package mergesort

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

func devide(nums *[]int, l, r int) []int {
	if l == r-1 {
		return []int{(*nums)[l]}
	}
	mid := (l + r) / 2
	left := devide(nums, l, mid)
	right := devide(nums, mid, r)
	return merge(&left, &right)
}

func Sort(nums *[]int) []int {
	return devide(nums, 0, len(*nums))
}
