package rankers

import (
	"sort"
)

// Return the modified rankings for elements in a `rankable`
// e.g. 1, 3, 3, 4, 5
func ModifiedRank(d Rankable) []float64 {
	r := make([]float64, d.Len())
	sort.Sort(d)

	for i := range r {
		k := i + 1
		for j := i + 1; j < len(r) && d.RankEqual(i, j); j++ {
			k = j + 1
		}
		r[i] = float64(k)
	}
	return r
}
