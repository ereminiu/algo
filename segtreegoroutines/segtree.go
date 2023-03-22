package segtreegoroutines

import "sync"

var limit int = 1 << 11
var net int = 0 // net is neutral value that doesn't make sense f(x, net) = f(net, x) = x

type Node struct {
	val int
}

func merge(x, y Node) Node {
	return Node{max(x.val, y.val)}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type Tree struct {
	tr []Node
	n  int
}

func NewTree(a *[]int) Tree {
	n := len(*a)
	t := Tree{make([]Node, 4*n), n}
	t.build(0, 0, n, a)
	return t
}

func (t *Tree) build(v, vl, vr int, a *[]int) {
	if vr-vl == 1 {
		t.tr[v] = Node{(*a)[vl]}
		return
	}
	mid := (vl + vr) / 2
	t.build(2*v+1, vl, mid, a)
	t.build(2*v+2, mid, vr, a)
	t.tr[v] = merge(t.tr[2*v+1], t.tr[2*v+2])
}

func (t *Tree) getMaxv1(v, vl, vr, l, r int) Node {
	if vl >= l && vr <= r {
		return t.tr[v]
	}
	if vl >= r || vr <= l {
		return Node{net}
	}
	mid := (vl + vr) / 2
	lf, rg := t.getMaxv1(2*v+1, vl, mid, l, r), t.getMaxv1(2*v+2, mid, vr, l, r)
	return merge(lf, rg)
}

func (t *Tree) getMaxv2(v, vl, vr, l, r int) Node {
	if vl >= l && vr <= r {
		return t.tr[v]
	}
	if vl >= r || vr <= l {
		return Node{net}
	}

	if vr-vl < limit {
		return t.getMaxv1(v, vl, vr, l, r)
	}

	mid := (vl + vr) / 2

	var wg sync.WaitGroup
	wg.Add(1)
	lf := Node{net}

	go func() {
		defer wg.Done()
		lf = t.getMaxv2(2*v+1, vl, mid, l, r)
	}()

	wg.Wait()
	rg := t.getMaxv2(2*v+2, mid, vr, l, r)
	return merge(lf, rg)
}

func (t *Tree) GetMax(l, r int) int {
	return t.getMaxv1(0, 0, t.n, l, r).val
}

func (t *Tree) update(v, vl, vr, idx, val int) {
	if vr-vl == 1 {
		t.tr[v] = Node{val}
		return
	}
	mid := (vl + vr) / 2
	if idx < mid {
		t.update(2*v+1, vl, mid, idx, val)
	} else {
		t.update(2*v+2, mid, vr, idx, val)
	}
	t.tr[v] = merge(t.tr[2*v+1], t.tr[2*v+2])
}

func (t *Tree) Update(idx, val int) {
	t.update(0, 0, t.n, idx, val)
}
