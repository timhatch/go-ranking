package rankers

type MockRankable struct {
	data []struct {
		score int
		name  string
	}
}

func (m *MockRankable) Len() int {
	return len(m.data)
}

func (m *MockRankable) RankEqual(i, j int) bool {
	return m.data[i].score == m.data[j].score
}

func (m *MockRankable) Swap(i, j int) {
	m.data[i], m.data[j] = m.data[j], m.data[i]
}

func (m *MockRankable) Less(i, j int) bool {
	if m.data[i].score != m.data[j].score {
		return m.data[i].score < m.data[j].score
	}
	return m.data[i].name < m.data[j].name
}

var mockRankable1 = &MockRankable{
	data: []struct {
		score int
		name  string
	}{
		{3, "A"},
		{2, "B"},
		{2, "C"},
		{1, "D"},
	},
}
