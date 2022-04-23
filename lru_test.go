package lru

import "testing"

func TestLRUCachePut(t *testing.T) {
	cache := New(2)
	cache.Put(1, 1)
	if cache.size != 1 {
		t.Errorf("LRU Cache Put failed")
	}
}

func TestLRUCacheGet(t *testing.T) {
	cache := New(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	if cache.size != 2 {
		t.Errorf("LRU Cache Put failed")
	}
	if cache.Get(1) != 1 {
		t.Errorf("LRU Cache Get failed")
	}
}
