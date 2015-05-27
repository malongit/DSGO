package avlt

//成功返回true，冲突返回false。
//AVL树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tree *Tree) Insert(key int32) bool {
	if tree.root == nil {
		tree.root = newNode(key)
		return true
	}
	tree.path.clear()
	var root = tree.root
	for {
		if key < root.key {
			tree.path.push(root, true)
			if root.left == nil {
				root.left = newNode(key)
				break
			}
			root = root.left
		} else if key > root.key {
			tree.path.push(root, false)
			if root.right == nil {
				root.right = newNode(key)
				break
			}
			root = root.right
		} else { //key == root.key
			return false
		}
	}

	var state, lf = int8(0), false
	for !tree.path.isEmpty() && state == 0 {
		root, lf = tree.path.pop()
		state = root.adjust(!lf)
	}
	if state != 0 && root.state != 0 { //2 || -2
		root, _ = root.rotate()
		if tree.path.isEmpty() {
			tree.root = root
		} else {
			tree.hookSubTree(root)
		}
	}
	return true
}

func newNode(key int32) (unit *node) {
	unit = new(node)
	unit.key, unit.state = key, 0
	unit.left, unit.right = nil, nil
	return unit
}
