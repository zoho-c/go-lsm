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

// 测试迭代器的使用
func TestSkipList_Iterator(t *testing.T) {
	sl := NewSkipList()

	// 插入一些键值对
	sl.Put("key1", []byte("value1"))
	sl.Put("key2", []byte("value2"))
	sl.Put("key3", []byte("value3"))

	// 创建迭代器
	it := sl.Iterator()

	// 遍历并验证顺序和值
	keys := make([]string, 0)
	values := make([][]byte, 0)

	for it.Valid() {
		keys = append(keys, it.Key())
		values = append(values, it.Value())
		it.Next()
	}

	// 期望的结果
	expectedKeys := []string{"key1", "key2", "key3"}
	expectedValues := [][]byte{[]byte("value1"), []byte("value2"), []byte("value3")}

	// 检查长度是否一致
	if len(keys) != len(expectedKeys) {
		t.Errorf("Expected %d keys, got %d", len(expectedKeys), len(keys))
		return
	}

	// 检查每个 key 和 value 是否正确
	for i := 0; i < len(expectedKeys); i++ {
		if keys[i] != expectedKeys[i] {
			t.Errorf("Expected key %s at index %d, got %s", expectedKeys[i], i, keys[i])
		}
		if string(values[i]) != string(expectedValues[i]) {
			t.Errorf("Expected value %s at index %d, got %s", expectedValues[i], i, values[i])
		}
	}
}

// 测试空跳表的迭代器行为
func TestEmptyIterator(t *testing.T) {
	sl := NewSkipList()
	it := sl.Iterator()

	if it.Valid() {
		t.Error("Expected iterator to be invalid on empty skiplist")
	}
}
