<<<<<<< HEAD
package binaryheap

func (heap *Heap) adjustDown(spot int) {
	var size = len(heap.core)
	var key = heap.core[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid = 0
		if heap.core[left] < heap.core[right] {
=======
package binary

func (hp *Heap) adjustDown(spot int) {
	var size = len(hp.core)
	var key = hp.core[spot]
	var left, right = spot*2 + 1, spot*2 + 2
	for right < size {
		var kid = 0
		if hp.core[left] < hp.core[right] {
>>>>>>> tmp
			kid = left
		} else {
			kid = right
		}
<<<<<<< HEAD
		if key <= heap.core[kid] {
			goto LabelOver
		}
		heap.core[spot] = heap.core[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key > heap.core[left] {
		heap.core[spot], heap.core[left] = heap.core[left], key
		return
	}
LabelOver:
	heap.core[spot] = key
}

func (heap *Heap) adjustUp(spot int) {
	var key = heap.core[spot]
	for spot > 0 {
		var parent = (spot - 1) / 2
		if heap.core[parent] <= key {
			break
		}
		heap.core[spot], spot = heap.core[parent], parent
	}
	heap.core[spot] = key
=======
		if key <= hp.core[kid] {
			goto LabelOver
		}
		hp.core[spot] = hp.core[kid]
		spot, left, right = kid, kid*2+1, kid*2+2
	}
	if right == size && key > hp.core[left] {
		hp.core[spot], hp.core[left] = hp.core[left], key
		return
	}
LabelOver:
	hp.core[spot] = key
}

func (hp *Heap) adjustUp(spot int) {
	var key = hp.core[spot]
	for spot > 0 {
		var parent = (spot - 1) / 2
		if hp.core[parent] <= key {
			break
		}
		hp.core[spot], spot = hp.core[parent], parent
	}
	hp.core[spot] = key
>>>>>>> tmp
}
