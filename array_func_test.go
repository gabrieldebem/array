package array_test

import (
	"testing"

	"github.com/gabrieldebem/array"
)

func TestFilterWithInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	odd := array.Filter[int](numbers, func(value int) bool {
		return value % 2 == 0
	})

	if len(odd) != 4 {
		t.Errorf("Expected length of 4, got %d", len(odd))
	}
}

func TestFilterWithString(t *testing.T) {
	words := []string{"gopher", "monkey", "robot", "ninja"}
	filtered := array.Filter[string](words, func(value string) bool {
		return len(value) > 5
	})

	if len(filtered) != 2 {
		t.Errorf("Expected length of 2, got %d", len(filtered))
	}
}

func TestMapWithInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	mapped := array.Map[int, int](numbers, func(value int) int {
		return value * 2
	})

	if len(mapped) != 9 {
		t.Errorf("Expected length of 9, got %d", len(mapped))
	}

	for i, v := range mapped {
		if v != numbers[i]*2 {
			t.Errorf("Expected %d, got %d", numbers[i]*2, v)
		}
	}
}

func TestMapWithString(t *testing.T) {
	words := []string{"gopher", "monkey", "robot", "ninja"}
	mapped := array.Map[string, string](words, func(value string) string {
		return value + "!"
	})

	if len(mapped) != 4 {
		t.Errorf("Expected length of 4, got %d", len(mapped))
	}

	for i, v := range mapped {
		if v != words[i]+"!" {
			t.Errorf("Expected %s, got %s", words[i]+"!", v)
		}
	}
}

func TestReduceWithInt(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	reduced := array.Reduce[int](numbers, func(carry, value int) int {
		return carry + value
	})

	if reduced != 45 {
		t.Errorf("Expected 45, got %d", reduced)
	}
}

func TestReduceWithString(t *testing.T) {
	words := []string{"gopher", "monkey", "robot", "ninja"}
	reduced := array.Reduce[string](words, func(carry, value string) string {
		return carry + " " + value
	})

	if reduced != "gopher monkey robot ninja" {
		t.Errorf("Expected 'gopher monkey robot ninja', got '%s'", reduced)
	}
}
