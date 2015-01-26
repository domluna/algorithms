// Implements a hashmap.
//
// The approach presented here is
// called chaining.
package main

import (
	"container/list"
	"fmt"
	"hash/adler32"
)

// entry represents a key/value pair stored
// in the hashMap.
type entry struct {
	key   string
	value interface{}
}

// hashMap represents the hash data structure.
type hashMap struct {
	buckets []*list.List
}

func newMap(buckets uint) *hashMap {
	return &hashMap{
		make([]*list.List, buckets),
	}
}

// resize resizes the hashMap based on the load factor.
// The load factor is the ratio of entries/buckets.
func (hm *hashMap) resize() {
}

func (hm *hashMap) set(key string, v interface{}) {
	i := int(adler32.Checksum([]byte(key))) % len(hm.buckets)

	// get the linked list, insert v
	if hm.buckets[i] == nil {
		hm.buckets[i] = list.New()
	}

	b := hm.buckets[i]
	// see if the key already exists
	for e := b.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*entry)
		if kv.key == key {
			kv.value = v
			return
		}
	}

	// no key is present
	b.PushBack(&entry{
		key:   key,
		value: v,
	})
}

// get returns the value
func (hm *hashMap) get(key string) interface{} {
	i := int(adler32.Checksum([]byte(key))) % len(hm.buckets)

	if hm.buckets[i] == nil {
		return nil
	}

	b := hm.buckets[i]
	for e := b.Front(); e != nil; e = e.Next() {
		kv := e.Value.(*entry)
		if kv.key == key {
			return kv.value
		}
	}
	return nil
}

// String provides a nice string representation for the hashmap.
func (hm *hashMap) String() string {
	var s string
	for i, b := range hm.buckets {
		s += fmt.Sprintf("Bucket %d\n", i)

		// skip empty buckets
		if b == nil {
			continue
		}

		// print out the elements for the bucket
		for e := b.Front(); e != nil; e = e.Next() {
			kv := e.Value.(*entry)
			s += fmt.Sprintf("\tKey: %s, Value: %v\n", kv.key, kv.value)
		}

	}
	return s
}

func main() {
	hm := newMap(5)
	hm.set("hello", 22)
	hm.set("hello1", 22)
	hm.set("hello2", 22)
	hm.set("hello3", 22)
	hm.set("hello4", 22)
	hm.set("hello5", 22)
	hm.set("hello6", 22)
	hm.set("hello world", "hello back")
	hm.set("hello", 35)
	fmt.Println(hm)
	fmt.Println(hm.get("hello world"))
	fmt.Println(hm.get("hello"))
	fmt.Println(hm.get("blank"))
}
