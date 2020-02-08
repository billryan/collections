package set

import (
	"encoding/json"
	"sync"
	"sync/atomic"
)

type (
	ConcurrentSet struct {
		hash sync.Map
		size uint32
	}
)

func (s *ConcurrentSet) Add(e interface{}) {
	_, exists := s.hash.LoadOrStore(e, nothing{})
	if !exists {
		atomic.AddUint32(&s.size, 1)
	}
}

func (s *ConcurrentSet) AddAll(es ...interface{}) {
	for _, e := range es {
		_, exists := s.hash.LoadOrStore(e, nothing{})
		if !exists {
			atomic.AddUint32(&s.size, 1)
		}
	}
}

func (s *ConcurrentSet) Clear() {
	s.hash.Range(func(k, v interface{}) bool {
		s.hash.Delete(k)
		return true
	})
	atomic.StoreUint32(&s.size, 0)
}

func (s *ConcurrentSet) Contains(e interface{}) bool {
	_, exists := s.hash.Load(e)
	return exists
}

func (s *ConcurrentSet) ContainsAll(es ...interface{}) bool {
	for _, e := range es {
		_, exists := s.hash.Load(e)
		if !exists {
			return false
		}
	}
	return true
}

func (s *ConcurrentSet) Foreach(f func(interface{})) {
	s.hash.Range(func(k, v interface{}) bool {
		f(k)
		return true
	})
}

func (s *ConcurrentSet) Map(f func(interface{}) interface{}) Set {
	sizeMap := make(map[interface{}]nothing)
	var n sync.Map
	s.hash.Range(func(k, v interface{}) bool {
		newK := f(k)
		n.Store(newK, nothing{})
		sizeMap[newK] = nothing{}
		return true
	})

	return &ConcurrentSet{n, uint32(len(sizeMap))}
}

// Returns true if this set contains no elements.
func (s *ConcurrentSet) IsEmpty() bool {
	return s.size == 0
}

// Removes the specified element from this set if it is present (optional operation).
func (s *ConcurrentSet) Remove(e interface{}) bool {
	_, exists := s.hash.Load(e)
	s.hash.Delete(e)
	if exists {
		atomic.AddUint32(&s.size, ^uint32(0))
	}
	return exists
}

// Removes the specified elements from this set if it is present (optional operation).
// Return true if all element exist.
func (s *ConcurrentSet) RemoveAll(es ...interface{}) bool {
	existAll := true
	for _, e := range es {
		_, exists := s.hash.Load(e)
		if exists {
			s.hash.Delete(e)
			atomic.AddUint32(&s.size, ^uint32(0))
		} else {
			existAll = false
		}
	}
	return existAll
}

// Return the number of elements in set s (cardinality of s).
func (s *ConcurrentSet) Len() uint32 {
	return atomic.LoadUint32(&s.size)
}

// Returns an slice containing all of the elements in this set.
func (s *ConcurrentSet) ToSlice() []interface{} {
	slice := make([]interface{}, 0)
	s.hash.Range(func(k, v interface{}) bool {
		slice = append(slice, k)
		return true
	})

	return slice
}

func (s *ConcurrentSet) Clone() Set {
	var n sync.Map
	s.hash.Range(func(k, v interface{}) bool {
		n.Store(k, nothing{})
		return true
	})

	return &ConcurrentSet{n, s.size}
}

// Return a new set with elements common to the set and all others.
func (s *ConcurrentSet) Intersection(others ...Set) Set {
	var n sync.Map

	size := uint32(0)
	s.hash.Range(func(k, v interface{}) bool {
		existAll := true
		for _, set := range others {
			if !set.Contains(k) {
				existAll = false
				break
			}
		}
		if existAll {
			n.Store(k, nothing{})
			size++
		}
		return true
	})

	return &ConcurrentSet{n, size}
}

// Return a new set with elements from the set and all others.
func (s *ConcurrentSet) Union(others ...Set) Set {
	var n sync.Map
	size := s.size
	s.hash.Range(func(k, v interface{}) bool {
		n.Store(k, nothing{})
		return true
	})

	for _, set := range others {
		cset := set.(*ConcurrentSet)
		cset.hash.Range(func(k, v interface{}) bool {
			_, exists := n.LoadOrStore(k, nothing{})
			if exists {
				size++
			}
			return true
		})
	}

	return &ConcurrentSet{n, size}
}

// Return a new set with elements in the set that are not in the others.
func (s *ConcurrentSet) Difference(others ...Set) Set {
	var n sync.Map
	size := uint32(0)

	s.hash.Range(func(k, v interface{}) bool {
		existAny := false
		for _, set := range others {
			if set.Contains(k) {
				existAny = true
			}
		}
		if !existAny {
			n.Store(k, nothing{})
			size++
		}
		return true
	})

	return &ConcurrentSet{n, size}
}

// Test whether every element in the set is in other. set <= other
func (s *ConcurrentSet) IsSubset(other Set) bool {
	if s.Len() > other.Len() {
		return false
	}

	isSubset := true
	s.hash.Range(func(k, v interface{}) bool {
		if !other.Contains(k) {
			isSubset = false
			return false
		}
		return true
	})
	return isSubset
}

// Test whether the set is a proper subset of other, that is, set <= other and set != other.
func (s *ConcurrentSet) IsProperSubset(other Set) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

// Test whether every element in other is in the set. set >= other
func (s *ConcurrentSet) IsSuperset(other Set) bool {
	return other.IsSubset(s)
}

// Test whether the set is a proper superset of other, that is, set >= other and set != other.
func (s *ConcurrentSet) IsProperSuperset(other Set) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}

func (s *ConcurrentSet) UnmarshalText(text []byte) error {
	var v []interface{}
	err := json.Unmarshal(text, &v)
	if err == nil {
		for _, k := range v {
			s.hash.Store(k, nothing{})
		}
		s.size = uint32(len(v))
	}
	return err
}

func (s *ConcurrentSet) ToSet() *HashSet {
	n := make(map[interface{}]nothing)

	s.hash.Range(func(k, v interface{}) bool {
		n[k] = nothing{}
		return true
	})

	return &HashSet{n}
}