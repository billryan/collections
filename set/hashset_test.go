package set

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	s := NewHashSet()
	if s.Len() != 0 {
		t.Error("Length of empty init set should be 0")
	}

	s = NewHashSet(1, 4, 8)
	if s.Len() != 3 {
		t.Error("Length should be 3")
	}
}

func TestHashSet_Add(t *testing.T) {
	s := NewHashSet()
	s.Add("k1")

	if !s.Contains("k1") {
		t.Error("Set should contain 'k1'")
	}
}

func TestHashSet_AddAll(t *testing.T) {
	s := NewHashSet()
	s.AddAll("k1", "k2")
	if s.Len() != 2 {
		t.Error("Length should be 2")
	}
}

func TestHashSet_Clear(t *testing.T) {
	s := NewHashSet(1, 2, 4, 8)
	s.Clear()
	if s.Len() != 0 {
		t.Error("Length should be 0 after clear()")
	}
}

func TestHashSet_Contains(t *testing.T) {
	s := NewHashSet(1, 2)
	if s.Contains(0) {
		t.Error("Set should not contain 0")
	}

	if !s.Contains(1) {
		t.Error("Set should contain 1")
	}
}

func TestHashSet_ContainsAll(t *testing.T) {
	s := NewHashSet(1, 2, 4)
	if !s.ContainsAll(1, 2) {
		t.Error("Set should contain 1 and 2")
	}

	if s.ContainsAll(1, 3) {
		t.Error("Set should not contain 1 and 3")
	}
}

func TestHashSet_Foreach(t *testing.T) {
	s := NewHashSet(1, 2, 3)
	f := func(x interface{}) { fmt.Printf("type of x is %T and value is %v\n", x, x) }
	s.Foreach(f)
}

func TestHashSet_Map(t *testing.T) {
	oldSet := NewHashSet(1, 2, 3)
	f := func(x int) int { return x * x }
	fWrapper := func(x interface{}) interface{} {
		xInt := x.(int)
		return f(xInt)
	}
	s := oldSet.Map(fWrapper)
	if s.Len() != oldSet.Len() {
		t.Errorf("Length should be %d", oldSet.Len())
	}
	if !s.ContainsAll(1, 4, 9) {
		t.Error("Set should be 1, 4, 9")
	}
}

func TestHashSet_Len(t *testing.T) {
	s := NewHashSet()
	if s.Len() != 0 {
		t.Error("Length should be 0")
	}

	s.AddAll(1, 2, 2)
	if s.Len() != 2 {
		t.Error("Length should be 2")
	}
}

func TestHashSet_Clone(t *testing.T) {
	s := NewHashSet(1, 2, 4)
	s2 := s.Clone()
	s.Remove(1)
	if !s2.Contains(1) {
		t.Error("Set s2 should contain 1")
	}
	if !s2.ContainsAll(1, 2, 4) {
		t.Error("Set s2 should contain 1, 2, 4")
	}
}

func TestHashSet_IsEmpty(t *testing.T) {
	s := NewHashSet()
	if !s.IsEmpty() {
		t.Error("Set should be empty")
	}
	s.Add(1)
	if s.IsEmpty() {
		t.Error("Set should not be empty")
	}
}

func TestHashSet_Remove(t *testing.T) {
	s := NewHashSet()
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

func TestHashSet_RemoveAll(t *testing.T) {
	s := NewHashSet(1, 2, 4)
	exist := s.RemoveAll(1, 2)
	if !exist {
		t.Error("Set s should contain 1 and 2 before RemoveAll(1, 2)")
	}

	if !s.Contains(4) {
		t.Error("Set s should contain 4")
	}
}

func TestHashSet_Intersection(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(1, 2, 8)
	s3 := NewHashSet(1, 2, 9, 12)
	s := s1.Intersection(s2, s3)
	if !s.ContainsAll(1, 2) {
		t.Error("Set should contain 1 and 2")
	}
	if s.Contains(4) || s.Contains(8) {
		t.Error("Set should not contain 4 and 8")
	}
}

func TestHashSet_Union(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(1, 2, 8)
	s3 := NewHashSet(2, 3)
	s := s1.Union(s2, s3)
	if !s.ContainsAll(1, 2, 3, 4, 8) {
		t.Error("Set should contain 1, 2, 3, 4, 8")
	}
}

func TestHashSet_Difference(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(4, 8)
	s3 := NewHashSet(4, 9, 12)
	s := s1.Difference(s2, s3)
	if s.Contains(4) || s.Contains(8) {
		t.Error("Set should not contain 4 or 8")
	}

	if !s.ContainsAll(1, 2) {
		t.Error("Set should contain 1 and 2")
	}
}

func TestHashSet_IsSubset(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(2, 4)
	if !s1.IsSubset(s1) {
		t.Error("Set s1 should be subset of s1")
	}
	if !s2.IsSubset(s1) {
		t.Error("Set s2 should be subset of s1")
	}
	if s1.IsSubset(s2) {
		t.Error("Set s1 should not be subset of s2")
	}
}

func TestHashSet_IsProperSubset(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(2, 4)
	if s1.IsProperSubset(s1) {
		t.Error("Set s1 should not be proper subset of s1")
	}
	if !s2.IsProperSubset(s1) {
		t.Error("Set s2 should be proper subset of s1")
	}
	if s1.IsProperSubset(s2) {
		t.Error("Set s1 should not be proper subset of s2")
	}
}

func TestHashSet_IsSuperset(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(2, 4)
	if !s1.IsSuperset(s1) {
		t.Error("Set s1 should be superset of s1")
	}
	if !s1.IsSuperset(s2) {
		t.Error("Set s1 should be superset of s2")
	}
	if s2.IsSuperset(s1) {
		t.Error("Set s2 should not be superset of s1")
	}
}

func TestHashSet_IsProperSuperset(t *testing.T) {
	s1 := NewHashSet(1, 2, 4)
	s2 := NewHashSet(2, 4)
	if s1.IsProperSuperset(s1) {
		t.Error("Set s1 should not be proper superset of s1")
	}
	if !s1.IsProperSuperset(s2) {
		t.Error("Set s1 should be proper superset of s2")
	}
	if s2.IsProperSuperset(s1) {
		t.Error("Set s2 should not be proper superset of s1")
	}
}

func TestHashSet_ToSlice(t *testing.T) {
	s := NewHashSet(1, 2, 4)
	slice := s.ToSlice()
	if len(slice) != 3 {
		t.Error("Slice length should be 3")
	}
	i1Cnt, i2Cnt, i4Cnt := 0, 0, 0
	for _, i := range slice {
		iInt := i.(int)
		if iInt == 1 {
			i1Cnt += 1
		} else if iInt == 2 {
			i2Cnt += 1
		} else if iInt == 4 {
			i4Cnt += 1
		}
	}
	if i1Cnt != 1 || i2Cnt != 1 || i4Cnt != 1 {
		t.Error("slice element should be 1, 2, 4")
	}
}

func TestHashSet_UnmarshalText(t *testing.T) {
	text := []byte(`["billryan", "test"]`)
	s := NewHashSet()
	err := s.UnmarshalText(text)
	if err != nil {
		t.Errorf("error while unmarshal text %s", err)
	}
	if !s.ContainsAll("billryan", "test") {
		t.Error("Set should contain 'billryan' and 'test'")
	}

	//	type Config struct {
	//		Users    HashSet
	//		Password string
	//	}
	//
	//	doc := []byte(`
	//Users = '["billryan", "test"]'
	//Password = "mypassword"`)
	//
	//	config := Config{}
	//	err = toml.Unmarshal(doc, &config)
	//	if err != nil {
	//		t.Errorf("toml unmarshal error %s", err)
	//	}
}
