package store

import "github.com/wangjohn/quickselect"

var _ Store = &NaiveStore{}

type NaiveStore struct {
	Data     map[int]float64
	PeakKeys int
}

func NewNaiveStore() *NaiveStore {
	return &NaiveStore{
		Data: map[int]float64{},
	}
}

func (c *NaiveStore) GetPeakKeys() int {
	return c.PeakKeys
}

func (c *NaiveStore) Add(key int, value float64) {
	c.Data[key] += value
	c.PeakKeys = max(c.PeakKeys, len(c.Data))
}

func (c *NaiveStore) FinishAdd() {
	// Nothing
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
