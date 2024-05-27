package rankers

import (
	"sort"
)

// Return the Ordinal rankings for elements in a `rankable`
// e.g. 1, 2, 3, 4, 5
func OrdinalRank(d Rankable) []float64 {
	r := make([]float64, d.Len())
	sort.Sort(d)

	for i := range r {
		r[i] = float64(i + 1)
	}
	return r
}
