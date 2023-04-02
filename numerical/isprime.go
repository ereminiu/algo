package numerical

import "golang.org/x/exp/constraints"

// don't push numbers less than 0
func Prime[T constraints.Signed](n T) bool {
	if n < 4 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for p := T(3); p*p <= n; p += 2 {
		if n%p == 0 {
			return false
		}
	}
	return true
}
