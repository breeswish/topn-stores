package store

type Store interface {
	Add(key int, value float64)
	FinishAdd()
	GetTopNItems(topN int) map[int]float64
	GetRoundTopNItems(topN int) map[int]float64
	GetPeakKeys() int
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

type KeyValue struct {
	key   int
	value float64
}

type SortByValue []KeyValue

func (a SortByValue) Len() int           { return len(a) }
func (a SortByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByValue) Less(i, j int) bool { return a[i].value > a[j].value }
