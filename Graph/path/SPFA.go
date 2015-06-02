package path

type Path struct {
	Next int
	Dist int
}

<<<<<<< HEAD
const MaxSignedDist = int((^uint(0)) >> 1)

//输入邻接表，返回某点到各点的最短路径的长度及标记是否可达。
//复杂度为O(EV)，逊于Dijkstra，但可以处理负权边。
//出错(例如有负回路)返回空数组。
func SPFA(roads [][]Path, start int) (dists []int, reachable []bool) {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return []int{}, []bool{}
=======
const MAX_DIST = int((^uint(0)) >> 1)

//输入邻接表，返回某点到各点的最短路径的长度(MAX_DIST指不通)。
//实为改良的Bellman-Ford算法，复杂度为O(EV)，逊于Dijkstra，但可以处理负权边。
func SPFA(roads [][]Path, start int) (dists []int, fail bool) {
	var size = len(roads)
	if size == 0 || start < 0 || start >= size {
		return []int{}, true
>>>>>>> tmp
	}

	var ready = make([]bool, size)
	dists = make([]int, size)
	var cnts = make([]int, size)
	for i := 0; i < size; i++ {
<<<<<<< HEAD
		ready[i], dists[i], cnts[i] = true, MaxSignedDist, 0
=======
		ready[i], dists[i], cnts[i] = true, MAX_DIST, 0
>>>>>>> tmp
	}
	var space = make([]int, size)
	var rpt, wpt = 0, 0

	ready[start], dists[start], cnts[start] = false, 0, 1
	space[wpt] = start
	for wpt++; rpt != wpt; rpt = (rpt + 1) % size { //队列非空
		var current = space[rpt]
		ready[current] = true
		for _, path := range roads[current] {
			var distance = dists[current] + path.Dist
			var peer = path.Next
			if distance < dists[peer] {
				dists[peer] = distance
				if ready[peer] { //未入队
					ready[peer] = false
					space[wpt], wpt = peer, (wpt+1)%size
					cnts[peer]++
					if cnts[peer] > size { //负回路
<<<<<<< HEAD
						return []int{}, []bool{}
=======
						return []int{}, true
>>>>>>> tmp
					}
				}
			}
		}
	}
<<<<<<< HEAD
	for i := 0; i < size; i++ {
		if dists[i] != MaxSignedDist {
			ready[i] = true
		} else {
			ready[i], dists[i] = false, 0
		}
	}
	return dists, ready
=======
	return dists, false
>>>>>>> tmp
}
