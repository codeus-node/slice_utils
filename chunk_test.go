package slice_utils_test

import (
	"testing"

	"github.com/codeus-node/slice_utils"
)

func Test_Chunk(t *testing.T) {
	chunks := slice_utils.Chunk([]string{"a", "b", "c"}, 2)
	if len(chunks) != 2 {
		t.Errorf("expect 2 chunks, got %v", len(chunks))
	}
	if len(chunks[0]) != 2 {
		t.Errorf("expect 2 elements in first chunk, got %v", len(chunks[0]))
	}
	if len(chunks[1]) != 1 {
		t.Errorf("expect 1 element in second chunk, got %v", len(chunks[1]))
	}
	if chunks[0][0] != "a" {
		t.Errorf("expect 'a' in first chunk, got %v", chunks[0][0])
	}
	if chunks[0][1] != "b" {
		t.Errorf("expect 'b' in first chunk, got %v", chunks[0][1])
	}
	if chunks[1][0] != "c" {
		t.Errorf("expect 'c' in second chunk, got %v", chunks[1][0])
	}
}
