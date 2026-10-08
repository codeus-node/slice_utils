package slice_utils_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/codeus-node/slice_utils"
)

func TestConvert(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		conF     func(int) string
		expected []string
	}{
		{
			name:  "Converting Numbers to Strings",
			input: []int{1, 2, 3},
			conF: func(v int) string {
				return strconv.Itoa(v)
			},
			expected: []string{"1", "2", "3"},
		},
		{
			name:  "Process an empty slice",
			input: []int{},
			conF: func(v int) string {
				return strconv.Itoa(v)
			},
			expected: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice_utils.Convert(tt.input, tt.conF)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("Convert() = %v, want %v", got, tt.expected)
			}
		})
	}
}
