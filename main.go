package main

import (
	"fmt"
	"sort"
	"time"

	"github.com/ereminiu/algo/fastmergesort"
	"github.com/ereminiu/algo/quicksortv3"
	"github.com/ereminiu/algo/utils"
)

func averageTimeBuiltin(times int) time.Duration {
	ret := time.Duration(0)

	for i := 0; i < times; i++ {
		start := time.Now()

		N := 1000000
		a := utils.RandSlice(N)
		sort.Ints(a)

		ret += time.Now().Sub(start)
	}

	return ret / time.Duration(times)
}

func averageTimeQuick(times int) time.Duration {
	ret := time.Duration(0)

	for i := 0; i < times; i++ {
		start := time.Now()

		N := 1000000
		a := utils.RandSlice(N)
		quicksortv3.Sort(&a)

		ret += time.Now().Sub(start)
	}

	return ret / time.Duration(times)
}

func averageTimeMerge(times int) time.Duration {
	ret := time.Duration(0)

	for i := 0; i < times; i++ {
		start := time.Now()

		N := 1000000
		a := utils.RandSlice(N)
		fastmergesort.Sort(&a)

		ret += time.Now().Sub(start)
	}

	return ret / time.Duration(times)
}

func main() {
	// a := utils.RandSlice(8)
	// fmt.Println(a)

	// Sort(&a)

	// fmt.Println(a)

	fmt.Println(averageTimeQuick(12))
	fmt.Println(averageTimeMerge(12))
	fmt.Println(averageTimeBuiltin(12))
}
