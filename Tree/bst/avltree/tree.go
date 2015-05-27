package avltree

type node struct {
	key    int32
	state  int8 //(2), 1, 0, -1, (-2)
	parent *node
	left   *node
	right  *node
}
type Tree struct {
	root *node
}

func (tree *Tree) IsEmpty() bool {
	return tree.root == nil
}

func (tree *Tree) Search(key int32) bool {
	var target = tree.root
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
