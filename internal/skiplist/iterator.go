package skiplist

type Iterator struct {
	curr *Node
}

func (sl *SkipList) Iterator() *Iterator {
	curr := sl.head
	return &Iterator{curr: curr.forward[0]}
}

func (it *Iterator) Next() bool {
	if it.curr == nil || it.curr.forward[0] == nil {
		it.curr = nil
		return false
	}
	it.curr = it.curr.forward[0]
	return true
}

// Key 返回当前节点的键
func (it *Iterator) Key() string {
	if it.curr == nil {
		return ""
	}
	return it.curr.key
}

// Value 返回当前节点的值
func (it *Iterator) Value() []byte {
	if it.curr == nil {
		return nil
	}
	return it.curr.value
}

// Valid 检查当前迭代器是否有效
func (it *Iterator) Valid() bool {
	return it.curr != nil
}
