package list

import (
	"unsafe"
)

type Node struct {
	Val  int
	Next *Node
}

<<<<<<< HEAD
func FakeHead(this **Node) *Node {
	var base = uintptr(unsafe.Pointer(this))
	var off = unsafe.Offsetof((*this).Next)
=======
func FakeHead(spt **Node) *Node {
	var base = uintptr(unsafe.Pointer(spt))
	var off = unsafe.Offsetof((*spt).Next)
>>>>>>> tmp
	return (*Node)(unsafe.Pointer(base - off))
}

func Merge(list1 *Node, list2 *Node) *Node {
	var head *Node = nil
	var tail = FakeHead(&head)
	for ; list1 != nil && list2 != nil; tail = tail.Next {
		if list1.Val > list2.Val {
			tail.Next, list2 = list2, list2.Next
		} else {
			tail.Next, list1 = list1, list1.Next
		}
	}
	for ; list1 != nil; tail = tail.Next {
		tail.Next, list1 = list1, list1.Next
	}
	for ; list2 != nil; tail = tail.Next {
		tail.Next, list2 = list2, list2.Next
	}
	return head
}

func Reverse(list *Node) *Node {
	var head *Node = nil
	for list != nil {
		var node = list
		list = list.Next
<<<<<<< HEAD
		node.Next = head
		head = node
=======
		node.Next, head = head, node
>>>>>>> tmp
	}
	return head
}
