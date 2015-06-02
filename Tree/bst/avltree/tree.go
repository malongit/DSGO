package avltree

type node struct {
<<<<<<< HEAD
	key     int
	balance int8
	left    *node
	right   *node
}

type Tree struct {
	root *node
	path stack
}

func (tree *Tree) Search(key int) bool {
	var target = tree.root
=======
	key    int32
	state  int8 //(2), 1, 0, -1, (-2)
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
		}
	}
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
>>>>>>> tmp
}
