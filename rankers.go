package rankers

// Ranking methods from Rosetta Code,
// Implements the following methods:
// 1. Standard. (Ties share their first ordinal number).
// 2. Modified. (Ties share their last ordinal number).
// 3. Dense. (Ties share the next available integer).
// 4. Ordinal. (All elements take the next available integer. ties are ignored).
// 5. Fractional. (Ties share the mean of their ordinal numbers).

// Define a `rankable` interface type.
// A `rankable` type must provide two methods:
type Rankable interface {
	// For rank allocation
	Len() int                // return the number of elements to be ranked
	RankEqual(int, int) bool // return true if two elements are equal
	// FOr initial sorting
	Swap(int, int)      // swap elemenes within a Rankable array
	Less(int, int) bool // return true if element i < j
}
