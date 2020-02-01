package set

type (
	Set struct {
		hash map[interface{}]nothing
	}

	nothing struct{}
)

// Create a new set
func New(initial ...interface{}) *Set {
	s := &Set{make(map[interface{}]nothing)}

	for _, v := range initial {
		s.Add(v)
	}

	return s
}

// Adds the specified element to this set if it is not already present (optional operation).
func (s *Set) Add(e interface{}) {
	s.hash[e] = nothing{}
}

// Adds all of the elements to this set if they're not already present (optional operation).
func (s *Set) AddAll(es ...interface{}) {
	for _, e := range es {
		s.hash[e] = nothing{}
	}
}

// Removes all of the elements from this set (optional operation).
func (s *Set) Clear() {
	for e := range s.hash {
		delete(s.hash, e)
	}
}

// Returns true if this set contains the specified element.
func (s *Set) Contains(e interface{}) bool {
	_, exist := s.hash[e]
	return exist
}

// Returns true if this set contains all of the elements of the specified collection.
func (s *Set) ContainsAll(es ...interface{}) bool {
	for _, e := range es {
		_, exist := s.hash[e]
		if !exist {
			return false
		}
	}
	return true
}

// Call f for each item in the set
func (s *Set) Foreach(f func(interface{})) {
	for k := range s.hash {
		f(k)
	}
}

// Returns true if this set contains no elements.
func (s *Set) IsEmpty() bool {
	return len(s.hash) == 0
}

// Removes the specified element from this set if it is present (optional operation).
func (s *Set) Remove(e interface{}) bool {
	_, exist := s.hash[e]
	delete(s.hash, e)
	return exist
}

// Removes the specified elements from this set if it is present (optional operation).
// Return true if all element exist.
func (s *Set) RemoveAll(es ...interface{}) bool {
	existAll := true
	for _, e := range es {
		_, exist := s.hash[e]
		if exist {
			delete(s.hash, e)
		} else {
			existAll = false
		}
	}
	return existAll
}

// Return the number of elements in set s (cardinality of s).
func (s *Set) Len() int {
	return len(s.hash)
}

// Returns an slice containing all of the elements in this set.
func (s *Set) ToSlice() []interface{} {
	slice := make([]interface{}, s.Len())
	for e := range s.hash {
		slice = append(slice, e)
	}
	return slice
}

// Return a new set with elements common to the set and all others.
func (s *Set) Intersection(others ...*Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		existAll := true
		for _, set := range others {
			if _, exist := set.hash[k]; !exist {
				existAll = false
				break
			}
		}
		if existAll {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

// Return a new set with elements from the set and all others.
func (s *Set) Union(others ...*Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		n[k] = nothing{}
	}
	for _, set := range others {
		for k := range set.hash {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

// Return a new set with elements in the set that are not in the others.
func (s *Set) Difference(others ...*Set) *Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		existAny := false
		for _, set := range others {
			if _, exist := set.hash[k]; exist {
				existAny = true
			}
		}
		if !existAny {
			n[k] = nothing{}
		}
	}

	return &Set{n}
}

func (s *Set) isSubset(other *Set, reverse bool) bool {
	src, dest := s, other
	if reverse {
		src, dest = dest, src
	}
	if src.Len() > dest.Len() {
		return false
	}
	for k := range src.hash {
		if _, exists := dest.hash[k]; !exists {
			return false
		}
	}
	return true
}

// Test whether every element in the set is in other. set <= other
func (s *Set) IsSubset(other *Set) bool {
	return s.isSubset(other, false)
}

// Test whether the set is a proper subset of other, that is, set <= other and set != other.
func (s *Set) IsProperSubset(other *Set) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

// Test whether every element in other is in the set. set >= other
func (s *Set) IsSuperset(other *Set) bool {
	return s.isSubset(other, true)
}

// Test whether the set is a proper superset of other, that is, set >= other and set != other.
func (s *Set) IsProperSuperset(other *Set) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}
