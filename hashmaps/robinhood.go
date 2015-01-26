// Implements a hashmap with the
// Robinhood scheme.
package main

// Entry represents an entry in the hashmap.
type Entry struct {
	key   string
	value interface{}
}

// Map represents hashmap using the
// Robinhood scheme.
type Map struct {
	buckets []*Entry
}

// Insert inserts the k/v pair into the hashmap.
func (m *Map) Insert(key string, v interface{}) {

}

// Delete deletes the k/v pair associated with the key
// from the hashmap.
func (m *Map) Delete(key string) {

}

// Get attempts to find the k/v pair associated with
// the key. If found returns the value if not, returns nil.
func (m *Map) Get(key string) interface{} {
	return nil

}

func main() {

}
