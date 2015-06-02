package rbtree

<<<<<<< HEAD
//红黑树删除过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整
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
//红黑树删除过程包括：O(log N)的搜索，O(1)的旋转，O(log N)的平衡因子调整。
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
		tree.path.push(target, false)
		victim = target.right //取中右
		for victim.left != nil {
			tree.path.push(victim, true)
			victim = victim.left
		}
		orphan = victim.right
	} //如果需要删除的节点有两个儿子，那么问题可以被转化成删除另一个只有一个儿子的节点的问题
	if tree.path.isEmpty() { //此时victim==target
		tree.root = orphan
		if tree.root != nil {
			tree.root.black = true
		}
	} else { //非根，parent!=nil
		tree.hookSubTree(orphan)
		if victim.black { //红victim随便删，黑的要考虑
			if orphan != nil && !orphan.black { //红子变黑顶上
				orphan.black = true
			} else {
				tree.adjustAfterDelete()
			}
		}
		target.key = victim.key //李代桃僵
=======
		victim = target.right
		for victim.left != nil {
			victim = victim.left
		}
		orphan = victim.right
	}

	var root = victim.parent
	if root == nil { //此时victim==target
		tr.root = root.tryHook(orphan)
		if tr.root != nil {
			tr.root.black = true
		}
	} else {
		if key < root.key {
			root.left = root.tryHook(orphan)
		} else {
			root.right = root.tryHook(orphan)
		}
		if victim.black { //红victim随便删，黑的要考虑
			if orphan != nil && !orphan.black {
				orphan.black = true //红子变黑顶上
			} else {
				tr.adjustAfterDelete(root, victim.key)
			}
		}
		target.key = victim.key
>>>>>>> tmp
	}
	return true
}

<<<<<<< HEAD
func (tree *Tree) adjustAfterDelete() {
	var parent, victim_is_left = tree.path.pop()
	for { //剩下情况：victim黑，orphan也黑，此时victim(orphan顶替)的兄弟必然存在
		if victim_is_left {
			var brother = parent.right //brother != nil
			var left_kid, right_kid = brother.left, brother.right
			if brother.black { //Rb，主战场
				if left_kid == nil || left_kid.black {
					if right_kid == nil || right_kid.black { //双黑，变色解决
						brother.black = false
						if parent.black && !tree.path.isEmpty() { //需要上溯
							parent, victim_is_left = tree.path.pop()
							continue
						}
						parent.black = true
					} else { //中黑外红
						parent.right, brother.left = left_kid, parent
						brother.black, parent.black, right_kid.black = parent.black, true, true
						tree.hookSubTree(brother)
					}
				} else { //中红
					brother.left, parent.right = left_kid.right, left_kid.left
					left_kid.right, left_kid.left = brother, parent
					left_kid.black, parent.black = parent.black, true
					tree.hookSubTree(left_kid)
				}
			} else { //Rr，来一次单旋转变出黑brother
				//红brohter下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				parent.right, brother.left = left_kid, parent
				brother.black, parent.black = true, false
				tree.hookSubTree(brother)
				tree.path.push(brother, victim_is_left)
				continue
			}
		} else {
			var brother = parent.left //brother != nil
			var right_kid, left_kid = brother.right, brother.left
			if brother.black { //Rb，主战场
				if right_kid == nil || right_kid.black {
					if left_kid == nil || left_kid.black { //双黑，变色解决
						brother.black = false
						if parent.black && !tree.path.isEmpty() { //需要上溯
							parent, victim_is_left = tree.path.pop()
							continue
						}
						parent.black = true
					} else { //中黑外红
						parent.left, brother.right = right_kid, parent
						brother.black, parent.black, left_kid.black = parent.black, true, true
						tree.hookSubTree(brother)
					}
				} else { //中红
					brother.right, parent.left = right_kid.left, right_kid.right
					right_kid.left, right_kid.right = brother, parent
					right_kid.black, parent.black = parent.black, true
					tree.hookSubTree(right_kid)
				}
			} else { //Rr，来一次单旋转变出黑brother
				//红brohter下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				parent.left, brother.right = right_kid, parent
				brother.black, parent.black = true, false
				tree.hookSubTree(brother)
				tree.path.push(brother, victim_is_left)
				continue
=======
//----------------红叔模式----------------
//|        bG        |        bU        |
//|       /  \       |       /  \       |
//|     bO    rU     |     rG    bR     |
//|          /  \    |    /  \          |
//|        bL    bR  |  bO    bL        |

//------------------双黑------------------
//|        xG        |        bG        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bO    rU     |
//|          /  \    |          /  \    |
//|        bL    bR  |        bL    bR  |

//----------------中黑外红----------------
//|        xG        |        xU        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bR     |
//|          /  \    |    /  \          |
//|        bL    rR  |  bO    bL        |

//------------------中红------------------
//|        xG        |        xL        |
//|       /  \       |       /  \       |
//|     bO    bU     |     bG    bU     |
//|          /  \    |    /  \  /  \    |
//|        rL    xR  |  bO   u  v   xR  |
//|       /  \       |                  |
//|      u    v      |                  |

func (tr *Tree) adjustAfterDelete(G *node, key int32) {
	for { //剩下情况：victim黑，orphan也黑，此时victim(orphan顶替)的兄弟必然存在
		var super = G.parent
		if key < G.key {
			var U = G.right //U != nil
			var L, R = U.left, U.right
			if !U.black { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G.right, U.left = G.hook(L), U.hook(G)
				U.black, G.black = true, false
				tr.hookSubTree(super, U)
				continue //变出黑U后再行解决
			} else {
				if L == nil || L.black {
					if R == nil || R.black { //双黑，变色解决
						U.black = false
						if G.black && super != nil {
							G = super
							continue //上溯
						}
						G.black = true
					} else { //中黑外红
						G.right, U.left = G.tryHook(L), U.hook(G)
						U.black, G.black, R.black = G.black, true, true
						tr.hookSubTree(super, U)
					}
				} else { //中红
					U.left, G.right = U.tryHook(L.right), G.tryHook(L.left)
					L.right, L.left = L.hook(U), L.hook(G)
					L.black, G.black = G.black, true
					tr.hookSubTree(super, L)
				}
			}
		} else {
			var U = G.left //U != nil
			var R, L = U.right, U.left
			if !U.black { //红U下必是两个实体黑，以保证每条支路至少双黑（与victim和orphan也黑双黑匹配）
				G.left, U.right = G.hook(R), U.hook(G)
				U.black, G.black = true, false
				tr.hookSubTree(super, U)
				continue //变出黑U后再行解决
			} else {
				if R == nil || R.black {
					if L == nil || L.black { //双黑，变色解决
						U.black = false
						if G.black && super != nil {
							G = super
							continue //上溯
						}
						G.black = true
					} else { //中黑外红
						G.left, U.right = G.tryHook(R), U.hook(G)
						U.black, G.black, L.black = G.black, true, true
						tr.hookSubTree(super, U)
					}
				} else { //中红
					U.right, G.left = U.tryHook(R.left), G.tryHook(R.right)
					R.left, R.right = R.hook(U), R.hook(G)
					R.black, G.black = G.black, true
					tr.hookSubTree(super, R)
				}
>>>>>>> tmp
			}
		}
		break //个别情况需要循环
	}
}
