package hashmap

import (
	"testing"
)

func TestHash_Add(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", "value1")
	if !hash.Exists("key1") {
		t.Error("Expected key1 to exist in hash after adding, but it doesn't.")
	}
}

func TestHash_Remove(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", "value1")
	hash.Remove("key1")
	if hash.Exists("key1") {
		t.Error("Expected key1 to be removed from hash, but it still exists.")
	}
}

func TestHash_Exists(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", "value1")
	if !hash.Exists("key1") {
		t.Error("Expected key1 to exist in hash, but it doesn't.")
	}
}

func TestHash_Get(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", "value1")
	value := hash.Get("key1")
	if value != "value1" {
		t.Errorf("Expected value 'value1', but got '%v'", value)
	}
}

func TestHash_Length(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", "value1")
	hash.Add("key2", "value2")
	length := hash.Length()
	if length != 2 {
		t.Errorf("Expected length of 2, but got %d", length)
	}
}

func TestHash_Copy(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", "value1")
	copy := hash.Copy()
	if len(copy) != 1 || copy["key1"] != "value1" {
		t.Errorf("Copy is not correct: %v", copy)
	}
}

func TestHash_Map(t *testing.T) {
	hash := NewHash(make(map[string]interface{}))
	hash.Add("key1", 5)
	hash.Add("key2", 10)
	var sum int
	hash.Map(func(record interface{}) bool {
		val := record.(int)
		sum += val
		return false
	})
	if sum != 15 {
		t.Errorf("Expected sum of 15, but got %d", sum)
	}
}
