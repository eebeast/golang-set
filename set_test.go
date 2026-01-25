package set

import (
	"testing"
)

func TestNew(t *testing.T) {
	t.Run("creates empty set", func(t *testing.T) {
		s := New[string]()
		if s == nil {
			t.Error("New() returned nil")
		}
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
	})

	t.Run("creates empty set for int", func(t *testing.T) {
		s := New[int]()
		if s == nil {
			t.Error("New() returned nil")
		}
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
	})
}

func TestNewWith(t *testing.T) {
	t.Run("creates set with string values", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
		if !s.Has("a") || !s.Has("b") || !s.Has("c") {
			t.Error("Set should contain all initial values")
		}
	})

	t.Run("creates set with int values", func(t *testing.T) {
		s := NewWith(1, 2, 3, 4, 5)
		if s.Len() != 5 {
			t.Errorf("Expected length 5, got %d", s.Len())
		}
		if !s.Has(1) || !s.Has(2) || !s.Has(3) || !s.Has(4) || !s.Has(5) {
			t.Error("Set should contain all initial values")
		}
	})

	t.Run("creates set with duplicate values", func(t *testing.T) {
		s := NewWith("a", "b", "a", "c", "b")
		if s.Len() != 3 {
			t.Errorf("Expected length 3 after deduplication, got %d", s.Len())
		}
		if !s.Has("a") || !s.Has("b") || !s.Has("c") {
			t.Error("Set should contain unique values only")
		}
	})

	t.Run("creates empty set with no arguments", func(t *testing.T) {
		s := NewWith[string]()
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
	})

	t.Run("creates set with single value", func(t *testing.T) {
		s := NewWith("single")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		if !s.Has("single") {
			t.Error("Set should contain the single value")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("adds single element", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		if !s.Has("test") {
			t.Error("Set should contain added element")
		}
	})

	t.Run("adds multiple elements", func(t *testing.T) {
		s := New[string]()
		s.Add("a")
		s.Add("b")
		s.Add("c")
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
		if !s.Has("a") || !s.Has("b") || !s.Has("c") {
			t.Error("Set should contain all added elements")
		}
	})

	t.Run("adding duplicate element doesn't increase length", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		s.Add("test")
		if s.Len() != 1 {
			t.Errorf("Expected length 1 after duplicate add, got %d", s.Len())
		}
		if !s.Has("test") {
			t.Error("Set should still contain the element")
		}
	})

	t.Run("adds int elements", func(t *testing.T) {
		s := New[int]()
		s.Add(1)
		s.Add(2)
		s.Add(3)
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
		if !s.Has(1) || !s.Has(2) || !s.Has(3) {
			t.Error("Set should contain all added int elements")
		}
	})

	t.Run("adds empty string", func(t *testing.T) {
		s := New[string]()
		s.Add("")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		if !s.Has("") {
			t.Error("Set should contain empty string")
		}
	})
}

func TestRemove(t *testing.T) {
	t.Run("removes existing element", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		s.Remove("b")
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
		if s.Has("b") {
			t.Error("Set should not contain removed element")
		}
		if !s.Has("a") || !s.Has("c") {
			t.Error("Set should still contain other elements")
		}
	})

	t.Run("removes non-existent element (no error)", func(t *testing.T) {
		s := NewWith("a", "b")
		s.Remove("c")
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
		if !s.Has("a") || !s.Has("b") {
			t.Error("Set should still contain original elements")
		}
	})

	t.Run("removes all elements", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		s.Remove("a")
		s.Remove("b")
		s.Remove("c")
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
		if s.Has("a") || s.Has("b") || s.Has("c") {
			t.Error("Set should not contain any elements")
		}
	})

	t.Run("removes from empty set (no error)", func(t *testing.T) {
		s := New[string]()
		s.Remove("test")
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
	})

	t.Run("removes int element", func(t *testing.T) {
		s := NewWith(1, 2, 3)
		s.Remove(2)
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
		if s.Has(2) {
			t.Error("Set should not contain removed int element")
		}
	})
}

func TestHas(t *testing.T) {
	t.Run("returns true for existing element", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		if !s.Has("a") {
			t.Error("Has() should return true for existing element")
		}
	})

	t.Run("returns false for non-existent element", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		if s.Has("d") {
			t.Error("Has() should return false for non-existent element")
		}
	})

	t.Run("returns false for empty set", func(t *testing.T) {
		s := New[string]()
		if s.Has("anything") {
			t.Error("Has() should return false for empty set")
		}
	})

	t.Run("returns true after add", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		if !s.Has("test") {
			t.Error("Has() should return true after Add()")
		}
	})

	t.Run("returns false after remove", func(t *testing.T) {
		s := NewWith("a", "b")
		s.Remove("a")
		if s.Has("a") {
			t.Error("Has() should return false after Remove()")
		}
	})

	t.Run("works with int type", func(t *testing.T) {
		s := NewWith(1, 2, 3)
		if !s.Has(2) {
			t.Error("Has() should work with int type")
		}
		if s.Has(99) {
			t.Error("Has() should return false for non-existent int")
		}
	})
}

func TestLen(t *testing.T) {
	t.Run("returns 0 for empty set", func(t *testing.T) {
		s := New[string]()
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
	})

	t.Run("returns correct length after adds", func(t *testing.T) {
		s := New[string]()
		s.Add("a")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		s.Add("b")
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
		s.Add("c")
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
	})

	t.Run("returns correct length after removes", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
		s.Remove("b")
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
		s.Remove("a")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		s.Remove("c")
		if s.Len() != 0 {
			t.Errorf("Expected length 0, got %d", s.Len())
		}
	})

	t.Run("length doesn't change with duplicate adds", func(t *testing.T) {
		s := New[string]()
		s.Add("test")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		s.Add("test")
		if s.Len() != 1 {
			t.Errorf("Expected length 1 after duplicate add, got %d", s.Len())
		}
	})

	t.Run("works with NewWith", func(t *testing.T) {
		s := NewWith("a", "b", "c", "d", "e")
		if s.Len() != 5 {
			t.Errorf("Expected length 5, got %d", s.Len())
		}
	})
}

func TestSetOperations(t *testing.T) {
	t.Run("add remove add cycle", func(t *testing.T) {
		s := New[string]()
		s.Add("a")
		s.Remove("a")
		s.Add("a")
		if s.Len() != 1 {
			t.Errorf("Expected length 1, got %d", s.Len())
		}
		if !s.Has("a") {
			t.Error("Set should contain element after add-remove-add cycle")
		}
	})

	t.Run("multiple operations sequence", func(t *testing.T) {
		s := New[string]()
		s.Add("a")
		s.Add("b")
		s.Add("c")
		s.Remove("b")
		s.Add("d")
		s.Remove("a")
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
		if !s.Has("c") || !s.Has("d") {
			t.Error("Set should contain correct elements after sequence")
		}
		if s.Has("a") || s.Has("b") {
			t.Error("Set should not contain removed elements")
		}
	})
}

func TestSetWithDifferentTypes(t *testing.T) {
	t.Run("float64 set", func(t *testing.T) {
		s := NewWith(1.1, 2.2, 3.3)
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
		if !s.Has(2.2) {
			t.Error("Set should work with float64")
		}
		s.Remove(2.2)
		if s.Len() != 2 {
			t.Errorf("Expected length 2, got %d", s.Len())
		}
	})

	t.Run("rune set", func(t *testing.T) {
		s := NewWith('a', 'b', 'c')
		if s.Len() != 3 {
			t.Errorf("Expected length 3, got %d", s.Len())
		}
		if !s.Has('b') {
			t.Error("Set should work with rune")
		}
	})
}

func TestToSlice(t *testing.T) {
	t.Run("returns slice of string values", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		slice := s.ToSlice()
		if len(slice) != 3 {
			t.Errorf("Expected slice length 3, got %d", len(slice))
		}
		// Check that all elements are present (order may vary)
		elements := make(map[string]bool)
		for _, v := range slice {
			elements[v] = true
		}
		if !elements["a"] || !elements["b"] || !elements["c"] {
			t.Error("Slice should contain all set elements")
		}
	})

	t.Run("returns empty slice for empty set", func(t *testing.T) {
		s := New[string]()
		slice := s.ToSlice()
		if len(slice) != 0 {
			t.Errorf("Expected empty slice, got length %d", len(slice))
		}
		if slice == nil {
			t.Error("ToSlice() should return empty slice, not nil")
		}
	})

	t.Run("returns slice of int values", func(t *testing.T) {
		s := NewWith(1, 2, 3, 4, 5)
		slice := s.ToSlice()
		if len(slice) != 5 {
			t.Errorf("Expected slice length 5, got %d", len(slice))
		}
		elements := make(map[int]bool)
		for _, v := range slice {
			elements[v] = true
		}
		for i := 1; i <= 5; i++ {
			if !elements[i] {
				t.Errorf("Slice should contain element %d", i)
			}
		}
	})

	t.Run("returns slice with single element", func(t *testing.T) {
		s := NewWith("single")
		slice := s.ToSlice()
		if len(slice) != 1 {
			t.Errorf("Expected slice length 1, got %d", len(slice))
		}
		if slice[0] != "single" {
			t.Errorf("Expected 'single', got %q", slice[0])
		}
	})

	t.Run("slice length matches set length", func(t *testing.T) {
		s := NewWith("a", "b", "c", "d", "e", "f")
		slice := s.ToSlice()
		if len(slice) != s.Len() {
			t.Errorf("Slice length %d should match set length %d", len(slice), s.Len())
		}
	})

	t.Run("returns slice after modifications", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		s.Remove("b")
		s.Add("d")
		slice := s.ToSlice()
		if len(slice) != 3 {
			t.Errorf("Expected slice length 3, got %d", len(slice))
		}
		elements := make(map[string]bool)
		for _, v := range slice {
			elements[v] = true
		}
		if !elements["a"] || !elements["c"] || !elements["d"] {
			t.Error("Slice should contain current set elements")
		}
		if elements["b"] {
			t.Error("Slice should not contain removed element")
		}
	})

	t.Run("works with float64 type", func(t *testing.T) {
		s := NewWith(1.1, 2.2, 3.3)
		slice := s.ToSlice()
		if len(slice) != 3 {
			t.Errorf("Expected slice length 3, got %d", len(slice))
		}
	})

	t.Run("returns new slice each time", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		slice1 := s.ToSlice()
		slice2 := s.ToSlice()
		if &slice1[0] == &slice2[0] {
			t.Error("ToSlice() should return a new slice each time")
		}
	})
}

func TestForeach(t *testing.T) {
	t.Run("calls function for each element", func(t *testing.T) {
		s := NewWith("a", "b", "c")
		visited := make(map[string]bool)
		s.Foreach(func(value string) {
			visited[value] = true
		})
		if len(visited) != 3 {
			t.Errorf("Expected 3 elements visited, got %d", len(visited))
		}
		if !visited["a"] || !visited["b"] || !visited["c"] {
			t.Error("Foreach should visit all elements")
		}
	})

	t.Run("doesn't call function for empty set", func(t *testing.T) {
		s := New[string]()
		called := false
		s.Foreach(func(value string) {
			called = true
		})
		if called {
			t.Error("Foreach should not call function for empty set")
		}
	})

	t.Run("works with int type", func(t *testing.T) {
		s := NewWith(1, 2, 3, 4, 5)
		sum := 0
		s.Foreach(func(value int) {
			sum += value
		})
		expectedSum := 1 + 2 + 3 + 4 + 5
		if sum != expectedSum {
			t.Errorf("Expected sum %d, got %d", expectedSum, sum)
		}
	})

	t.Run("works with single element", func(t *testing.T) {
		s := NewWith("single")
		called := false
		s.Foreach(func(value string) {
			if value != "single" {
				t.Errorf("Expected 'single', got %q", value)
			}
			called = true
		})
		if !called {
			t.Error("Foreach should call function for single element")
		}
	})

	t.Run("works with float64 type", func(t *testing.T) {
		s := NewWith(1.1, 2.2, 3.3)
		var values []float64
		s.Foreach(func(value float64) {
			values = append(values, value)
		})
		if len(values) != 3 {
			t.Errorf("Expected 3 values, got %d", len(values))
		}
	})
}
