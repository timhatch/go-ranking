package rankers

import (
	"reflect"
	"testing"
)

type scores []struct {
	score int
	name  string
}

func (s scores) Len() int                { return len(s) }
func (s scores) RankEqual(i, j int) bool { return s[i].score == s[j].score }
func (s scores) Swap(i, j int)           { s[i], s[j] = s[j], s[i] }
func (s scores) Less(i, j int) bool {
	if s[i].score != s[j].score {
		return s[i].score > s[j].score
	}
	return s[i].name < s[j].name
}

var data = scores{
	{44, "Solomon"},
	{42, "Jason"},
	{42, "Errol"},
	{41, "Garry"},
	{41, "Bernard"},
	{41, "Barry"},
	{39, "Stephen"},
}

func TestStandardRank(t *testing.T) {
	r := StandardRank(data)
	w := []float64{1, 2, 2, 4, 4, 4, 7}

	if !reflect.DeepEqual(r, w) {
		t.Errorf("want ranking %v, got %v", w, r)
	}
}

func TestModifiedRank(t *testing.T) {
	r := ModifiedRank(data)
	w := []float64{1, 3, 3, 6, 6, 6, 7}

	if !reflect.DeepEqual(r, w) {
		t.Errorf("want ranking %v, got %v", w, r)
	}
}

func TestDenseRank(t *testing.T) {
	r := DenseRank(data)
	w := []float64{1, 2, 2, 3, 3, 3, 4}

	if !reflect.DeepEqual(r, w) {
		t.Errorf("want ranking %v, got %v", w, r)
	}
}

func TestOrdinalRank(t *testing.T) {
	r := OrdinalRank(data)
	w := []float64{1, 2, 3, 4, 5, 6, 7}

	if !reflect.DeepEqual(r, w) {
		t.Errorf("want ranking %v, got %v", w, r)
	}
}

func TestFractionalRank(t *testing.T) {
	r := FractionalRank(data)
	w := []float64{1, 2.5, 2.5, 5, 5, 5, 7}

	if !reflect.DeepEqual(r, w) {
		t.Errorf("want ranking %v, got %v", w, r)
	}
}
