package store

import "github.com/wangjohn/quickselect"

var _ Store = &StoreV2{}

type StoreV2Value struct {
	Value         float64
	OccupyWindows int
}

type StoreV2 struct {
	Data     map[int]*StoreV2Value
	TopN     int
	PeakKeys int
}

func NewStoreV2(topN int) *StoreV2 {
	return &StoreV2{
		Data:     map[int]*StoreV2Value{},
		TopN:     topN,
		PeakKeys: 0,
	}
}

func (c *StoreV2) GetPeakKeys() int {
	return c.PeakKeys
}

func (c *StoreV2) Add(key int, value float64) {
	_, ok := c.Data[key]
	if !ok {
		c.Data[key] = &StoreV2Value{
			Value:         value,
			OccupyWindows: 1,
		}
	} else {
		c.Data[key].Value += value
		c.Data[key].OccupyWindows += 1
	}
}

type KeyValueAndTimes struct {
	key           int
	value         float64
	occupyWindows int
}

type SortByAverageValue []KeyValueAndTimes

func (a SortByAverageValue) Len() int      { return len(a) }
func (a SortByAverageValue) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a SortByAverageValue) Less(i, j int) bool {
	return a[i].value/float64(a[i].occupyWindows) > a[j].value/float64(a[j].occupyWindows)
}

func (c *StoreV2) FinishAdd() {
	defer func() {
		c.PeakKeys = max(c.PeakKeys, len(c.Data))
	}()

	slices := make([]KeyValueAndTimes, 0)
	for key, value := range c.Data {
		slices = append(slices, KeyValueAndTimes{
			key:           key,
			value:         value.Value,
			occupyWindows: value.OccupyWindows,
		})
	}
	if len(slices) <= c.TopN {
		return
	}
	_ = quickselect.QuickSelect(SortByAverageValue(slices), c.TopN)
	for _, item := range slices[c.TopN:] {
		delete(c.Data, item.key)
	}
}

func (c *StoreV2) GetRoundTopNItems(topN int) map[int]float64 {
	return c.GetTopNItems(topN)
}

func (c *StoreV2) GetTopNItems(topN int) map[int]float64 {
	slices := make([]KeyValue, 0)
	for key, value := range c.Data {
		slices = append(slices, KeyValue{
			key:   key,
			value: value.Value,
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
