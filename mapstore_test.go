package mapstore

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
)

func BenchmarkRead50Write50(b *testing.B) {
	keys := make([]string, 10000)
	for i := 0; i < len(keys); i++ {
		keys[i] = fmt.Sprintf("%d", i)
	}
	store := NewWithSize(4)
	wg := sync.WaitGroup{}
	b.ResetTimer()
	threads := 4
	wg.Add(threads * 2)
	for i := 0; i < threads; i++ {
		go testWrites(store, keys, b.N, &wg)
		go testReads(store, keys, b.N, &wg)
	}
	wg.Wait()
	b.Logf("stat: %v", store.ShardStats())
}

func testWrites(s *Store, keys []string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	lenKeys := len(keys)
	for i := 0; i < num; i++ {
		key := keys[rand.Int()%lenKeys]
		s.Set(key, key)
	}
}

func testReads(s *Store, keys []string, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	lenKeys := len(keys)
	for i := 0; i < num; i++ {
		key := keys[rand.Int()%lenKeys]
		s.Get(key, key)
	}
}
