package go_data_structures

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Key generic.Type

type Value generic.Type

type ValueHashTable struct {
	items map[int]Value
	lock  sync.RWMutex
}

func hash(k Key) int {
	key := fmt.Sprintf("%s", k)
	h := 0
	for i := 0; i < len(key); i++ {
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *ValueHashTable) Remove(k Key) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	delete(ht.items, i)
}

func (ht *ValueHashTable) Get(k Key) Value {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	i := hash(k)
	return ht.items[i]
}

func (ht *ValueHashTable) Size() int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.items)
}

func (ht *ValueHashTable) Put(k Key, v Value) {
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	if ht.items == nil {
		ht.items = make(map[int]Value)
	}
	ht.items[i] = v
}

func (ht *ValueHashTable) IsEmpty() bool {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	itemsCount := len(ht.items)
	return itemsCount == 0
}

func (ht *ValueHashTable) Keys() []int {
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	keys := make([]int, 0, len(ht.items))
	for k := range ht.items {
		keys = append(keys, k)
	}
	return keys
}
