package skiplist

import (
	"strconv"
	"testing"
)

// 测试基础操作
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

	// 删除操作
	sl.Remove("key1")
	if ok, _ := sl.Get("key1"); ok {
		t.Error("Should return false for non-existent key")
	}
	sl.Remove("key11") // 删除不存在的key，不应该出现报错
}

// 测试大规模插入与读取操作
func TestLargeScaleInsertAndGet(t *testing.T) {
	sl := NewSkipList()
	numElements := 10000

	for i := 0; i < numElements; i++ {
		key := "key" + strconv.Itoa(i)
		value := []byte("value" + strconv.Itoa(i))
		sl.Put(key, value)
	}

	for i := 0; i < numElements; i++ {
		key := "key" + strconv.Itoa(i)
		value := []byte("value" + strconv.Itoa(i))
		if ok, val := sl.Get(key); !ok || string(val) != string(value) {
			t.Fatalf("Get failed, got %v, expect %v", string(val), string(value))
		}
	}
}

// 测试大规模删除操作
func TestLargeScaleDelete(t *testing.T) {
	sl := NewSkipList()
	numElements := 10000

	for i := 0; i < numElements; i++ {
		key := "key" + strconv.Itoa(i)
		value := []byte("value" + strconv.Itoa(i))
		sl.Put(key, value)
	}

	for i := 0; i < numElements; i++ {
		key := "key" + strconv.Itoa(i)
		sl.Remove(key)
	}

	for i := 0; i < numElements; i++ {
		key := "key" + strconv.Itoa(i)
		if ok, val := sl.Get(key); ok {
			t.Fatalf("Should not get value, but get %v", string(val))
		}
	}
}
