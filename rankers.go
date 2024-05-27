package rankers

// Ranking methods from Rosetta Code,
// Implements the following methods:
// 1. Standard. (Ties share what would have been their first ordinal number).
// 2. Modified. (Ties share what would have been their last ordinal number).
// 3. Dense. (Ties share the next available integer).
// 4. Ordinal. ((Competitors take the next available integer. Ties are not
//    treated otherwise).
// 5. Fractional. (Ties share the mean of what would have been their ordinal
//    numbers).

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
