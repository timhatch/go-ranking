package rankers

// Ranking methods from Rosetta Code,
// Implements the following methods:
// 1. Standard. (Ties share what would have been their first ordinal number).
// 2. Modified. (Ties share what would have been their last ordinal number).
// 3. Dense. (Ties share the next available integer).
// 4. Ordinal. ((Competitors take the next available integer. Ties are not treated otherwise).
// 5. Fractional. (Ties share the mean of what would have been their ordinal numbers).

// Define a `rankable` interface type. A `rankable` type must provide two methods:
type rankable interface {
	Len() int                // return the number of elements to be ranked
	RankEqual(int, int) bool // return true if two elements are equal
}

// Return the standard rankings for elements in a `rankable`
// e.g. 1, 2, 2, 4, 5
func StandardRank(d rankable) []float64 {
	r := make([]float64, d.Len())
	var k int
	for i := range r {
		if i == 0 || !d.RankEqual(i, i-1) {
			k = i + 1
		}
		r[i] = float64(k)
	}
	return r
}

// Return the modified rankings for elements in a `rankable`
// e.g. 1, 3, 3, 4, 5
func ModifiedRank(d rankable) []float64 {
	r := make([]float64, d.Len())
	for i := range r {
		k := i + 1
		for j := i + 1; j < len(r) && d.RankEqual(i, j); j++ {
			k = j + 1
		}
		r[i] = float64(k)
	}
	return r
}

// Return the dense rankings for elements in a `rankable`
// e.g. 1, 2, 2, 3, 4
func DenseRank(d rankable) []float64 {
	r := make([]float64, d.Len())
	var k int
	for i := range r {
		if i == 0 || !d.RankEqual(i, i-1) {
			k++
		}
		r[i] = float64(k)
	}
	return r
}

// Return the Ordinal rankings for elements in a `rankable`
// e.g. 1, 2, 3, 4, 5
func OrdinalRank(d rankable) []float64 {
	r := make([]float64, d.Len())
	for i := range r {
		r[i] = float64(i + 1)
	}
	return r
}

// Return the fractional rankings for elements in a `rankable`
// e.g. 1, 2.5, 2.5, 4, 5
func FractionalRank(d rankable) []float64 {
	r := make([]float64, d.Len())
	for i := 0; i < len(r); {
		var j int
		f := float64(i + 1)
		for j = i + 1; j < len(r) && d.RankEqual(i, j); j++ {
			f += float64(j + 1)
		}
		f /= float64(j - i)
		for ; i < j; i++ {
			r[i] = f
		}
	}
	return r
}
