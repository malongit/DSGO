<<<<<<< HEAD
package pairingheap
=======
package pairing
>>>>>>> tmp

import (
	"unsafe"
)

func fakeHead(spt **Node) *Node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*Node)(unsafe.Pointer(base - off))
}

<<<<<<< HEAD
func (heap *Heap) PopNode() (unit *Node) {
	if heap.root == nil {
		return nil
	}
	unit, heap.root = heap.root, heap.root.child
	if heap.root == nil {
=======
func (hp *Heap) PopNode() (unit *Node) {
	if hp.root == nil {
		return nil
	}
	unit, hp.root = hp.root, hp.root.child
	if hp.root == nil {
>>>>>>> tmp
		return
	}
	//一次整理最坏情况下代价为O(N)，摊还代价则为O(log N)
	//这里采用线性聚拢是不合适的，复杂之余不能持久降低宽度
<<<<<<< HEAD
	for heap.root.next != nil {
		var list, knot = heap.root, fakeHead(&heap.root)
=======
	for hp.root.next != nil {
		var list, knot = hp.root, fakeHead(&hp.root)
>>>>>>> tmp
		for list != nil && list.next != nil { //两两配对
			var one, another = list, list.next
			list = another.next
			knot.next = merge(one, another)
			knot = knot.next
		}
		knot.next = list
	}
<<<<<<< HEAD
	heap.root.prev = nil
	return
}
func (heap *Heap) Pop() (key int, fail bool) {
	var node = heap.PopNode()
=======
	hp.root.prev = nil
	return
}
func (hp *Heap) Pop() (key int, fail bool) {
	var node = hp.PopNode()
>>>>>>> tmp
	if node == nil {
		return 0, true
	}
	return node.key, false
}
