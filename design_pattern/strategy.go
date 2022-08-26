package design_pattern

import (
	"container/list"
	"fmt"
)

type evictionAlgo interface {
	evict()
	resort(int)
}

type FIFO struct {
	cache *Cache
}

func (f *FIFO) evict() {
	fmt.Println("evicted by FIFO")
	elem := f.cache.list.Front()
	f.cache.list.Remove(elem)

	delete(f.cache.storageSet, elem.Value.(int))
}

func (f *FIFO) resort(key int) {
	for e := f.cache.list.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			return
		}
	}
	f.cache.list.PushBack(key)
}

type LRU struct {
	cache *Cache
}

func (l *LRU) evict() {
	fmt.Println("evicted by LRU")
	elem := l.cache.list.Back()
	l.cache.list.Remove(elem)

	delete(l.cache.storageSet, elem.Value.(int))
}

func (l *LRU) resort(key int) {
	for e := l.cache.list.Front(); e != nil; e = e.Next() {
		if e.Value == key {
			l.cache.list.MoveToFront(e)
		}
	}
	l.cache.list.PushFront(key)
}

type Cache struct {
	maxCapacity  int
	storageSet   map[int]string
	evictionAlgo evictionAlgo
	list         *list.List
}

func InitCache(capacity int, strategy string) *Cache {
	cache := &Cache{
		maxCapacity: capacity,
		storageSet:  make(map[int]string),
		list:        list.New(),
	}
	switch strategy {
	case "FIFO":
		cache.evictionAlgo = &FIFO{
			cache: cache,
		}
	case "LRU":
		cache.evictionAlgo = &LRU{
			cache: cache,
		}
	default:
		cache.evictionAlgo = &FIFO{
			cache: cache,
		}
	}
	return cache
}

func (c *Cache) SetEvictionAlgo(eviction evictionAlgo) {
	c.evictionAlgo = eviction
}

func (c *Cache) Put(key int, value string) {
	if len(c.storageSet) >= c.maxCapacity {
		c.evictionAlgo.evict()
	}
	c.storageSet[key] = value
	c.evictionAlgo.resort(key)
}

func (c *Cache) Get(key int) string {
	if item, exist := c.storageSet[key]; exist {
		c.evictionAlgo.resort(key)
		return item
	} else {
		fmt.Println("key not exist")
		return ""
	}
}

func (c *Cache) Print() {
	if len(c.storageSet) == 0 {
		return
	}

	for key, value := range c.storageSet {
		fmt.Printf("key:%d, value:%s \n", key, value)
	}
}
