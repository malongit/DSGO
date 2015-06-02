<<<<<<< HEAD
package binomialheap
=======
package binomial
>>>>>>> tmp

import (
	"unsafe"
)

<<<<<<< HEAD
func fakeHead(this **node) *node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).brother)
	return (*node)(unsafe.Pointer(base - off))
}

func (heap *Heap) merge(list *node) {
	var knot = fakeHead(&heap.list)
	for knot.brother != nil && list != nil {
		if knot.brother.height == list.height {
			var root = knot.brother
			knot.brother = root.brother
			var another = list
			list = another.brother
=======
func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).peer)
	return (*node)(unsafe.Pointer(base - off))
}

//list是从少到多的
func (hp *Heap) merge(list *node) {
	var knot = fakeHead(&hp.list)
	for knot.peer != nil && list != nil {
		if knot.peer.level == list.level {
			var root = knot.peer
			knot.peer = root.peer
			var another = list
			list = another.peer
>>>>>>> tmp

			if root.key > another.key {
				root, another = another, root
			}
<<<<<<< HEAD
			another.brother, root.child = root.child, another
			root.height++

			root.brother = list
			list = root
		} else {
			if knot.brother.height > list.height {
				var target = list
				list = list.brother
				target.brother = knot.brother
				knot.brother = target
			}
			knot = knot.brother
		}
	}
	if list != nil {
		knot.brother = list
	}
}
func (heap *Heap) Merge(victim *Heap) {
	if heap != victim {
		heap.merge(victim.list)
		victim.list = nil
=======
			another.peer, root.child = root.child, another
			root.level++

			root.peer, list = list, root
		} else {
			if knot.peer.level > list.level {
				var target = list
				list = list.peer
				target.peer, knot.peer = knot.peer, target
			}
			knot = knot.peer
		}
	}
	if list != nil {
		knot.peer = list
	}
}
func (hp *Heap) Merge(victim *Heap) {
	if hp != victim {
		if hp.top.key > victim.top.key {
			hp.top = victim.top
		}
		hp.merge(victim.list)
		victim.list, victim.top = nil, nil
>>>>>>> tmp
	}
}
