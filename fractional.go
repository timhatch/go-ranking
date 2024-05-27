package rankers

import (
	"sort"
)

// Return the fractional rankings for elements in a `rankable`
// e.g. 1, 2.5, 2.5, 4, 5
func FractionalRank(d Rankable) []float64 {
	r := make([]float64, d.Len())
	sort.Sort(d)

	for i := 0; i < d.Len(); {
		j := i + 1
		f := float64(i + 1)

		for ; j < d.Len() && d.RankEqual(i, j); j++ {
			f += float64(j + 1)
		}

		f /= float64(j - i)
		for ; i < j; i++ {
			r[i] = f
		}
	}
	return r
}
