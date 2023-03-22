package main

import (
	"fmt"
	"math/rand"
	"time"

	segtree "github.com/ereminiu/algo/segmenttree"
	"github.com/ereminiu/algo/utils"
)

func averageTimeSegtree(times int) time.Duration {
	start := time.Now()

	for it := 0; it < 100; it++ {
		n := 150000
		a := utils.RandSlice(n)
		st := segtree.NewTree(&a)
		k := rand.Intn(100)
		if k%2 == 0 {
			idx := rand.Intn(n)
			st.Update(idx, 19)
			a[idx] = 19
		} else {
			l, r := rand.Intn(n-1), rand.Intn(n)
			for r <= l {
				r = rand.Intn(n)
			}
			_ = st.GetMax(l, r)
		}
	}

	return time.Now().Sub(start)
}

func averageTimeSegtreeWithGoroutines(times int) time.Duration {
	start := time.Now()

	for it := 0; it < 100; it++ {
		n := 150000
		a := utils.RandSlice(n)
		st := segtree.NewTree(&a)
		k := rand.Intn(100)
		if k%2 == 0 {
			idx := rand.Intn(n)
			st.Update(idx, 19)
			a[idx] = 19
		} else {
			l, r := rand.Intn(n-1), rand.Intn(n)
			for r <= l {
				r = rand.Intn(n)
			}
			_ = st.GetMax(l, r)
		}
	}

	return time.Now().Sub(start)
}

func main() {
	fmt.Println(averageTimeSegtree(150))
	fmt.Println(averageTimeSegtreeWithGoroutines(150))
}
