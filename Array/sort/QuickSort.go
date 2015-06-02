package sort

//快速排序，改进的冒泡排序，不具有稳定性。
//平均复杂度为O(NlogN) & O(logN)，最坏情况是O(N^2) & O(N)。
//其中比较操作是O(NlogN)，常数与MergeSort相当；挪移操作是O(NlogN)，常数小于MergeSort。
<<<<<<< HEAD
=======
//QuickSort不适合递归实现(有爆栈风险)。
>>>>>>> tmp
func QuickSort(list []int) {
	var tasks stack
	tasks.push(0, len(list))
	for !tasks.isEmpty() {
		var start, end = tasks.pop()
		if end-start < sz_limit {
			InsertSort(list[start:end])
		} else {
<<<<<<< HEAD
			var knot = part(list[start:end]) + start
=======
			var knot = partition(list[start:end]) + start
>>>>>>> tmp
			tasks.push(knot+1, end)
			tasks.push(start, knot)
		} //每轮保证至少解决一个，否则最坏情况可能是死循环
	}
}

<<<<<<< HEAD
func part(list []int) int {
=======
func partition(list []int) int {
>>>>>>> tmp
	var size = len(list)

	//三点取中法，最后保证seed落入中间，每轮至少解决此一处
	var seed = list[0]
	var mid, last = size / 2, size - 1
	if list[0] > list[mid] {
		if list[mid] > list[last] {
			seed, list[mid] = list[mid], list[0]
		} else { //c >= b
			if list[0] > list[last] {
				seed, list[last] = list[last], list[0]
			}
		}
	} else { //b >= a
		if list[last] > list[0] {
			if list[mid] > list[last] {
				seed, list[last] = list[last], list[0]
			} else {
				seed, list[mid] = list[mid], list[0]
			}
		}
	}

	var left, right = 1, last
	for { //注意对称性
		for list[left] < seed {
			left++
		}
		for list[right] > seed {
			right--
		}
		if left >= right {
			break
		}
		list[left], list[right] = list[right], list[left]
		left++
		right--
	}
	list[0], list[right] = list[right], seed

	return right
}
<<<<<<< HEAD
=======

type pair struct {
	start int
	end   int
}
type stack struct {
	core []pair
}

func (s *stack) size() int {
	return len(s.core)
}
func (s *stack) isEmpty() bool {
	return len(s.core) == 0
}
func (s *stack) push(start int, end int) {
	s.core = append(s.core, pair{start, end})
}
func (s *stack) pop() (start int, end int) {
	var sz = len(s.core) - 1
	var unit = s.core[sz]
	s.core = s.core[:sz]
	return unit.start, unit.end
}
>>>>>>> tmp
