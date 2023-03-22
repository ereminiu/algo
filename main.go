package main

import (
	"fmt"
	"math/rand"

	"github.com/ereminiu/algo/segmenttree"
	"github.com/ereminiu/algo/utils"
)

func testSegTree(times int) {
	n := 1000
	for it := 0; it < times; it++ {
		a := utils.RandSlice(n)
		st := segmenttree.NewTree(&a)

		getmax := func(l, r int) int {
			ret := 0
			for i := l; i < r; i++ {
				ret = utils.Max(ret, a[i])
			}
			return ret
		}

		for rep := 0; rep < 100; rep++ {
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
				if getmax(l, r) != st.GetMax(l, r) {
					fmt.Printf("Failed on %v %v %v\n", a, l, r)
					fmt.Printf("Should be %v instead of %v", getmax(l, r), st.GetMax(l, r))
					panic("Aboba")
				}
			}
		}

	}
	fmt.Printf("Tests passed :)\n")
}

func main() {
	// a := []int{1, 2, 3, -2, 7, 8}
	// st := segmenttree.NewTree(&a)

	// fmt.Println(st.GetMax(0, 5))
	// st.Update(0, 29)
	// fmt.Println(st.GetMax(0, 5))

	// fmt.Println(st.GetMax())

	testSegTree(1200)
}
