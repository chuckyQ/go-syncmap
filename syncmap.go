package syncmap

import (
	"sync"
)

type SyncMap[K comparable, V any] struct {
	m            *sync.Map
	defaultValue V
}

// Store sets the value for a key.
//
// Added in go1.9
func (s *SyncMap[K, V]) Store(key K, value V) {
	s.m.Store(key, value)
}

// Load returns the value stored in the map for a key, or nil if no value is present. The ok result indicates whether value was found in the map.
//
// Added in go1.9
func (s *SyncMap[K, V]) Load(key K) (V, bool) {

	obj, exists := s.m.Load(key)
	if exists {
		return obj.(V), exists
	}

	return s.defaultValue, exists

}

// Delete deletes the value for a key. If the key is not in the map, Delete does nothing.
//
// Added in go1.9
func (s *SyncMap[K, V]) Delete(key K) {
	s.m.Delete(key)

}

// Clear deletes all the entries, resulting in an empty Map.
//
// Added in go1.23
func (s *SyncMap[K, V]) Clear() {
	s.m.Clear()
}

// Swap swaps the value for a key and returns the previous value if any. The loaded result reports whether the key was present.
//
// Added in go1.20
func (s *SyncMap[K, V]) Swap(key K, value V) (V, bool) {
	previous, exists := s.m.Swap(key, value)
	if exists {
		return previous.(V), exists
	}
	return s.defaultValue, exists
}

// CompareAndDelete deletes the entry for key if its value is equal to old. The old value must be of a comparable type.
//
// If there is no current value for key in the map, CompareAndDelete returns false (even if the old value is the nil interface value).
//
// Added in go1.20
func (s *SyncMap[K, V]) CompareAndDelete(key K, value V) bool {
	return s.m.CompareAndDelete(key, value)
}

// CompareAndSwap swaps the old and new values for key if the value stored in the map is equal to old. The old value must be of a comparable type.
//
// Added in go1.20
func (s *SyncMap[K, V]) CompareAndSwap(key K, previous V, new V) bool {
	return s.m.CompareAndSwap(key, previous, new)
}

// LoadOrStore returns the existing value for the key if present. Otherwise, it stores and returns the given value. The loaded result is true if the value was loaded, false if stored.
//
// Added in go1.9
func (s *SyncMap[K, V]) LoadOrStore(key K, value V) (V, bool) {
	actual, loaded := s.m.LoadOrStore(key, value)
	if loaded {
		return actual.(V), loaded
	}
	return value, loaded
}

// LoadAndDelete deletes the value for a key, returning the previous value if any. The loaded result reports whether the key was present.
//
// Added in go1.15
func (s *SyncMap[K, V]) LoadAndDelete(key K) (V, bool) {

	value, exists := s.m.LoadAndDelete(key)
	if exists {
		return value.(V), exists
	}
	return s.defaultValue, exists

}

// Range calls f sequentially for each key and value present in the map. If f returns false, range stops the iteration.
//
// Range does not necessarily correspond to any consistent snapshot of the Map's contents: no key will be visited more than once, but if the value for any key is stored or deleted concurrently (including by f), Range may reflect any mapping for that key from any point during the Range call. Range does not block other methods on the receiver; even f itself may call any method on m.
//
// Range may be O(N) with the number of elements in the map even if f returns false after a constant number of calls.
//
// Added in go1.9
func (s *SyncMap[K, V]) Range(f func(key K, value V) bool) {
	s.m.Range(func(key, value any) bool {
		return f(key.(K), value.(V))
	})
}

func New[K comparable, V any](defaultValue V) *SyncMap[K, V] {
	return &SyncMap[K, V]{
		m:            &sync.Map{},
		defaultValue: defaultValue,
	}

}
