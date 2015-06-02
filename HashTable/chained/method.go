package chained

<<<<<<< HEAD
func (table *HashTable) Search(key string) bool {
	var index = table.hash(key) % uint(len(table.bucket))
	for unit := table.bucket[index]; unit != nil; unit = unit.next {
=======
import (
	"unsafe"
)

func (tb *hashTable) Search(key string) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	for unit := tb.bucket[index]; unit != nil; unit = unit.next {
>>>>>>> tmp
		if key == unit.key {
			return true
		}
	}
	return false
}

//成功返回true，没有返回false
<<<<<<< HEAD
func (table *HashTable) Remove(key string) bool {
	var index = table.hash(key) % uint(len(table.bucket))
	for knot := fakeHead(&table.bucket[index]); knot.next != nil; knot = knot.next {
		if key == knot.next.key {
			knot.next = knot.next.next
			table.cnt--
=======
func (tb *hashTable) Remove(key string) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	for knot := fakeHead(&tb.bucket[index]); knot.next != nil; knot = knot.next {
		if key == knot.next.key {
			knot.next = knot.next.next
			tb.cnt--
>>>>>>> tmp
			return true
		}
	}
	return false
}
<<<<<<< HEAD

//成功返回true，冲突返回false
func (table *HashTable) Insert(key string) bool {
	var tail = table.bucket[table.hash(key)%uint(len(table.bucket))]
	for ; tail != nil; tail = tail.next {
		if key == tail.key {
=======
func fakeHead(spt **node) *node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).next)
	return (*node)(unsafe.Pointer(base - off))
}

//成功返回true，冲突返回false
func (tb *hashTable) Insert(key string) bool {
	var index = tb.hash(key) % uint(len(tb.bucket))
	for unit := tb.bucket[index]; unit != nil; unit = unit.next {
		if key == unit.key {
>>>>>>> tmp
			return false
		}
	}
	var unit = new(node)
	unit.key = key
<<<<<<< HEAD
	tail.next = unit

	table.cnt++
	if table.isCrowded() {
		if newsz, ok := nextSize(uint(len(table.bucket))); ok {
			table.resize(newsz)
=======
	unit.next, tb.bucket[index] = tb.bucket[index], unit

	tb.cnt++
	if tb.isCrowded() {
		if newsz, ok := nextSize(uint(len(tb.bucket))); ok {
			tb.resize(newsz)
>>>>>>> tmp
		}
	}
	return true
}
<<<<<<< HEAD
func (table *HashTable) resize(size uint) {
	var old_bucket = table.bucket
	table.bucket = make([]*node, size)
	for _, unit := range old_bucket {
		for ; unit != nil; unit = unit.next {
			var index = table.hash(unit.key) % size
			unit.next, table.bucket[index] = table.bucket[index], unit
=======
func (tb *hashTable) resize(size uint) {
	var old_bucket = tb.bucket
	tb.bucket = make([]*node, size)
	for _, unit := range old_bucket {
		for unit != nil {
			var current, index = unit, tb.hash(unit.key) % size
			unit = unit.next
			current.next, tb.bucket[index] = tb.bucket[index], current
>>>>>>> tmp
		}
	}
}
