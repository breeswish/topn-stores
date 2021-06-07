package store

import "github.com/wangjohn/quickselect"

type StoreV3 struct {
	Data    map[int]float64
	DataBuf map[int]float64
	TopN    int
}

func (c *StoreV3) Add(key int, value float64) {
	c.DataBuf[key] += value
}

func (c *StoreV3) FinishAdd() {
	slices := make([]KeyValue, 0)
	for key, value := range c.DataBuf {
		slices = append(slices, KeyValue{
			key:           key,
			value:         value,
		})
	}

	if len(slices) > c.TopN {
		_ = quickselect.QuickSelect(SortByValue(slices), c.TopN)
		slices = slices[:c.TopN]
	}
	for _, item := range slices {
		c.Data[item.key] += item.value
	}
	c.DataBuf = map[int]float64{}
}

func (c *StoreV3) GetTopNItems(topN int) map[int]float64 {
	slices := make([]KeyValue, 0)
	for key, value := range c.Data {
		slices = append(slices, KeyValue{
			key:           key,
			value:         value,
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
