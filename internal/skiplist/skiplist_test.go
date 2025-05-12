package skiplist

import (
	"testing"
)

func TestBasicOperation(t *testing.T) {
	sl := NewSkipList()
	// 基础的插入与查找操作
	sl.Put("key1", []byte("value1"))
	if ok, val := sl.Get("key1"); !ok || string(val) != "value1" {
		t.Fatalf("Get failed, got %v, expect value1", string(val))
	}

	// 插入同key的不同value
	sl.Put("key1", []byte("value2"))
	if ok, val := sl.Get("key1"); !ok || string(val) != "value2" {
		t.Fatalf("Get failed, got %v, expect value2", string(val))
	}

	// 空值测试
	if ok, _ := sl.Get("not_exist"); ok {
		t.Error("Should return false for non-existent key")
	}
}

// TODO: 测试删除操作
// func TestRemove(t *testing.T) {
// }
