package array

type Queue interface {
	Clear()
	IsFull() bool
	IsEmpty() bool
	Push(key int) (fail bool)
	Pop() (key int, fail bool)
	Front() (key int, fail bool)
}

<<<<<<< HEAD
func NewQueue(lv uint) Queue {
	var core = new(queue)
	core.initialize(lv)
	return core
}

=======
func NewQueue(size int) Queue {
	var core = new(queue)
	core.initialize(size)
	return core
}

//实际容量为len(space)-1
>>>>>>> tmp
type queue struct {
	space    []int
	mask     int
	rpt, wpt int
}

<<<<<<< HEAD
func (this *queue) initialize(lv uint) {
	if lv < 4 {
		lv = 4
	}
	this.space = make([]int, lv)
	this.mask = ^(int(-1) << lv)
	this.Clear()
}
func (this *queue) Clear() {
	this.rpt, this.wpt = 0, 0
}

func (this *queue) IsFull() bool {
	var next = (this.wpt + 1) & this.mask
	return next == this.rpt
}
func (this *queue) IsEmpty() bool {
	return this.rpt == this.wpt
}

func (this *queue) Push(key int) (fail bool) {
	var next = (this.wpt + 1) & this.mask
	if next == this.rpt {
		return true
	}
	this.space[this.wpt] = key
	this.wpt = next
	return false
}

func (this *queue) Front() (key int, fail bool) {
	return this.space[this.rpt], this.IsEmpty()
}
func (this *queue) Pop() (key int, fail bool) {
	if this.IsEmpty() {
		return 0, true
	}
	key, this.rpt = this.space[this.rpt], (this.rpt+1)&this.mask
	return key, fail
=======
func (q *queue) initialize(size int) {
	if size > 0x10000 {
		size = 0x10000
	} else if size < 4 {
		size = 4
	}
	var sz = 4
	for sz < size {
		sz *= 2
	}
	q.space = make([]int, sz)
	q.mask = sz - 1
	q.Clear()
}
func (q *queue) Clear() {
	q.rpt, q.wpt = 0, 0
}

func (q *queue) IsFull() bool {
	var next = (q.wpt + 1) & q.mask
	return next == q.rpt
}
func (q *queue) IsEmpty() bool {
	return q.rpt == q.wpt
}

func (q *queue) Push(key int) (fail bool) {
	var next = (q.wpt + 1) & q.mask
	if next == q.rpt {
		return true
	}
	q.space[q.wpt] = key
	q.wpt = next
	return false
}

func (q *queue) Front() (key int, fail bool) {
	return q.space[q.rpt], q.IsEmpty()
}
func (q *queue) Pop() (key int, fail bool) {
	if q.IsEmpty() {
		return 0, true
	}
	key, q.rpt = q.space[q.rpt], (q.rpt+1)&q.mask
	return key, false
>>>>>>> tmp
}
