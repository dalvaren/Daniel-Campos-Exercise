package tasks

import (
	"testing"
)

func TestSumNumbers(t *testing.T) {
	testCases := []struct {a, b, result int}{
		{1, 2, 3},
		{3, 2, 5},
		{0, 4, 4},
		{-1, 2, 1},
	}

	for _, testCase := range testCases {
		if SumNumbers(testCase.a, testCase.b) != testCase.result {
			t.Errorf("SumNumbers(%q, %q) should be equal %q", testCase.a, testCase.b, testCase.result)
		}
	}

}
