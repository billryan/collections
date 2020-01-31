package set

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Error("Length of empty init set should be 0")
	}

	s = New(1, 4, 8)
	if s.Len() != 3 {
		t.Error("Length should be 3")
	}
}

func TestSet_Add(t *testing.T) {
	s := New()
	s.Add("k1")

	if !s.Contains("k1") {
		t.Error("Set should contain 'k1'")
	}
}

func TestSet_AddAll(t *testing.T) {
	s := New()
	s.AddAll("k1", "k2")
	if s.Len() != 2 {
		t.Error("Length should be 2")
	}
}

func TestSet_Clear(t *testing.T) {
	s := New(1, 2, 4, 8)
	s.Clear()
	if s.Len() != 0 {
		t.Error("Length should be 0 after clear()")
	}
}

func TestSet_Contains(t *testing.T) {
	s := New(1, 2)
	if s.Contains(0) {
		t.Error("Set should not contain 0")
	}

	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
}

func TestSet_ContainsAll(t *testing.T) {
	s := New(1, 2, 4)
	if !s.ContainsAll(1, 2) {
		t.Error("Set should contain 1 and 2")
	}

	if s.ContainsAll(1, 3) {
		t.Error("Set should not contain 1 and 3")
	}
}

func TestSet_Difference(t *testing.T) {
	s1 := New(1, 2, 4)
	s2 := New(4, 8)
	s3 := s1.Difference(s2)
	if s3.Contains(4) || s3.Contains(8) {
		t.Error("Set should not contain 4 or 8")
	}

	if !s3.ContainsAll(1, 2) {
		t.Error("Set should contain 1 and 2")
	}
}

func TestSet_Foreach(t *testing.T) {
	s := New(1, 2, 3)
	f := func(x interface{}) { fmt.Printf("type of x is %T and value is %v\n", x, x) }
	s.Foreach(f)
}

func TestSet_Len(t *testing.T) {
	s := New()
	if s.Len() != 0 {
		t.Error("Length should be 0")
	}

	s.AddAll(1, 2, 4)
	if s.Len() != 3 {
		t.Error("Length should be 3")
	}
}

func TestSet_Remove(t *testing.T) {
	s := New()
	s.Remove(2)
	if s.Len() != 0 {
		t.Error("Length should be 0")
	}

	s.AddAll(1, 2, 4)
	s.Remove(2)
	if s.Contains(2) {
		t.Error("Set s should not contain 2")
	}
}

func TestSet_RemoveAll(t *testing.T) {
	s := New(1, 2, 4)
	exist := s.RemoveAll(1, 2)
	if !exist {
		t.Error("Set s should contain 1 and 2 before RemoveAll(1, 2)")
	}

	if !s.Contains(4) {
		t.Error("Set s should contain 4")
	}
}

func Test(t *testing.T) {
	s := New()

	s.Add(5)

	if s.Len() != 1 {
		t.Errorf("Length should be 1")
	}

	s.Remove(5)

	if s.Len() != 0 {
		t.Errorf("Length should be 0")
	}

	// Difference
	s1 := New(1, 2, 3, 4, 5, 6)
	s2 := New(4, 5, 6)
	s3 := s1.Difference(s2)

	if s3.Len() != 3 {
		t.Errorf("Length should be 3")
	}

	// Intersection
	s3 = s1.Intersection(s2)
	if s3.Len() != 3 {
		t.Errorf("Length should be 3 after intersection")
	}

	// Union
	s4 := New(7, 8, 9)
	s3 = s2.Union(s4)

	if s3.Len() != 6 {
		t.Errorf("Length should be 6 after union")
	}

	// Subset
	if !s1.SubsetOf(s1) {
		t.Errorf("set should be a subset of itself")
	}
	// Proper Subset
	if s1.ProperSubsetOf(s1) {
		t.Errorf("set should not be a subset of itself")
	}

}
