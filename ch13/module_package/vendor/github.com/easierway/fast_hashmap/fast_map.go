package fast_hashmap

import (
	"container/list"
)

func hashCode(str string) int {
	h := 0
	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		h = h*31 + int(bytes[i])
	}
	return h
}

func indexFor(h int, capacity int) int {
	return h & (capacity - 1)
}

type MapWithStringKey interface {
	Put(key string, v interface{})
	Get(key string) (interface{}, bool)
}

type Entry struct {
	Key      string
	Value    interface{}
	lenOfKey int
}

type BucketWithSlice struct {
	entries []*Entry
}

type BucketWithList struct {
	entryList *list.List
}

type FasterMap struct {
	buckets         []*BucketWithSlice
	length          int
	capacity        int
	numOfCollisions int
}

type FastMap struct {
	buckets  []*BucketWithList
	length   int
	capacity int
}

func calTableCapacity(length int) int {
	capacity := 1
	for capacity < length*3 {
		capacity <<= 1
	}
	return capacity
}

func NewFasterMap(length int) *FasterMap {
	capacity := calTableCapacity(length)
	buckets := make([]*BucketWithSlice, capacity, capacity)
	return &FasterMap{
		buckets,
		length,
		capacity,
		0,
	}
}

func (m *FasterMap) Put(key string, v interface{}) {
	isFirstBucket := false
	hCode := hashCode(key)
	pos := indexFor(hCode, m.capacity)
	bucket := m.buckets[pos]
	if bucket == nil {
		entries := make([]*Entry, 0, 10)
		m.buckets[pos] = &BucketWithSlice{entries}
		isFirstBucket = true
	}
	for _, entry := range m.buckets[pos].entries {
		if entry.Key == key {
			entry.Value = v
			return
		}
	}
	m.buckets[pos].entries = append(m.buckets[pos].entries, &Entry{key, v, len(key)})
	if !isFirstBucket {
		m.numOfCollisions += 1
	}
}

func (m *FasterMap) Get(key string) (interface{}, bool) {
	hCode := hashCode(key)
	pos := indexFor(hCode, m.capacity)
	bucket := m.buckets[pos]
	if bucket == nil {
		return nil, false
	}
	if len(bucket.entries) == 1 {
		return bucket.entries[0].Value, true
	}
	var lenOfKey int
	for _, entry := range bucket.entries {
		lenOfKey = len(key)
		if lenOfKey == entry.lenOfKey && entry.Key == key {
			return entry.Value, true
		}

	}
	panic("Fatal for hash bucket")
}

func NewFastMap(length int) *FastMap {
	capacity := calTableCapacity(length)
	buckets := make([]*BucketWithList, capacity, capacity)
	return &FastMap{
		buckets,
		length,
		capacity,
	}
}

func (m *FastMap) Put(key string, v interface{}) {
	hCode := hashCode(key)
	pos := indexFor(hCode, m.capacity)
	bucket := m.buckets[pos]
	if bucket == nil {
		m.buckets[pos] = &BucketWithList{list.New()}
	}
	for i := m.buckets[pos].entryList.Front(); i != nil; i = i.Next() {
		entry, ok := (i.Value).(*Entry)
		if !ok {
			panic("Wrong entry")
		}
		if entry.Key == key {
			entry.Value = v
			return
		}
	}
	m.buckets[pos].entryList.PushBack(&Entry{key, v, len(key)})
}

func (m *FastMap) Get(key string) (interface{}, bool) {
	hCode := hashCode(key)
	pos := indexFor(hCode, m.capacity)
	bucket := m.buckets[pos]
	if bucket == nil {
		return nil, false
	}

	for i := bucket.entryList.Front(); i != nil; i = i.Next() {
		entry, ok := (i.Value).(*Entry)
		if !ok {
			panic("Wrong entry")
		}
		if entry.Key == key {
			return entry.Value, true
		}
	}
	panic("Fatal for hash bucket")
}
