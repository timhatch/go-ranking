package rankers

import (
	"reflect"
	"sort"
	"testing"
)

func TestDenseRank(t *testing.T) {
	tests := []struct {
		name     string
		input    *MockRankable
		expected []float64
	}{
		{
			name:     "Test case 1",
			input:    mockRankable1,
			expected: []float64{1, 2, 2, 3},
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := DenseRank(tt.input)
			sort.Sort(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Expected %v, but got %v", tt.expected, result)
			}
		})
	}
}
