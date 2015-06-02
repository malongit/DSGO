package avltree

<<<<<<< HEAD
func (parent *node) rotate() (root *node, over bool) {
	over = false
	//root = nil
	if parent.balance == 2 { //左倾右旋
		var child = parent.left
		if child.balance == -1 { //LR
			var grandchild = child.right //一定非nil
			child.right, parent.left = grandchild.left, grandchild.right
			grandchild.left, grandchild.right = child, parent
			switch grandchild.balance {
			case -1:
				parent.balance, child.balance = 0, 1
			case 1:
				parent.balance, child.balance = -1, 0
			default:
				parent.balance, child.balance = 0, 0
			}
			grandchild.balance = 0
			root = grandchild
		} else { //LL
			parent.left, child.right = child.right, parent
			if child.balance == 0 { //不降高旋转
				parent.balance, child.balance = 1, -1
				over = true
			} else { //child.balance == 1
				parent.balance, child.balance = 0, 0
			}
			root = child
		}
	} else { //右倾左旋(parent.balance==-2)
		var child = parent.right
		if child.balance == 1 { //RL
			var grandchild = child.left //一定非nil
			child.left, parent.right = grandchild.right, grandchild.left
			grandchild.right, grandchild.left = child, parent
			switch grandchild.balance {
			case -1:
				parent.balance, child.balance = 1, 0
			case 1:
				parent.balance, child.balance = 0, -1
			default:
				parent.balance, child.balance = 0, 0
			}
			grandchild.balance = 0
			root = grandchild
		} else { //RR
			parent.right, child.left = child.left, parent
			if child.balance == 0 { //不降高旋转
				parent.balance, child.balance = -1, 1
				over = true
			} else { //child.balance == -1
				parent.balance, child.balance = 0, 0
			}
			root = child
		}
	}
	return root, over
=======
//--------------LR形式--------------
//|       G       |       C       |
//|      / \      |      / \      |
//|     P         |     P   G     |
//|    / \        |    / \ / \    |
//|       C       |      u v      |
//|      / \      |               |
//|     u   v     |               |

//--------------LL形式--------------
//|       G       |       P       |
//|      / \      |      / \      |
//|     P         |     C   G     |
//|    / \        |    / \ / \    |
//|   C   x       |        x      |
//|  / \          |               |
//|               |               |

//旋转后高度不发生变化时stop为true
func (G *node) rotate() (root *node, stop bool) {
	stop = false
	//root = nil
	if G.state == 2 { //左倾右旋
		var P = G.left
		if P.state == -1 { //LR
			var C = P.right //一定非nil
			P.right, G.left = P.tryHook(C.left), G.tryHook(C.right)
			C.left, C.right = C.hook(P), C.hook(G)
			switch C.state {
			case 1:
				G.state, P.state = -1, 0
			case -1:
				G.state, P.state = 0, 1
			default:
				G.state, P.state = 0, 0
			}
			C.state = 0
			root = C
		} else { //LL
			G.left, P.right = G.tryHook(P.right), P.hook(G)
			if P.state == 0 { //不降高旋转
				G.state, P.state = 1, -1
				stop = true
			} else { //P.state == 1
				G.state, P.state = 0, 0
			}
			root = P
		}
	} else { //右倾左旋(P.state==-2)
		var P = G.right
		if P.state == 1 { //RL
			var C = P.left //一定非nil
			P.left, G.right = P.tryHook(C.right), G.tryHook(C.left)
			C.right, C.left = C.hook(P), C.hook(G)
			switch C.state {
			case -1:
				G.state, P.state = 1, 0
			case 1:
				G.state, P.state = 0, -1
			default:
				G.state, P.state = 0, 0
			}
			C.state = 0
			root = C
		} else { //RR
			G.right, P.left = G.tryHook(P.left), P.hook(G)
			if P.state == 0 { //不降高旋转
				G.state, P.state = -1, 1
				stop = true
			} else { //P.state == -1
				G.state, P.state = 0, 0
			}
			root = P
		}
	}
	return root, stop
>>>>>>> tmp
}
