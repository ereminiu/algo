package utils

import "math/rand"

func RandSlice(n int) []int {
	ret := rand.Perm(n)
	for i := 0; i < n; i++ {
		ret[i] += rand.Intn(10)
	}
	return ret
}
