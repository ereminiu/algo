package numerical

func sieve(n int) []int {
	n++
	isprime := make([]bool, n)
	for i := 0; i < n; i++ {
		isprime[i] = true
	}
	isprime[0], isprime[1] = false, false
	ans := make([]int, 0)
	for i := 2; i < n; i++ {
		if !isprime[i] {
			continue
		}
		ans = append(ans, i)
		for j := i * i; j < n; j += i {
			isprime[j] = false
		}
	}

	return ans
}
