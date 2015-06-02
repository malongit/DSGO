<<<<<<< HEAD
package skiplist
=======
package skiplist

import (
	"time"
	"unsafe"
)

type SkipList interface {
	Size() int
	IsEmpty() bool
	Insert(key int) bool
	Search(key int) bool
	Remove(key int) bool
	Travel(doit func(int))
}

const factor = 3

type node struct {
	next []*node
	key  int
}
type skipList struct {
	heads []*node
	knots []*node
	level int //非零
	mark  int //非零
	ceil  int //非零
	floor int //非零
	rand  mt19937
}

func NewSkipList() SkipList {
	var l = new(skipList)
	l.initialize()
	return l
}

func (l *skipList) initialize() {
	l.rand.initialize(uint32(time.Now().Unix()))
	l.heads, l.knots = make([]*node, 1), make([]*node, 1)
	l.level, l.mark, l.ceil, l.floor = 1, 1, factor, 1
}

func (l *skipList) IsEmpty() bool {
	return l.Size() == 0
}
func (l *skipList) Size() int {
	return l.mark - 1
}

func (l *skipList) Travel(doit func(int)) {
	for unit := l.heads[0]; unit != nil; unit = unit.next[0] {
		doit(unit.key)
	}
}
func (l *skipList) Search(key int) bool {
	var knot = (*node)(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
	}
	var target = knot.next[0]
	return target != nil && target.key == key
}

//成功返回true，冲突返回false
func (l *skipList) Insert(key int) bool {
	var knot = (*node)(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		l.knots[i] = knot
	}
	var target = knot.next[0]
	if target != nil && target.key == key {
		return false
	}

	l.mark++
	if l.mark == l.ceil {
		l.floor = l.ceil
		l.ceil *= factor
		l.level++
		l.heads = append(l.heads, nil)
		l.knots = append(l.knots, (*node)(unsafe.Pointer(l)))
	}

	var lv = 1
	for l.rand.next() <= (^uint32(0)/uint32(factor)) &&
		lv < l.level {
		lv++
	}
	target = new(node)
	target.key = key
	target.next = make([]*node, lv)
	for i := 0; i < lv; i++ {
		target.next[i] = l.knots[i].next[i]
		l.knots[i].next[i] = target
	}
	return true
}

//成功返回true，没有返回false
func (l *skipList) Remove(key int) bool {
	var knot = (*node)(unsafe.Pointer(l))
	for i := l.level - 1; i >= 0; i-- {
		for knot.next[i] != nil && knot.next[i].key < key {
			knot = knot.next[i]
		}
		l.knots[i] = knot
	}
	var target = knot.next[0]
	if target == nil || target.key != key {
		return false
	}

	var lv = min(len(target.next), l.level)
	for i := 0; i < lv; i++ {
		l.knots[i].next[i] = target.next[i]
	}

	l.mark--
	if l.mark < l.floor { //注意不能==
		l.ceil = l.floor
		l.floor /= factor
		l.level--
		l.heads = l.heads[:l.level]
		l.knots = l.knots[:l.level]
	}
	return true
}
func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}
>>>>>>> tmp
