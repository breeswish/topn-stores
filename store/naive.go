package store

import "github.com/wangjohn/quickselect"

type NaiveStore struct {
	Data map[int]float64
}

func (c *NaiveStore) Add(key int, value float64) {
	c.Data[key] += value
}

func (c *NaiveStore) FinishAdd() {

}

func (c *NaiveStore) GetTopNItems(topN int) map[int]float64 {
	slices := make([]KeyValue, 0)
	for key, value := range c.Data {
		slices = append(slices, KeyValue{
			key:   key,
			value: value,
		})
	}
	if len(slices) > topN {
		_ = quickselect.QuickSelect(SortByValue(slices), topN)
		slices = slices[0:topN]
	}
	r := map[int]float64{}
	for _, item := range slices {
		r[item.key] = item.value
	}
	return r
}

type KeyValue struct {
	key   int
	value float64
}

type SortByValue []KeyValue

func (a SortByValue) Len() int           { return len(a) }
func (a SortByValue) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByValue) Less(i, j int) bool { return a[i].value > a[j].value }
