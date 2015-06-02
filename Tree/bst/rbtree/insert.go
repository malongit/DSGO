package rbtree

<<<<<<< HEAD
func newNode(key int) (unit *node) {
	unit = new(node)
	unit.key, unit.black = key, false
	unit.left, unit.right = nil, nil
	return unit
}

//红黑树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整
//成功返回true，冲突返回false
func (tree *Tree) Insert(key int) bool {
	if tree.root == nil {
		tree.root = newNode(key) //默认为红
		tree.root.black = true
		return true
	}

	tree.path.clear()
	var parent = tree.root
	var kid *node
	var kid_is_left bool
	for {
		if key == parent.key {
			return false
		}
		if key < parent.key {
			kid_is_left = true
			if parent.left == nil {
				kid = newNode(key) //默认为红
				parent.left = kid
				break
			}
			tree.path.push(parent, kid_is_left)
			parent = parent.left
		} else {
			kid_is_left = false
			if parent.right == nil {
				kid = newNode(key) //默认为红
				parent.right = kid
				break
			}
			tree.path.push(parent, kid_is_left)
			parent = parent.right
		}
	}

	for !parent.black { //违法双红禁
		var grandparent, parent_is_left = tree.path.pop() //必然存在，根为黑，parent非根
		if parent_is_left {
			var another = grandparent.right
			if another != nil && !another.black { //红叔模式，变色解决
				parent.black, another.black = true, true
				if !tree.path.isEmpty() {
					grandparent.black = false
					kid = grandparent
					parent, kid_is_left = tree.path.pop() //上溯，检查双红禁
					continue
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if kid_is_left { //LL
					grandparent.left, parent.right = parent.right, grandparent
					grandparent.black, parent.black = false, true
					tree.hookSubTree(parent)
				} else { //LR
					parent.right, grandparent.left = kid.left, kid.right
					kid.left, kid.right = parent, grandparent
					grandparent.black, kid.black = false, true
					tree.hookSubTree(kid)
				}
			}
		} else {
			var another = grandparent.left
			if another != nil && !another.black { //红叔模式，变色解决
				parent.black, another.black = true, true
				if !tree.path.isEmpty() {
					grandparent.black = false
					kid = grandparent
					parent, kid_is_left = tree.path.pop() //上溯，检查双红禁
					continue
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if kid_is_left { //RL
					parent.left, grandparent.right = kid.right, kid.left
					kid.right, kid.left = parent, grandparent
					grandparent.black, kid.black = false, true
					tree.hookSubTree(kid)
				} else { //RR
					grandparent.right, parent.left = parent.left, grandparent
					grandparent.black, parent.black = false, true
					tree.hookSubTree(parent)
=======
//成功返回true，冲突返回false。
//红黑树插入过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
func (tr *Tree) Insert(key int32) bool {
	if tr.root == nil {
		tr.root = newNode(nil, key) //默认为红
		tr.root.black = true
		return true
	}

	var root = tr.root
	for {
		if key < root.key {
			if root.left == nil {
				root.left = newNode(root, key) //默认为红
				break
			}
			root = root.left
		} else if key > root.key {
			if root.right == nil {
				root.right = newNode(root, key) //默认为红
				break
			}
			root = root.right
		} else { //key == root.key
			return false
		}
	}

	//------------红叔模式------------
	//|      bG      |      rG      |
	//|     /  \     |     /  \     |
	//|   rP    rU   |   bP    bU   |
	//|   |          |   |          |
	//|   rC         |   rC         |

	//-----------------LL形式-----------------
	//|        bG        |        bP        |
	//|       /  \       |       /  \       |
	//|     rP    bU     |     rC     rG    |
	//|    /  \          |          /  \    |
	//|  rC    x         |         x    bU  |

	//-----------------LR形式-----------------
	//|        bG        |        bC        |
	//|       /  \       |       /  \       |
	//|     rP    bU     |     rP    rG     |
	//|    / \           |    / \    / \    |
	//|      rC          |       u  v   bU  |
	//|     /  \         |                  |
	//|    u    v        |                  |

	var P = root
	for !P.black { //违法双红禁
		var G = P.parent //必然存在，根为黑，P非根
		var super = G.parent
		if key < G.key {
			var U = G.right
			if U != nil && !U.black { //红叔模式，变色解决
				P.black, U.black = true, true
				if super != nil {
					G.black = false
					P = G.parent
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if key < P.key { //LL
					G.left, P.right = G.tryHook(P.right), P.hook(G)
					G.black, P.black = false, true
					tr.hookSubTree(super, P)
				} else { //LR
					var C = P.right
					P.right, G.left = P.tryHook(C.left), G.tryHook(C.right)
					C.left, C.right = C.hook(P), C.hook(G)
					G.black, C.black = false, true
					tr.hookSubTree(super, C)
				}
			}
		} else {
			var U = G.left
			if U != nil && !U.black { //红叔模式，变色解决
				P.black, U.black = true, true
				if super != nil {
					G.black = false
					P = G.parent
					continue //上溯，检查双红禁
				} //遇根终止
			} else { //黑叔模式，旋转解决
				if key > P.key { //RR
					G.right, P.left = G.tryHook(P.left), P.hook(G)
					G.black, P.black = false, true
					tr.hookSubTree(super, P)
				} else { //RL
					var C = P.left
					P.left, G.right = P.tryHook(C.right), G.tryHook(C.left)
					C.right, C.left = C.hook(P), C.hook(G)
					G.black, C.black = false, true
					tr.hookSubTree(super, C)
>>>>>>> tmp
				}
			}
		}
		break //变色时才需要循环
	}
	return true
}
<<<<<<< HEAD
=======

func newNode(parent *node, key int32) (unit *node) {
	unit = new(node)
	unit.key, unit.black = key, false
	unit.parent = parent
	unit.left, unit.right = nil, nil
	return unit
}
>>>>>>> tmp
