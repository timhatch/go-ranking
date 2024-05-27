package rankers

import (
	"sort"
)

// Return the dense rankings for elements in a `rankable`
// e.g. 1, 2, 2, 3, 4
func DenseRank(d Rankable) []float64 {
	r := make([]float64, d.Len())
	sort.Sort(d)

	var k int
	for i := range r {
		if i == 0 || !d.RankEqual(i, i-1) {
			k++
		}
		r[i] = float64(k)
	}
	return r
}
