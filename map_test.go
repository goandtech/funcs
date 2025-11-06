package funcs

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {

	type testCase struct {
		name           string
		input          []int
		expectedOutput []string
	}

	testCases := []testCase{
		{name: "empty"},
		{name: "simple", input: []int{1, 2}, expectedOutput: []string{"1", "2"}},
	}

	for _, tc := range testCases {
		f := func(i int) string { return fmt.Sprint(i) }
		t.Run(tc.name, func(t *testing.T) {
			output := Map(tc.input, f)

			if len(output) != len(tc.input) {
				t.Errorf("ouput length [%d] != input lenght [%d]", len(output), len(tc.input))
			}
			for index, value := range tc.input {
				if f(value) != output[index] {
					t.Errorf("expected value %v at index %d but got %v", f(value), index, output[index])
				}
			}
		})
	}
}
