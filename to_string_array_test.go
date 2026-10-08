package slice_utils_test

import (
	"reflect"
	"testing"

	"github.com/codeus-node/slice_utils"
)

func Test_ToStringArray_WithInts(t *testing.T) {
	arr := slice_utils.ToStringArray([]int{1, 2, 3})

	expected := []string{"1", "2", "3"}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func Test_ToStringArray_WithStrings(t *testing.T) {
	arr := slice_utils.ToStringArray([]string{"a", "b", "c"})

	expected := []string{"a", "b", "c"}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func Test_ToStringArray_WithFloats(t *testing.T) {
	arr := slice_utils.ToStringArray([]float64{1.5, 2.25, 3.75})

	expected := []string{"1.5", "2.25", "3.75"}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func Test_ToStringArray_WithBools(t *testing.T) {
	arr := slice_utils.ToStringArray([]bool{true, false})

	expected := []string{"true", "false"}

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}

func Test_ToStringArray_WithEmptySlice(t *testing.T) {
	arr := slice_utils.ToStringArray([]int{})

	expected := make([]string, 0)

	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("Expected %v, got %v", expected, arr)
	}
}
