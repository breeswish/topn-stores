package store

import (
	"github.com/wangjohn/quickselect"
)

var _ Store = &StoreV1{}

type StoreV1 struct {
	Data     map[int]float64
	TopN     int
	PeakKeys int
}

func NewStoreV1(topN int) *StoreV1 {
	return &StoreV1{
		Data:     map[int]float64{},
		TopN:     topN,
		PeakKeys: 0,
	}
}

func (c *StoreV1) GetPeakKeys() int {
	return c.PeakKeys
}

func (c *StoreV1) Add(key int, value float64) {
	c.Data[key] += value
}

func (c *StoreV1) FinishAdd() {
	defer func() {
		c.PeakKeys = max(c.PeakKeys, len(c.Data))
	}()

	slices := make([]KeyValue, 0)
	for key, value := range c.Data {
		slices = append(slices, KeyValue{
			key:   key,
			value: value,
		})
	}
	if len(slices) <= c.TopN {
		return
	}
	_ = quickselect.QuickSelect(SortByValue(slices), c.TopN)
	for _, item := range slices[c.TopN:] {
		delete(c.Data, item.key)
	}
}

func (c *StoreV1) GetTopNItems(topN int) map[int]float64 {
	slices := make([]KeyValue, 0)
	for key, value := range c.Data {
		slices = append(slices, KeyValue{
			key:   key,
			value: value,
		})
	}
	if len(slices) > topN {
		_ = quickselect.QuickSelect(SortByValue(slices), topN)
		slices = slices[:topN]
	}
	r := map[int]float64{}
	for _, item := range slices {
		r[item.key] = item.value
	}
	return r
}
