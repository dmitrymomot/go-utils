package utils_test

import (
	"testing"

	"github.com/dmitrymomot/go-utils"
)

func TestGetVarType(t *testing.T) {
	// Define test cases as a map with input values and expected outputs.
	tests := map[string]struct {
		input  interface{}
		output string
	}{
		"test1": {input: 123, output: "int"},
		"test2": {input: "Hello, world!", output: "string"},
		"test3": {input: true, output: "bool"},
		"test4": {input: 3.14159, output: "float64"},
	}

	// Iterate over each test case and run the function with the input value.
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := utils.GetVarType(tc.input)

			// Compare the actual output with the expected output.
			if result != tc.output {
				t.Errorf("Expected %s, but got %s", tc.output, result)
			}
		})
	}
}
