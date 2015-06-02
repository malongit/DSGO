package bst

import (
	"testing"
)

<<<<<<< HEAD
func benchInsert(b *testing.B, hint int) {
	b.StopTimer()
	var tree, list = newTree(hint), mixArray(b.N)
=======
func Test_Nothing(t *testing.T) {
	//...
}

func Benchmark_SimpleInsert(b *testing.B) {
	benchInsert(b, SIMPLE_BST)
}
func Benchmark_SimpleSearch(b *testing.B) {
	benchSearch(b, SIMPLE_BST)
}
func Benchmark_SimpleRemove(b *testing.B) {
	benchRemove(b, SIMPLE_BST)
}
func Benchmark_AVLtreeInsert(b *testing.B) {
	benchInsert(b, AVL_TREE)
}
func Benchmark_AVLtreeSearch(b *testing.B) {
	benchSearch(b, AVL_TREE)
}
func Benchmark_AVLtreeRemove(b *testing.B) {
	benchRemove(b, AVL_TREE)
}
func Benchmark_RBtreeInsert(b *testing.B) {
	benchInsert(b, RB_TREE)
}
func Benchmark_RBtreeSearch(b *testing.B) {
	benchSearch(b, RB_TREE)
}
func Benchmark_RBtreeRemove(b *testing.B) {
	benchRemove(b, RB_TREE)
}

func benchInsert(b *testing.B, hint int) {
	b.StopTimer()
	var tree, list = newTree(hint), mixedArray(b.N)
>>>>>>> tmp
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
}
<<<<<<< HEAD
func benchFind(b *testing.B, hint int) {
	b.StopTimer()
	var tree, list = newTree(hint), mixArray(b.N)
=======
func benchSearch(b *testing.B, hint int) {
	b.StopTimer()
	var tree, list = newTree(hint), mixedArray(b.N)
>>>>>>> tmp
	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		tree.Search(list[i])
	}
}
func benchRemove(b *testing.B, hint int) {
	b.StopTimer()
<<<<<<< HEAD
	var tree, list = newTree(hint), mixArray(b.N)
=======
	var tree, list = newTree(hint), mixedArray(b.N)
>>>>>>> tmp
	for i := 0; i < b.N; i++ {
		tree.Insert(list[i])
	}
	b.StartTimer()
	//for i := 0; i < b.N; i++ { //对非平衡树而言删除顺序重要
	for i := b.N - 1; i >= 0; i-- {
		tree.Remove(list[i])
	}
}
<<<<<<< HEAD

func Benchmark_SimpleInsert(b *testing.B) {
	benchInsert(b, SIMPLE_BST)
}
func Benchmark_SimpleFind(b *testing.B) {
	benchFind(b, SIMPLE_BST)
}
func Benchmark_SimpleRemove(b *testing.B) {
	benchRemove(b, SIMPLE_BST)
}
func Benchmark_AVLtreeInsert(b *testing.B) {
	benchInsert(b, AVL_TREE)
}
func Benchmark_AVLtreeFind(b *testing.B) {
	benchFind(b, AVL_TREE)
}
func Benchmark_AVLtreeRemove(b *testing.B) {
	benchRemove(b, AVL_TREE)
}
func Benchmark_RBtreeInsert(b *testing.B) {
	benchInsert(b, RB_TREE)
}
func Benchmark_RBtreeFind(b *testing.B) {
	benchFind(b, RB_TREE)
}
func Benchmark_RBtreeRemove(b *testing.B) {
	benchRemove(b, RB_TREE)
}
=======
>>>>>>> tmp
