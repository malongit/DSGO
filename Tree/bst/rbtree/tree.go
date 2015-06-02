package rbtree

<<<<<<< HEAD
type node struct {
	key   int
	black bool
	left  *node
	right *node
}

type Tree struct {
	root *node
	path stack
}

func (tree *Tree) Search(key int) bool {
	var target = tree.root
=======
//AVL树的平衡因子有5态，需要3bit存储空间。
//而红黑树的平衡因子只需1bit，有时候可以巧妙地隐藏掉。
type node struct {
	key    int32
	black  bool
	parent *node
	left   *node
	right  *node
}
type Tree struct {
	root *node
}

func (tr *Tree) IsEmpty() bool {
	return tr.root == nil
}

func (tr *Tree) Search(key int32) bool {
	var target = tr.root
>>>>>>> tmp
	for target != nil {
		if key == target.key {
			return true
		}
		if key < target.key {
			target = target.left
		} else {
			target = target.right
		}
	}
	return false
}

<<<<<<< HEAD
func (tree *Tree) hookSubTree(subtree *node) {
	if tree.path.isEmpty() {
		tree.root = subtree
	} else {
		if super, lf := tree.path.top(); lf {
			super.left = subtree
		} else {
			super.right = subtree
=======
func (parent *node) tryHook(child *node) *node {
	if child != nil {
		child.parent = parent
	}
	return child
}
func (parent *node) hook(child *node) *node {
	child.parent = parent
	return child
}
func (tr *Tree) hookSubTree(super *node, root *node) {
	if super == nil {
		tr.root = super.hook(root)
	} else {
		if root.key < super.key {
			super.left = super.hook(root)
		} else {
			super.right = super.hook(root)
>>>>>>> tmp
		}
	}
}
