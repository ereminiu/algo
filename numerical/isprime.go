package numerical

// don't push numbers less than 0
func Prime(n int) bool {
	if n < 4 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for p := 3; p*p <= n; p += 2 {
		if n%p == 0 {
			return false
		}
	}
	return true
}
