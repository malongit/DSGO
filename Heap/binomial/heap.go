<<<<<<< HEAD
package binomialheap

type node struct {
	key     int
	height  uint
	brother *node
	child   *node
}
type Heap struct {
	list *node
} //二项堆的Top和Push可以到O(1)，但本实现的Top为O(log N)

func (heap *Heap) IsEmpty() bool {
	return heap.list == nil
}

func (heap *Heap) Top() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}
	var best = heap.list.key
	for pt := heap.list.brother; pt != nil; pt = pt.brother {
		if pt.key < best {
			best = pt.key
		}
	}
	return best, false
}
func (heap *Heap) Push(key int) {
	var peer = new(node)
	peer.key = key
	peer.height = 0
	peer.brother, peer.child = nil, nil
	heap.merge(peer)
}

func reverse(head *node) *node {
	if head == nil {
		return nil
	}
	var tail = head.brother
	head.brother = nil
	for tail != nil {
		var current = tail
		tail = tail.brother
		current.brother, head = head, current
	}
	return head
}
func (heap *Heap) Pop() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}

	var knot = fakeHead(&heap.list)
	for pt := heap.list; pt.brother != nil; pt = pt.brother {
		if pt.brother.key < knot.brother.key {
			knot = pt
		}
	}
	key = knot.brother.key
	var list = reverse(knot.brother.child)
	knot.brother = knot.brother.brother
	heap.merge(list)
=======
package binomial

//二项堆的Push和Top操作的复杂度为O(1)，其余核心操作复杂度为O(logN)。
type Heap struct {
	list *node
	top  *node
}
type node struct {
	key   int
	level uint
	peer  *node
	child *node
}

func (hp *Heap) IsEmpty() bool {
	return hp.list == nil
}

func (hp *Heap) Top() (key int, fail bool) {
	if hp.IsEmpty() {
		return 0, true
	}
	return hp.top.key, false
}
func (hp *Heap) Push(key int) {
	var unit = new(node)
	unit.key, unit.level = key, 0
	unit.peer, unit.child = nil, nil
	if hp.IsEmpty() {
		hp.list, hp.top = unit, unit
	} else {
		if key < hp.top.key {
			hp.top = unit
		}
		hp.merge(unit)
	}
}

//list是从少到多的，而child相反
func reverse(list *node) *node {
	var head *node = nil
	for list != nil {
		var current = list
		list = list.peer
		current.peer, head = head, current
	}
	return head
}
func (hp *Heap) Pop() (key int, fail bool) {
	if hp.IsEmpty() {
		return 0, true
	}
	key = hp.top.key

	var knot = fakeHead(&hp.list)
	for knot.peer != hp.top {
		knot = knot.peer
	}
	knot.peer = knot.peer.peer

	hp.merge(reverse(hp.top.child))
	hp.top = hp.list
	if hp.list != nil {
		for pt := hp.list.peer; pt != nil; pt = pt.peer {
			if pt.key < hp.top.key {
				hp.top = pt
			}
		}
	}
>>>>>>> tmp
	return key, false
}
