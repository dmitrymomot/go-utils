package utils_test

import (
	"reflect"
	"testing"

	"github.com/dmitrymomot/go-utils"
)

func TestMergeIntoMap(t *testing.T) {
	// Test case #1: Merge maps with unique keys
	dst1 := map[string]interface{}{
		"a": 1,
		"b": "two",
	}
	src1 := map[string]interface{}{
		"c": 3.5,
		"d": true,
	}
	expectedResult1 := map[string]interface{}{
		"a": 1,
		"b": "two",
		"c": 3.5,
		"d": true,
	}
	result1 := utils.MergeIntoMap(dst1, src1)
	if !reflect.DeepEqual(result1, expectedResult1) {
		t.Errorf("Test case #1 failed: Expected %v but got %v", expectedResult1, result1)
	}

	// Test case #2: Merge maps with overlapping keys (values from src should overwrite dst)
	dst2 := map[string]interface{}{
		"a": 1,
		"b": "original",
		"c": true,
	}
	src2 := map[string]interface{}{
		"b": "new value",
		"d": 42,
	}
	expectedResult2 := map[string]interface{}{
		"a": 1,
		"b": "new value",
		"c": true,
		"d": 42,
	}
	result2 := utils.MergeIntoMap(dst2, src2)
	if !reflect.DeepEqual(result2, expectedResult2) {
		t.Errorf("Test case #2 failed: Expected %v but got %v", expectedResult2, result2)
	}

	// Test case #3: dst is nil, src is not nil
	var dst3 map[string]interface{}
	src3 := map[string]interface{}{
		"a": 1,
	}
	expectedResult3 := map[string]interface{}{
		"a": 1,
	}
	result3 := utils.MergeIntoMap(dst3, src3)
	if !reflect.DeepEqual(result3, expectedResult3) {
		t.Errorf("Test case #3 failed: Expected %v but got %v", expectedResult3, result3)
	}

	// Test case #4: src is nil, dst is not nil
	dst4 := map[string]interface{}{
		"a": 1,
	}
	var src4 map[string]interface{}
	expectedResult4 := map[string]interface{}{
		"a": 1,
	}
	result4 := utils.MergeIntoMap(dst4, src4)
	if !reflect.DeepEqual(result4, expectedResult4) {
		t.Errorf("Test case #4 failed: Expected %v but got %v", expectedResult4, result4)
	}

	// Test case #5: Both maps are nil
	var dst5, src5 map[string]interface{}
	expectedResult5 := map[string]interface{}{}
	result5 := utils.MergeIntoMap(dst5, src5)
	if !reflect.DeepEqual(result5, expectedResult5) {
		t.Errorf("Test case #5 failed: Expected %v but got %v", expectedResult5, result5)
	}
}

func TestMergeMapsRecursively(t *testing.T) {
	type testCase struct {
		dst      map[string]interface{}
		src      map[string]interface{}
		expected map[string]interface{}
	}

	testCases := []testCase{
		{
			// Merging empty maps should return an empty map
			dst:      map[string]interface{}{},
			src:      map[string]interface{}{},
			expected: map[string]interface{}{},
		},
		{
			// Merging nil maps should return the other map
			dst:      nil,
			src:      map[string]interface{}{"foo": 123},
			expected: map[string]interface{}{"foo": 123},
		},
		{
			// Merging non-nested maps should combine values
			dst:      map[string]interface{}{"foo": 123},
			src:      map[string]interface{}{"bar": "hello"},
			expected: map[string]interface{}{"foo": 123, "bar": "hello"},
		},
		{
			// Merging nested maps should combine recursively
			dst: map[string]interface{}{
				"foo": map[string]interface{}{
					"a": 1,
					"b": 2,
				},
			},
			src: map[string]interface{}{
				"foo": map[string]interface{}{
					"b": 3,
					"c": 4,
				},
			},
			expected: map[string]interface{}{
				"foo": map[string]interface{}{
					"a": 1,
					"b": 3,
					"c": 4,
				},
			},
		},
		{
			// Merging maps with different value types should replace
			dst: map[string]interface{}{
				"foo": map[string]interface{}{
					"a": 1,
					"b": 2,
				},
			},
			src: map[string]interface{}{
				"foo": "hello",
			},
			expected: map[string]interface{}{
				"foo": "hello",
			},
		},
	}

	for i, tc := range testCases {
		result := utils.MergeIntoMapRecursively(tc.dst, tc.src)

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("Test case %d failed: expected %v but got %v", i, tc.expected, result)
		}
	}
}
