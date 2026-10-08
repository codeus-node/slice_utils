package slice_utils_test

import (
	"reflect"
	"testing"

	"github.com/codeus-node/slice_utils"
)

func Test_ToAnyArray(t *testing.T) {
	arr := slice_utils.ToAnyArray([]int{1, 2, 3})

	expected := []any{1, 2, 3}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}
