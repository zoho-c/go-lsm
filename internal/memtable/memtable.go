package memtable

import (
	"github.com/zhouhao/go-lsm/internal/skiplist"
	"sync"
)

type Memtable struct {
	currTable    *skiplist.SkipList   // 活跃的SkipList
	frozenTables []*skiplist.SkipList // 冻结的SkipLists
	rwMtx        sync.RWMutex         // 读写锁
}

// NewMemtable 创建一个新的Memtable
func NewMemtable() *Memtable {
	return &Memtable{
		currTable: skiplist.NewSkipList(),
	}
}

// PutWithLock 向Memtable中存入一个key value 对，并保证并发安全
func (m *Memtable) PutWithLock(key string, value []byte) {
	m.rwMtx.Lock()
	defer m.rwMtx.Unlock()
	m.currTable.Put(key, value)
}

// PutNoLock 向Memtable中存入一个key value
func (m *Memtable) PutNoLock(key string, value []byte) {
	m.currTable.Put(key, value)
	// 判断SkipList的大小
}

//// currGet 从当前活跃的跳表中获取查找key
//func (m *Memtable) currGet(key string) *skiplist.Iterator {
//
//}

// GetWithLock 从Memtable中读取一个key value对，并保证并发安全
func (m *Memtable) GetWithLock(key string) (bool, []byte) {
	m.rwMtx.RLock()
	defer m.rwMtx.RUnlock()
	if ok, value := m.currTable.Get(key); ok {
		return true, value
	}
	// 如果当前的活跃跳表没有找到，则开始从冻结的跳表中寻找
	for _, fzTable := range m.frozenTables {
		// TODO: 思考能不能在锁上进行优化
		if ok, value := fzTable.Get(key); ok {
			return true, value
		}
	}
	// 如果内存中的跳表都没有找到，开始从磁盘上的SST开始找
	return false, nil
}

// GetNoLock 从Memtable中读取一个key value对
func (m *Memtable) GetNoLock(key string) (bool, []byte) {
	if ok, value := m.currTable.Get(key); ok {
		return true, value
	}
	// 如果当前的活跃跳表没有找到，则开始从冻结的跳表中寻找
	for _, fzTable := range m.frozenTables {
		if ok, value := fzTable.Get(key); ok {
			return true, value
		}
	}
	// 如果内存中的跳表都没有找到，开始从磁盘上的SST开始找
	return false, nil
}

// RemoveWithLock 从Memtable中删除一个key
func (m *Memtable) RemoveWithLock(key string) {
	m.rwMtx.Lock()
	defer m.rwMtx.Unlock()
	m.currTable.Put(key, make([]byte, 0))
}

// 实现对于Memtable的迭代器
