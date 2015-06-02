package chained

import (
<<<<<<< HEAD
	"unsafe"
)

//此数列借鉴自SGI STL
var size_primes = []uint{
	53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func nextSize(size uint) (newsz uint, ok bool) {
	var start, end = 0, len(size_primes)
	for start < end {
		var mid = (start + end) / 2
		if size < size_primes[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(size_primes) {
		return size, false
	}
	return size_primes[start], true
}

=======
	"HashTable/hash"
)

>>>>>>> tmp
type node struct {
	key  string
	next *node
}
<<<<<<< HEAD

func fakeHead(this **node) *node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).next)
	return (*node)(unsafe.Pointer(base - off))
}

type HashTable struct {
=======
type hashTable struct {
>>>>>>> tmp
	hash   func(str string) uint
	bucket []*node
	cnt    int
}

<<<<<<< HEAD
func (table *HashTable) Size() int {
	return table.cnt
}
func (table *HashTable) IsEmpty() bool {
	return table.cnt == 0
}
func (table *HashTable) isCrowded() bool {
	return table.cnt*2 > len(table.bucket)*3
}

func (table *HashTable) Initialize(fn func(str string) uint) {
	table.cnt = 0
	table.hash = fn
	table.bucket = make([]*node, size_primes[0])
=======
func (tb *hashTable) Size() int {
	return tb.cnt
}
func (tb *hashTable) IsEmpty() bool {
	return tb.cnt == 0
}
func (tb *hashTable) isCrowded() bool {
	return tb.cnt*2 > len(tb.bucket)*3
}

func (tb *hashTable) initialize(fn func(str string) uint) {
	tb.cnt = 0
	tb.hash = fn
	tb.bucket = make([]*node, primes[0])
}
func NewHashTable(fn func(str string) uint) hash.HashTable {
	var tb = new(hashTable)
	tb.initialize(fn)
	return tb
}

//此数列借鉴自SGI STL
var primes = []uint{
	17, 29, 53, 97, 193, 389, 769, 1543, 3079, 6151, 12289, 24593, 49157, 98317, 196613,
	393241, 786433, 1572869, 3145739, 6291469, 12582917, 25165843, 50331653, 1610612741}

func nextSize(size uint) (newsz uint, ok bool) {
	var start, end = 0, len(primes)
	for start < end {
		var mid = (start + end) / 2
		if size < primes[mid] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	if start == len(primes) {
		return size, false
	}
	return primes[start], true
>>>>>>> tmp
}
