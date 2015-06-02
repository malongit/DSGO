package cuckoo

<<<<<<< HEAD
func (table *HashTable) Search(key string) bool {
	return table.first.search(key) || table.second.search(key)
}
func (table *coreTable) search(key string) bool {
	var code = table.hash(key)
	var pt = table.bucket[code%uint(len(table.bucket))]
	return pt != nil && pt.code[table.id] == code && pt.key == key
}

//成功返回true，没有返回false
func (table *HashTable) Remove(key string) bool {
	if table.first.search(key) || table.second.search(key) {
		table.cnt--
		return true
	}
	return false
}
func (table *coreTable) remove(key string) bool {
	var code = table.hash(key)
	var index = code & table.mask
	var pt = table.bucket[index]
	if pt != nil && pt.code[table.id] == code && pt.key == key {
		table.bucket[index] = nil
		return true
=======
func (tb *hashTable) Search(key string) bool {
	for i := 0; i < WAYS; i++ {
		var idx = (tb.idx + i) % WAYS
		var table = &tb.core[idx]
		var code = table.hash(key)
		//var index = code % uint(len(table.bucket))
		var index = code & (uint(len(table.bucket)) - 1)
		var target = table.bucket[index]
		if target != nil &&
			target.code[idx] == code &&
			target.key == key {
			return true
		}
	}
	return false
}

//成功返回true，没有返回false
func (tb *hashTable) Remove(key string) bool {
	for i := 0; i < WAYS; i++ {
		var idx = (tb.idx + i) % WAYS
		var table = &tb.core[idx]
		var code = table.hash(key)
		//var index = code % uint(len(table.bucket))
		var index = code & (uint(len(table.bucket)) - 1)
		var target = table.bucket[index]
		if target != nil &&
			target.code[idx] == code &&
			target.key == key {
			tb.cnt--
			table.bucket[index] = nil
			return true
		}
>>>>>>> tmp
	}
	return false
}

//成功返回true，冲突返回false
<<<<<<< HEAD
func (table *HashTable) Insert(key string) bool {
	if table.Search(key) {
		return false
	}
	var obj = new(node)
	obj.key = key
	obj.code[table.first.id], obj.code[table.second.id] = table.first.hash(key), table.second.hash(key)

	for age := 0; ; age++ {
		var pt = obj
		for { //震荡调整
			var index = obj.code[table.first.id] & table.first.mask
			if table.first.bucket[index] == nil {
				table.first.bucket[index] = pt
				return true
			}
			pt, table.first.bucket[index] = table.first.bucket[index], pt
			if pt == obj {
				break
			}
			index = obj.code[table.second.id] & table.second.mask
			if table.second.bucket[index] == nil {
				table.second.bucket[index] = pt
				return true
			}
			pt, table.second.bucket[index] = table.second.bucket[index], pt
		}

		if age == 2 {
			panic("hash fail!")
		} //实际上不能解决大量hash重码的情况

		//调整失败(回绕)，扩容
		table.first, table.second = table.second, table.first
		var old_bucket = table.first.bucket
		table.first.bucket = make([]*node, len(old_bucket)<<2)
		table.first.mask = (table.first.mask << 2) | 0x3
		for _, unit := range old_bucket {
			var index = unit.code[table.first.id] & table.first.mask
			table.first.bucket[index] = unit //倍扩，绝对不会冲突
=======
func (tb *hashTable) Insert(key string) bool {
	var code [WAYS]uint
	for i := 0; i < WAYS; i++ {
		var table = &tb.core[i]
		code[i] = table.hash(key)
		//var index = code[i] % uint(len(table.bucket))
		var index = code[i] & (uint(len(table.bucket)) - 1)
		var target = table.bucket[index]
		if target != nil &&
			target.code[i] == code[i] &&
			target.key == key {
			return false
		}
	}
	tb.cnt++
	var unit = new(node)
	unit.key = key
	unit.code = code
	for obj, age := unit, 0; ; age++ {
		for { //循环调整
			var table = &tb.core[tb.idx]
			//var index = obj.code[tb.idx] % uint(len(table.bucket))
			var index = obj.code[tb.idx] & (uint(len(table.bucket)) - 1)
			if table.bucket[index] == nil {
				table.bucket[index] = obj
				return true
			}
			obj, table.bucket[index] = table.bucket[index], obj
			if obj == unit {
				break //回绕
			}
			for idx := (tb.idx + 1) % WAYS; idx != tb.idx; idx = (idx + 1) % WAYS {
				table = &tb.core[idx]
				//index = obj.code[idx] % uint(len(table.bucket))
				index = obj.code[idx] & (uint(len(table.bucket)) - 1)
				if table.bucket[index] == nil {
					table.bucket[index] = obj
					return true
				}
				obj, table.bucket[index] = table.bucket[index], obj
			}
		}

		if age != 0 { //这里设定一个阈值，限制扩容次数
			panic("hash fail!")
		} //实际上不能解决大量hash重码的情况，最坏情况只能报错

		//调整失败(回绕)，扩容
		tb.idx = (tb.idx + (WAYS - 1)) % WAYS
		var table = &tb.core[tb.idx]
		var old_bucket = table.bucket
		table.bucket = make([]*node, len(old_bucket)<<WAYS)
		for _, u := range old_bucket {
			if u != nil {
				//var index = u.code[tb.idx] % uint(len(table.bucket))
				var index = u.code[tb.idx] & (uint(len(table.bucket)) - 1)
				table.bucket[index] = u //倍扩，绝对不会冲突
			}
>>>>>>> tmp
		}
	}
	return false
}
