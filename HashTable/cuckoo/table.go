package cuckoo

<<<<<<< HEAD
type node struct {
	code [2]uint
	key  string
}
type coreTable struct {
	hash   func(str string) uint
	mask   uint
	id     uint8 //0或1
	bucket []*node
}
type HashTable struct {
	first, second coreTable
	cnt           int
}

func (table *HashTable) Size() int {
	return table.cnt
}
func (table *HashTable) IsEmpty() bool {
	return table.cnt == 0
}

func (table *HashTable) Initialize(fn1 func(str string) uint, fn2 func(str string) uint) {
	table.cnt = 0
	table.first.hash, table.second.hash = fn1, fn2
	table.first.id, table.second.id = 0, 1
	table.first.bucket, table.second.bucket = make([]*node, 32), make([]*node, 16) //保持first为second的两倍
	table.first.mask, table.second.mask = 0x1f, 0xf
=======
import (
	"HashTable/hash"
)

const WAYS = 3

type node struct {
	code [WAYS]uint
	key  string
}
type table struct {
	hash   func(str string) uint
	bucket []*node
}
type hashTable struct {
	core [WAYS]table
	idx  int
	cnt  int
}

func (tb *hashTable) Size() int {
	return tb.cnt
}
func (tb *hashTable) IsEmpty() bool {
	return tb.cnt == 0
}

func (tb *hashTable) initialize(fn [WAYS]func(str string) uint) {
	tb.idx, tb.cnt = 0, 0
	var sz = 8 //2^n
	for i := WAYS - 1; i >= 0; i-- {
		tb.core[i].hash = fn[i]
		tb.core[i].bucket = make([]*node, sz)
		sz *= 2
	}
}
func NewHashTable(fn [WAYS]func(str string) uint) hash.HashTable {
	var tb = new(hashTable)
	tb.initialize(fn)
	return tb
>>>>>>> tmp
}
