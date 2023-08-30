package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type CacheItem struct {
	val interface{}
	key Key
}

func (lc *lruCache) Get(key Key) (interface{}, bool) {
	value, exists := lc.items[key]
	if exists == false {
		return nil, false
	}
	lc.queue.MoveToFront(value)
	item := value.Value.(CacheItem)
	return item.val, true
}

func (lc *lruCache) Clear() {
	lc.queue = NewList()
	lc.items = make(map[Key]*ListItem, lc.capacity)
}

func (lc *lruCache) Set(key Key, value interface{}) bool {
	item := CacheItem{
		key: key,
		val: value,
	}
	v, exists := lc.items[key]
	if exists == true {
		v.Value = item
		lc.queue.MoveToFront(v)
		return true
	} else {
		lc.items[key] = lc.queue.PushFront(item)
		for lc.queue.Len() > lc.capacity {
			di := lc.queue.Back()
			lc.queue.Remove(di)
			dv := di.Value.(CacheItem)
			delete(lc.items, dv.key)
		}
		return false
	}

}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
