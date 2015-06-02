<<<<<<< HEAD
package binaryheap

//二叉堆，底层采用数组，缺点是合并代价较高
=======
package binary

//二叉堆，底层采用数组。
//Build的复杂度为O(N)，Top的复杂度为O(1)，其余核心操作复杂度为O(logN)。
>>>>>>> tmp
type Heap struct {
	core []int
}

<<<<<<< HEAD
func (heap *Heap) Size() int {
	return len(heap.core)
}
func (heap *Heap) IsEmpty() bool {
	return len(heap.core) == 0
}

func (heap *Heap) Top() (key int, fail bool) {
	if heap.IsEmpty() {
		return 0, true
	}
	return heap.core[0], false
}

func (heap *Heap) Build(list []int) {
	var size = len(list)
	heap.core = list
	for idx := size/2 - 1; idx >= 0; idx-- {
		heap.adjustDown(idx)
	}
}
func (heap *Heap) Push(key int) {
	var place = len(heap.core)
	heap.core = append(heap.core, key)
	heap.adjustUp(place)
}
func (heap *Heap) Pop() (key int, fail bool) {
	var size = heap.Size()
=======
func (hp *Heap) Size() int {
	return len(hp.core)
}
func (hp *Heap) IsEmpty() bool {
	return len(hp.core) == 0
}

func (hp *Heap) Top() (key int, fail bool) {
	if hp.IsEmpty() {
		return 0, true
	}
	return hp.core[0], false
}

func (hp *Heap) Build(list []int) {
	var size = len(list)
	hp.core = list
	for idx := size/2 - 1; idx >= 0; idx-- {
		hp.adjustDown(idx)
	}
}
func (hp *Heap) Push(key int) {
	var place = len(hp.core)
	hp.core = append(hp.core, key)
	hp.adjustUp(place)
}
func (hp *Heap) Pop() (key int, fail bool) {
	var size = hp.Size()
>>>>>>> tmp
	if size == 0 {
		return 0, true
	}

<<<<<<< HEAD
	key = heap.core[0]
	if size == 1 {
		heap.core = heap.core[:0]
	} else {
		heap.core[0] = heap.core[size-1]
		heap.core = heap.core[:size-1]
		heap.adjustDown(0)
=======
	key = hp.core[0]
	if size == 1 {
		hp.core = hp.core[:0]
	} else {
		hp.core[0] = hp.core[size-1]
		hp.core = hp.core[:size-1]
		hp.adjustDown(0)
>>>>>>> tmp
	}
	return key, false
}
