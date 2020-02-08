package set

import "encoding/json"

type (
	HashSet struct {
		hash map[interface{}]nothing
	}

	nothing struct{}
)

// Adds the specified element to this set if it is not already present (optional operation).
func (s *HashSet) Add(e interface{}) {
	s.hash[e] = nothing{}
}

// Adds all of the elements to this set if they're not already present (optional operation).
func (s *HashSet) AddAll(es ...interface{}) {
	for _, e := range es {
		s.hash[e] = nothing{}
	}
}

// Removes all of the elements from this set (optional operation).
func (s *HashSet) Clear() {
	for e := range s.hash {
		delete(s.hash, e)
	}
}

// Returns true if this set contains the specified element.
func (s *HashSet) Contains(e interface{}) bool {
	_, exist := s.hash[e]
	return exist
}

// Returns true if this set contains all of the elements of the specified collection.
func (s *HashSet) ContainsAll(es ...interface{}) bool {
	for _, e := range es {
		_, exist := s.hash[e]
		if !exist {
			return false
		}
	}
	return true
}

// Call f for each item in the set
func (s *HashSet) Foreach(f func(interface{})) {
	for k := range s.hash {
		f(k)
	}
}

// Call f for each item in the set, set result as new key
func (s *HashSet) Map(f func(interface{}) interface{}) Set {
	n := make(map[interface{}]nothing)
	for k := range s.hash {
		n[f(k)] = nothing{}
	}
	return &HashSet{n}
}

// Returns true if this set contains no elements.
func (s *HashSet) IsEmpty() bool {
	return len(s.hash) == 0
}

// Removes the specified element from this set if it is present (optional operation).
func (s *HashSet) Remove(e interface{}) bool {
	_, exist := s.hash[e]
	delete(s.hash, e)
	return exist
}

// Removes the specified elements from this set if it is present (optional operation).
// Return true if all element exist.
func (s *HashSet) RemoveAll(es ...interface{}) bool {
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
func (s *HashSet) Len() uint32 {
	return uint32(len(s.hash))
}

// Returns an slice containing all of the elements in this set.
func (s *HashSet) ToSlice() []interface{} {
	slice := make([]interface{}, 0)
	for e := range s.hash {
		slice = append(slice, e)
	}
	return slice
}

func (s *HashSet) Clone() Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		n[k] = nothing{}
	}

	return &HashSet{n}
}

// Return a new set with elements common to the set and all others.
func (s *HashSet) Intersection(others ...Set) Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		existAll := true
		for _, set := range others {
			if !set.Contains(k) {
				existAll = false
				break
			}
		}
		if existAll {
			n[k] = nothing{}
		}
	}

	return &HashSet{n}
}

// Return a new set with elements from the set and all others.
func (s *HashSet) Union(others ...Set) Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		n[k] = nothing{}
	}
	for _, set := range others {
		hashset := set.(*HashSet)
		for k := range hashset.hash {
			n[k] = nothing{}
		}
	}

	return &HashSet{n}
}

// Return a new set with elements in the set that are not in the others.
func (s *HashSet) Difference(others ...Set) Set {
	n := make(map[interface{}]nothing)

	for k := range s.hash {
		existAny := false
		for _, set := range others {
			if set.Contains(k) {
				existAny = true
			}
		}
		if !existAny {
			n[k] = nothing{}
		}
	}

	return &HashSet{n}
}

// Test whether every element in the set is in other. set <= other
func (s *HashSet) IsSubset(other Set) bool {
	if s.Len() > other.Len() {
		return false
	}
	for k := range s.hash {
		if !other.Contains(k) {
			return false
		}
	}
	return true
}

// Test whether the set is a proper subset of other, that is, set <= other and set != other.
func (s *HashSet) IsProperSubset(other Set) bool {
	return s.Len() < other.Len() && s.IsSubset(other)
}

// Test whether every element in other is in the set. set >= other
func (s *HashSet) IsSuperset(other Set) bool {
	return other.IsSubset(s)
}

// Test whether the set is a proper superset of other, that is, set >= other and set != other.
func (s *HashSet) IsProperSuperset(other Set) bool {
	return s.Len() > other.Len() && s.IsSuperset(other)
}

func (s *HashSet) UnmarshalText(text []byte) error {
	var v []interface{}
	err := json.Unmarshal(text, &v)
	if err == nil {
		s.hash = make(map[interface{}]nothing)
		for _, k := range v {
			s.hash[k] = nothing{}
		}
	}
	return err
}
