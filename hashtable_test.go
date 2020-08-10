package go_data_structures

import (
	"fmt"
	"testing"
)

func populateHashTable(count int, start int) *ValueHashTable {
	dict := ValueHashTable{}
	for i := start; i < (start + count); i++ {
		dict.Put(fmt.Sprintf("key%d", i), fmt.Sprintf("value%d", i))
	}
	return &dict
}

func TestValueHashTable_Put(t *testing.T) {
	dict := populateHashTable(3, 0)
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	dict.Put("key1", "value1")
	if size := dict.Size(); size != 3 {
		t.Errorf("wrong count, expected 3 and got %d", size)
	}
	dict.Put("key4", "value4")
	if size := dict.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}

func TestValueHashTable_Get(t *testing.T) {
	dict := populateHashTable(3, 0)
	if value := dict.Get("key1"); value != "value1" {
		t.Errorf("wrong value, expected value1 and got %s", value)
	}
}

func TestValueHashTable_Remove(t *testing.T) {
	dict := populateHashTable(5, 0)
	dict.Remove("key1")
	if size := dict.Size(); size != 4 {
		t.Errorf("wrong count, expected 4 and got %d", size)
	}
}

func TestValueHashTable_IsEmpty(t *testing.T) {
	dict := populateHashTable(1, 0)
	if isEmpty := dict.IsEmpty(); isEmpty {
		t.Errorf("the dict is empty")
	}
	dict.Remove("key0")
	if isEmpty := dict.IsEmpty(); !isEmpty {
		t.Errorf("the dict has items")
	}
}

func TestValueHashTable_Keys(t *testing.T) {
	dict := populateHashTable(3, 0)
	if keysLen := len(dict.Keys()); keysLen != 3 {
		t.Errorf("wrong count, expected 3 and got %d", keysLen)
	}
}
