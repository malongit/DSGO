package bptree

import (
	"math/rand"
	"testing"
	"time"
)

<<<<<<< HEAD
var tree_ut Tree
var dirty_ut = false

const size_ut = 200

var list1_ut [size_ut]int
var list2_ut [size_ut]int
var list3_ut [size_ut]int
var list4_ut [size_ut]int

func prepareTestData() {
	rand.Seed(time.Now().Unix())
	for i := 0; i < size_ut; i++ {
		list1_ut[i] = rand.Int()
		list2_ut[i] = rand.Int()
		list3_ut[i] = rand.Int()
		list4_ut[i] = rand.Int()
	}
}

func detectAndMark(t *testing.T) {
	if err := recover(); err != nil {
		dirty_ut = true
		t.Fail()
	}
}
func markFail(t *testing.T) {
	dirty_ut = true
	t.FailNow()
}

func Test_Insert(t *testing.T) {
	defer detectAndMark(t)
	prepareTestData()

	for i := 0; i < size_ut; i++ {
		tree_ut.Insert(list1_ut[i])
		tree_ut.Insert(list2_ut[i])
		tree_ut.Insert(list3_ut[i])
		tree_ut.Insert(list4_ut[i])
	}
	for i := 0; i < size_ut; i++ {
		if tree_ut.Insert(list2_ut[i]) {
			markFail(t)
		}
	}
}

func Test_Search(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		if !tree_ut.Search(list4_ut[i]) ||
			!tree_ut.Search(list3_ut[i]) ||
			!tree_ut.Search(list2_ut[i]) ||
			!tree_ut.Search(list1_ut[i]) {
			markFail(t)
		}
	}
}

func Test_Remove(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		tree_ut.Remove(list3_ut[i])
		tree_ut.Remove(list4_ut[i])
	}
	for i := 0; i < size_ut; i++ {
		tree_ut.Remove(list1_ut[i])
		tree_ut.Remove(list2_ut[i])
	}
}

func Test_InsertAfterRemove(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		tree_ut.Insert(list1_ut[i])
		tree_ut.Insert(list3_ut[i])
	}
}

func Test_FindAfterRemove(t *testing.T) {
	if dirty_ut {
		t.SkipNow()
	}
	defer detectAndMark(t)

	for i := 0; i < size_ut; i++ {
		if !tree_ut.Search(list1_ut[i]) {
			markFail(t)
		}
	}
	for i := 0; i < size_ut; i++ {
		if tree_ut.Search(list2_ut[i]) ||
			tree_ut.Search(list4_ut[i]) {
			markFail(t)
		}
	}
	for i := 0; i < size_ut; i++ {
		if !tree_ut.Search(list3_ut[i]) {
			markFail(t)
		}
	}
=======
func Test_Tree(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var tree Tree
	var cnt = 0
	const size = 2000
	var list [size * 2]int

	rand.Seed(time.Now().Unix())
	for i := 0; i < size*2; i++ {
		list[i] = rand.Int()
	}

	//插入两份
	for i := 0; i < size*2; i++ {
		if tree.Insert(list[i]) {
			cnt++
		}
	}
	if tree.Insert(list[size]) {
		t.Fail()
	}
	for i := 0; i < size*2; i++ {
		if !tree.Search(list[i]) {
			t.Fail()
		}
	}

	//遍历
	mark_ut = -(int(^uint(0) >> 1)) - 1
	tree.Travel(checkNum)

	//删除第一份
	for i := 0; i < size; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
	}
	for i := 0; i < size; i++ {
		if tree.Search(list[i]) {
			t.Fail()
		}
	}

	//删除第二份
	for i := size; i < size*2; i++ {
		if tree.Remove(list[i]) {
			cnt--
		}
	}
	if !tree.IsEmpty() || cnt != 0 {
		t.Fail()
	}
}

var mark_ut = 0

func checkNum(val int) {
	if val < mark_ut {
		panic(val)
	}
	mark_ut = val
>>>>>>> tmp
}
