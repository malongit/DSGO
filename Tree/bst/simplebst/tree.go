package simplebst

type node struct {
<<<<<<< HEAD
	key   int
=======
	key   int32
>>>>>>> tmp
	left  *node
	right *node
}
type Tree struct {
	root *node
}

<<<<<<< HEAD
func (tree *Tree) Search(key int) bool {
	var target = tree.root
=======
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
func newNode(key int) (unit *node) {
=======
func newNode(key int32) (unit *node) {
>>>>>>> tmp
	unit = new(node)
	unit.key = key
	unit.left, unit.right = nil, nil
	return unit
}

//成功返回true，冲突返回false
<<<<<<< HEAD
func (tree *Tree) Insert(key int) bool {
	if tree.root == nil {
		tree.root = newNode(key)
		return true
	}

	var parrent = tree.root
=======
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(key)
		return true
	}

	var parrent = tr.root
>>>>>>> tmp
	for {
		if key == parrent.key {
			return false
		}
		if key < parrent.key {
			if parrent.left == nil {
				parrent.left = newNode(key)
				return true
			}
			parrent = parrent.left
		} else {
			if parrent.right == nil {
				parrent.right = newNode(key)
				return true
			}
			parrent = parrent.right
		}
	}
	return true
}

//成功返回true，没有返回false
<<<<<<< HEAD
func (tree *Tree) Remove(key int) bool {
	var target, parrent, lf = tree.root, (*node)(nil), false
=======
func (tr *Tree) Remove(key int32) bool {
	var target, parrent, lf = tr.root, (*node)(nil), false
>>>>>>> tmp
	for target != nil && key != target.key {
		if key < target.key {
			target, parrent, lf = target.left, target, true
		} else {
			target, parrent, lf = target.right, target, false
		}
	}
	if target == nil {
		return false
	}

	var victim, orphan *node = nil, nil
	if target.left == nil {
		victim, orphan = target, target.right
	} else if target.right == nil {
		victim, orphan = target, target.left
	} else { //取中右，取中左也是可以的
		victim, parrent, lf = target.right, target, false
		for victim.left != nil {
			victim, parrent, lf = victim.left, victim, true
		}
		orphan = victim.right
	}

	if parrent == nil { //此时victim==target
<<<<<<< HEAD
		tree.root = orphan
=======
		tr.root = orphan
>>>>>>> tmp
	} else {
		if lf {
			parrent.left = orphan
		} else {
			parrent.right = orphan
		}
		target.key = victim.key //李代桃僵
	}

	return true
}
