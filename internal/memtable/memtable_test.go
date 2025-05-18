package memtable

import "testing"

func TestBasicOperation(t *testing.T) {
	mt := NewMemtable()

	// 插入一对Key Value
	mt.PutWithLock("key1", []byte("value1"))
	if ok, val := mt.GetWithLock("key1"); !ok || string(val) != "value1" {
		t.Fatalf("Get failed, got %v, expect value1", string(val))
	}

	// 插入同key的不同value
	mt.PutWithLock("key1", []byte("value2"))
	if ok, val := mt.GetWithLock("key1"); !ok || string(val) != "value2" {
		t.Fatalf("Get failed, got %v, expect value2", string(val))
	}

	// 空值测试
	if ok, _ := mt.GetWithLock("not_exist"); ok {
		t.Error("Should return false for non-existent key")
	}
	// 注意，这里暂时不测试Remove，需要在实现了在磁盘上的SST上的查找之后再测试
}
