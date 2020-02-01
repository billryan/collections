package set

type Set interface {
	// Adds the specified element to this set if it is not already present (optional operation).
	Add(e interface{})

	// Adds all of the elements to this set if they're not already present (optional operation).
	AddAll(es ...interface{})

	// Removes all of the elements from this set (optional operation).
	Clear()

	// Returns true if this set contains the specified element.
	Contains(e interface{}) bool

	// Returns true if this set contains all of the elements of the specified collection.
	ContainsAll(es ...interface{}) bool

	// Call f for each item in the set
	Foreach(f func(interface{}))

	// Returns true if this set contains no elements.
	IsEmpty() bool

	// Removes the specified element from this set if it is present (optional operation).
	Remove(e interface{}) bool

	// Removes the specified elements from this set if it is present (optional operation).
	// Return true if all element exist.
	RemoveAll(es ...interface{}) bool

	// Return the number of elements in set s (cardinality of s).
	Len() int

	// Returns an slice containing all of the elements in this set.
	ToSlice() []interface{}

	// Return a new set with elements common to the set and all others.
	Intersection(others ...Set) Set

	// Return a new set with elements from the set and all others.
	Union(others ...Set) Set

	// Return a new set with elements in the set that are not in the others.
	Difference(others ...Set) Set

	// Test whether every element in the set is in other. set <= other
	IsSubset(other Set) bool

	// Test whether the set is a proper subset of other, that is, set <= other and set != other.
	IsProperSubset(other Set) bool

	// Test whether every element in other is in the set. set >= other
	IsSuperset(other Set) bool

	// Test whether the set is a proper superset of other, that is, set >= other and set != other.
	IsProperSuperset(other Set) bool
}

// Create a new hash set
func NewHashSet(initial ...interface{}) Set {
	s := &HashSet{make(map[interface{}]nothing)}

	for _, v := range initial {
		s.Add(v)
	}

	return s
}
