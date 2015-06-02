<<<<<<< HEAD
package heap

import (
	"testing"
)

func Test_Nothing(t *testing.T) {
	//Do Nothing
}

/*

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const size = 16
	rand.Seed(time.Now().Unix())

	var list [size]int
	for i := 0; i < size; i++ {
		var key = rand.Int() % 1000
		fmt.Printf("%d ", key)
		list[i] = key
	}

	var heap BinaryHeap
	heap.Build(list[:])
	for i := 0; i < size; i++ {
		var key = rand.Int() % 1000
		fmt.Printf("%d ", key)
		heap.Push(key)
	}
	fmt.Println()

	for !heap.IsEmpty() {
		fmt.Printf("%d ", heap.Pop())
	}
	fmt.Println()

}
*/
=======
package binomial

import (
	"math/rand"
	"testing"
	"time"
)

func Test_Heap(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			t.Fail()
		}
	}()
	var heap, another Heap
	const size = 200
	var list [size * 2]int

	const INT_MAX = int(^uint(0) >> 1)
	var mark = INT_MAX

	rand.Seed(time.Now().Unix())
	for i := 0; i < size*2; i++ {
		list[i] = rand.Int()
		if list[i] < mark {
			mark = list[i]
		}
	}

	//插入
	for i := 0; i < size; i++ {
		heap.Push(list[i])
		another.Push(list[size+i])
	}

	//合并
	heap.Merge(&another)
	var key, fail = heap.Top()
	if fail || key != mark || !another.IsEmpty() {
		t.Fail()
	}

	//删除
	for i := 0; i < size*2; i++ {
		key, fail = heap.Pop()
		if fail || key < mark {
			t.Fail()
		}
		mark = key
	}
	if !heap.IsEmpty() {
		t.Fail()
	}
}
>>>>>>> tmp
