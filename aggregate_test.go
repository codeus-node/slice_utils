package slice_utils_test

import (
	"testing"

	"github.com/codeus-node/slice_utils"
)

func TestAggregate(t *testing.T) {
	tests := []struct {
		name    string
		input   []int
		initial int
		aggF    func(int, int, int) int
		want    int
	}{
		{
			name:    "Calculate the sum",
			input:   []int{1, 2, 3, 4},
			initial: 0,
			aggF: func(i int, val int, acc int) int {
				return acc + val
			},
			want: 10,
		},
		{
			name:    "Include the index in the calculation",
			input:   []int{10, 10, 10},
			initial: 0,
			aggF: func(i int, val int, acc int) int {
				return acc + val + i
			},
			want: 33,
		},
		{
			name:    "An empty slice returns the initial value",
			input:   []int{},
			initial: 42,
			aggF: func(i int, val int, acc int) int {
				return acc + val
			},
			want: 42,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := slice_utils.Aggregate(tt.input, tt.aggF, tt.initial)
			if got != tt.want {
				t.Errorf("Aggregate() = %v, want %v", got, tt.want)
			}
		})
	}
}
