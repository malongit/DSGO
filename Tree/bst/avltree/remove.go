package avltree

<<<<<<< HEAD
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整
//成功返回true，没有返回false
func (tree *Tree) Remove(key int) bool {
	tree.path.clear()
	var target = tree.root
	for target != nil && key != target.key {
		if key < target.key {
			tree.path.push(target, true)
			target = target.left
		} else {
			tree.path.push(target, false)
=======
//成功返回true，没有返回false。
//AVL树删除过程包括：O(log N)的搜索，O(log N)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Remove(key int32) bool {
	var target = tr.root
	for target != nil && key != target.key {
		if key < target.key {
			target = target.left
		} else {
>>>>>>> tmp
			target = target.right
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
	} else {
<<<<<<< HEAD
		if target.balance == 1 {
			tree.path.push(target, true)
			victim = target.left //取中左
			for victim.right != nil {
				tree.path.push(victim, false)
=======
		if target.state == 1 {
			victim = target.left
			for victim.right != nil {
>>>>>>> tmp
				victim = victim.right
			}
			orphan = victim.left
		} else {
<<<<<<< HEAD
			tree.path.push(target, false)
			victim = target.right //取中右
			for victim.left != nil {
				tree.path.push(victim, true)
=======
			victim = target.right
			for victim.left != nil {
>>>>>>> tmp
				victim = victim.left
			}
			orphan = victim.right
		}
<<<<<<< HEAD
	} //如果需要删除的节点有两个儿子，那么问题可以被转化成删除另一个只有一个儿子的节点的问题

	if tree.path.isEmpty() { //此时victim==target
		tree.root = orphan
	} else {
		var parent, victim_is_left = tree.path.pop() //非根，parent!=nil
		var tips = parent.balance
		if victim_is_left {
			parent.left = orphan
			parent.balance-- //左支减短
		} else {
			parent.right = orphan
			parent.balance++ //右支减短
		}

		for tips != 0 { //如果原balance为0则子树高度不变
			if parent.balance == 0 { //无需旋转，super无需重接子树
				if tree.path.isEmpty() {
					break
				}
				parent, victim_is_left = tree.path.pop()
				tips = parent.balance
				if victim_is_left {
					parent.balance-- //左支减短
				} else {
					parent.balance++ //右支减短
				}
				continue
			}
			var subtree, keep_height = parent.rotate()
			if tree.path.isEmpty() { //到根
				tree.root = subtree
				break
			}
			var super, sub_is_left = tree.path.pop()
			if keep_height {
				if sub_is_left {
					super.left = subtree
				} else {
					super.right = subtree
				}
				break
			}
			//降高旋转
			tips = super.balance
			if sub_is_left {
				super.left = subtree
				super.balance-- //左支减短
			} else {
				super.right = subtree
				super.balance++ //右支减短
			}
			parent = super
		}
		target.key = victim.key //李代桃僵
	}
=======
	}

	var root = victim.parent
	if root == nil { //此时victim==target
		tr.root = root.tryHook(orphan)
		return true
	}
	key = victim.key

	var state, stop = root.state, false
	if key < root.key {
		root.left = root.tryHook(orphan)
		root.state--
	} else {
		root.right = root.tryHook(orphan)
		root.state++
	}

	for state != 0 { //如果原平衡因子为0则子树高度不变
		var super = root.parent
		if super == nil {
			if root.state != 0 { //2 || -2
				root, _ = root.rotate()
				tr.root = super.hook(root)
			}
			break
		} else {
			if root.state != 0 { //2 || -2
				root, stop = root.rotate()
				if key < super.key {
					super.left = super.hook(root)
				} else {
					super.right = super.hook(root)
				}
				if stop {
					break
				}
			}
			root, state = super, super.state
			if key < root.key {
				root.state--
			} else {
				root.state++
			}
		}
	}
	target.key = key
>>>>>>> tmp
	return true
}
