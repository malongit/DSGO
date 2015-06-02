package bst

import (
<<<<<<< HEAD
	"SearchTree/BST/avltree"
	"SearchTree/BST/rbtree"
	"SearchTree/BST/simplebst"
=======
	"Tree/bst/avltree"
	"Tree/bst/compact/avlt"
	"Tree/bst/compact/rbt"
	"Tree/bst/rbtree"
	"Tree/bst/simplebst"
>>>>>>> tmp
	"fmt"
	"math/rand"
	"time"
	"unsafe"
)

type BST interface {
<<<<<<< HEAD
	Insert(key int) bool
	Search(key int) bool
	Remove(key int) bool
=======
	IsEmpty() bool
	Insert(key int32) bool
	Search(key int32) bool
	Remove(key int32) bool
>>>>>>> tmp
}

const (
	SIMPLE_BST = iota
	AVL_TREE
	RB_TREE
<<<<<<< HEAD
=======
	PLAIN_AVL
	PLAIN_RB
>>>>>>> tmp
)

func showName(hint int) string {
	switch hint {
	case SIMPLE_BST:
		return "simple BST"
	case AVL_TREE:
		return "AVL tree"
	case RB_TREE:
		return "red-black tree"
<<<<<<< HEAD
=======
	case PLAIN_AVL:
		return "AVL tree [no uplink]"
	case PLAIN_RB:
		return "red-black [no uplink]"
>>>>>>> tmp
	default:
		panic("Illegal BST type")
	}
}
func newTree(hint int) BST {
	switch hint {
	case SIMPLE_BST:
		return new(simplebst.Tree)
	case AVL_TREE:
		return new(avltree.Tree)
	case RB_TREE:
		return new(rbtree.Tree)
<<<<<<< HEAD
=======
	case PLAIN_AVL:
		return new(avlt.Tree)
	case PLAIN_RB:
		return new(rbt.Tree)
>>>>>>> tmp
	default:
		panic("Illegal BST type")
	}
}

<<<<<<< HEAD
func mixArray(size int) []int {
	var list = make([]int, size)

	const bits_of_int = uint(unsafe.Sizeof(size)) * 8
=======
func mixedArray(size int) []int32 {
	var list = make([]int32, size)

	const bits_of_int = uint(unsafe.Sizeof(list[0])) * 8
>>>>>>> tmp
	var tmp = uint(size)
	var cnt uint = 0
	for cnt < bits_of_int && tmp != 0 {
		cnt++
		tmp >>= 1
	}
	cnt = bits_of_int - cnt - 11
	var mask = ^((^0) << cnt)

<<<<<<< HEAD
	var num = 0
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		if i%32 == 0 { //局部参入有序数列
			num += rand.Int() & mask
			list[i] = num
		} else {
			list[i] = rand.Int()
=======
	var num = int32(0)
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		if i%32 == 0 { //局部参入有序数列
			num += int32(rand.Int() & mask)
			list[i] = num
		} else {
			list[i] = int32(rand.Int())
>>>>>>> tmp
		}
	}
	return list
}

<<<<<<< HEAD
func DoOneTry(list []int, hint int) {
=======
func benchMark(list []int32, hint int) {
>>>>>>> tmp
	var tree = newTree(hint)
	var size = len(list)
	fmt.Println(showName(hint))

	var start = time.Now()
	for i := 0; i < size; i++ {
		tree.Insert(list[i])
	}
	fmt.Println("insert:", time.Since(start))
	start = time.Now()
	for i := 0; i < size; i++ {
		tree.Search(list[i])
	}
	fmt.Println("search:", time.Since(start))
	start = time.Now()
	//for i := 0; i < size; i++ { //对非平衡树而言删除顺序重要
	for i := size - 1; i >= 0; i-- {
		tree.Remove(list[i])
	}
	fmt.Println("remove:", time.Since(start))
}

<<<<<<< HEAD
func DoBenchMark(size int) {
	if size < 1000 {
		fmt.Println("too small")
		return
	}
	var list = mixArray(size)

	DoOneTry(list, SIMPLE_BST)
	DoOneTry(list, AVL_TREE)
	DoOneTry(list, RB_TREE)
=======
func BenchMark(size int) {
	if size < 10000 {
		fmt.Println("too small")
		return
	}
	var list = mixedArray(size)

	benchMark(list, AVL_TREE)
	benchMark(list, RB_TREE)
	benchMark(list, PLAIN_AVL)
	benchMark(list, PLAIN_RB)
>>>>>>> tmp
}
