package syncmap_test

import (
	"sync"
	"testing"

	"github.com/chuckyQ/go-syncmap"
)

func TestMap(t *testing.T) {

	m := sync.Map{}
	s := syncmap.New[int, int](0)

	num, exists := m.LoadOrStore(10, 20)
	num2, exists2 := s.LoadOrStore(10, 20)
	if num.(int) != num2 || exists != exists2 {
		t.Fatalf("test failed expected %v, got %v", num, num2)
	}

	num, previousValue := m.LoadAndDelete(10)
	num2, previousValue2 := s.LoadAndDelete(10)
	if num.(int) != num2 || previousValue != previousValue2 {
		t.Fatalf("test failed expected %v, got %v", num, num2)
	}

	m.Store(10, 20)
	s.Store(10, 20)
	num, exists = m.Load(10)
	num2, exists2 = s.Load(10)
	if num.(int) != num2 || exists != exists2 {
		t.Fatalf("test failed expected %v, got %v", num, num2)
	}

	deleted := m.CompareAndDelete(10, 20)
	deleted2 := s.CompareAndDelete(10, 20)
	if deleted != deleted2 {
		t.Fatalf("test failed expected %v, got %v", deleted, deleted2)
	}

	swapped := m.CompareAndSwap(10, 20, 30)
	swapped2 := s.CompareAndSwap(10, 20, 30)
	if swapped != swapped2 {
		t.Fatalf("test failed expected %v, got %v", deleted, deleted2)
	}

}
