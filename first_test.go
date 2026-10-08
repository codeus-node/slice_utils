package slice_utils_test

import (
	"testing"

	"github.com/codeus-node/slice_utils"
)

func TestFirst(t *testing.T) {
	t.Run("returns the first matching element when a match exists", func(t *testing.T) {
		numbers := []int{10, 20, 30, 40}

		result, err := slice_utils.First(numbers, func(n int) bool {
			return n > 25
		})

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if result != 30 {
			t.Errorf("expected 30, got: %d", result)
		}
	})

	t.Run("returns an error and zero value when no match is found", func(t *testing.T) {
		words := []string{"apple", "banana", "cherry"}

		result, err := slice_utils.First(words, func(s string) bool {
			return s == "dragonfruit"
		})

		if err == nil {
			t.Fatal("expected an error, got nil")
		}
		if result != "" {
			t.Errorf("expected empty string as zero value, got: %q", result)
		}
	})

	t.Run("returns an error and zero value when the input slice is empty", func(t *testing.T) {
		var emptySlice []int

		result, err := slice_utils.First(emptySlice, func(n int) bool {
			return n == 1
		})

		if err == nil {
			t.Fatal("expected an error for an empty slice, got nil")
		}
		if result != 0 {
			t.Errorf("expected zero value 0, got: %d", result)
		}
	})

	t.Run("works correctly with custom struct types", func(t *testing.T) {
		type User struct {
			ID   int
			Name string
		}

		users := []User{
			{ID: 1, Name: "Alice"},
			{ID: 2, Name: "Bob"},
			{ID: 3, Name: "Charlie"},
		}

		result, err := slice_utils.First(users, func(u User) bool {
			return u.ID == 2
		})

		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if result.Name != "Bob" {
			t.Errorf("expected user name 'Bob', got: %q", result.Name)
		}
	})
}

func TestFirst_TableDriven(t *testing.T) {
	tests := []struct {
		name          string
		input         []int
		filter        func(int) bool
		expectedValue int
		expectError   bool
	}{
		{
			name:  "match at start of slice",
			input: []int{5, 10, 15},
			filter: func(n int) bool {
				return n == 5
			},
			expectedValue: 5,
			expectError:   false,
		},
		{
			name:  "match at end of slice",
			input: []int{5, 10, 15},
			filter: func(n int) bool {
				return n == 15
			},
			expectedValue: 15,
			expectError:   false,
		},
		{
			name:  "no match present",
			input: []int{1, 2, 3},
			filter: func(n int) bool {
				return n == 99
			},
			expectedValue: 0,
			expectError:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := slice_utils.First(tt.input, tt.filter)

			if (err != nil) != tt.expectError {
				t.Fatalf("expectError = %v, but got error: %v", tt.expectError, err)
			}
			if got != tt.expectedValue {
				t.Errorf("expected result %d, got %d", tt.expectedValue, got)
			}
		})
	}
}
