package skiplist

import (
	"math/rand"
	"time"
)

const (
	maxLevel = 16
)

type Node struct {
	key     string
	value   []byte
	forward []*Node
}

type SkipList struct {
	head       *Node
	level      int
	randSource *rand.Rand
	// TODO: 实现对于内存使用量的追踪
	// TODO: 添加读写锁机制实现对于SkipList的并发访问
}

// NewSkipList 创建一个新的跳表
func NewSkipList() *SkipList {
	head := &Node{
		forward: make([]*Node, maxLevel),
	}

	return &SkipList{
		head:       head,
		level:      1,
		randSource: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// randomLevel 返回一个随机的层数
func (sl *SkipList) randomLevel() int {
	lvl := 1
	for (sl.randSource.Int()&1) == 1 && lvl < maxLevel {
		lvl++
	}
	return lvl
}

// Put 向SkipList中插入一个键值对
func (sl *SkipList) Put(key string, value []byte) {
	nodesUpdates := make([]*Node, maxLevel)
	curr := sl.head

	for i := sl.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].key < key {
			curr = curr.forward[i]
		}
		nodesUpdates[i] = curr
	}

	curr = curr.forward[0]

	if curr != nil && curr.key == key {
		curr.value = value
		return
	}

	newLevel := sl.randomLevel()
	if newLevel > sl.level {
		for i := sl.level; i < newLevel; i++ {
			nodesUpdates[i] = sl.head // 对于新的Node，使用head作为前驱节点
		}
		sl.level = newLevel
	}

	newNode := &Node{
		key:     key,
		value:   value,
		forward: make([]*Node, maxLevel),
	}

	for i := 0; i < sl.level; i++ {
		newNode.forward[i] = nodesUpdates[i].forward[i]
		nodesUpdates[i].forward[i] = newNode
	}

}

// Get 从SkipList中取出一个key对应的value
func (sl *SkipList) Get(key string) (bool, []byte) {
	curr := sl.head
	for i := sl.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].key < key {
			curr = curr.forward[i]
		}
	}

	curr = curr.forward[0]
	if curr != nil && curr.key == key {
		return true, curr.value
	}
	return false, nil
}

// Remove 从SkipList中按照key删除一个键值对
func (sl *SkipList) Remove(key string) {
	nodesUpdates := make([]*Node, maxLevel)
	curr := sl.head
	// 查找插入位置
	for i := sl.level - 1; i >= 0; i-- {
		for curr.forward[i] != nil && curr.forward[i].key < key {
			curr = curr.forward[i]
		}
		nodesUpdates[i] = curr
	}

	curr = curr.forward[0]

	// 不存在这个key
	if curr == nil || curr.key != key {
		return
	}

	for i := 0; i < sl.level; i++ {
		if nodesUpdates[i].forward[i] != curr {
			break
		}
		nodesUpdates[i].forward[i] = curr.forward[i]
	}
	// 更新当前有效层数（当高层为空时降低层数）
	for sl.level > 1 && sl.head.forward[sl.level-1] == nil {
		sl.level--
	}
}

// TODO: 实现前缀搜索与范围搜索
