package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"topn-stores/generator"
	"topn-stores/store"

	"go.uber.org/zap"
)

var distributions = []string{
	"constant",
	"uniform",
	"zipfian",
	"hotspot",
	"sequential",
}

func getDistribution(n int, name string, r *rand.Rand) generator.Generator {
	var g generator.Generator
	switch name {
	case "constant":
		g = generator.NewConstant(r.Int63n(int64(n)) + 1)
	case "uniform":
		g = generator.NewUniform(1, int64(n))
	case "zipfian":
		g = generator.NewZipfianWithRange(1, int64(n), generator.ZipfianConstant)
	case "hotspot":
		g = generator.NewHotspot(1, int64(n), 0.2, 0.8)
	case "sequential":
		g = generator.NewSequential(1, int64(n))
	default:
		log.Fatalf("unknown distribution %s", name)
	}
	return g
}

func CalcTopNEffectiveRateKey(base map[int]float64, target map[int]float64) float64 {
	hit := 0
	all := 0

	for key := range base {
		all += 1
		if _, ok := target[key]; ok {
			hit += 1
		}
	}

	return float64(hit) / float64(all)
}

func CalcTopNEffectiveRateValue(base map[int]float64, target map[int]float64) float64 {
	var baseSum float64
	var targetSum float64

	for key := range base {
		if _, ok := target[key]; ok {
			baseSum += base[key]
			targetSum += target[key]
		}
	}

	return targetSum / baseSum
}

func DecorateValue(v float64) string {
	if v < 0.3 {
		return fmt.Sprintf("🔴%f", v)
	} else if v < 0.5 {
		return fmt.Sprintf("🟡%f", v)
	} else if v > 0.9 {
		return fmt.Sprintf("🟢%f", v)
	} else {
		return fmt.Sprintf("  %f", v)
	}
}

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	const iterRounds = 60
	const topN = 200
	const finalTopN = 20

	r := rand.New(rand.NewSource(time.Now().Unix()))

	for _, dataRange := range []int{150, 200, 201, 400, 1000} {
		for _, dataPortation := range []float64{0.1, 0.5, 1, 2} {
			for _, keyDist := range distributions {
				for _, valueDist := range distributions {
					if keyDist == "sequential" && valueDist == "constant" {
						// This doesn't make sense when comparing top N, since every item is the same
						continue
					}

					gKey := getDistribution(dataRange, keyDist, r)
					gValue := getDistribution(101, valueDist, r)
	
					baseCache := store.NaiveStore{Data: map[int]float64{}}
					v1Cache := store.StoreV1{
						Data: map[int]float64{},
						TopN: topN,
					}
					v2Cache := store.StoreV2{
						Data: map[int]*store.StoreV2Value{},
						TopN: topN,
					}
					v3Cache := store.StoreV3{
						Data:    map[int]float64{},
						DataBuf: map[int]float64{},
						TopN:    topN,
					}
					//v4Cache := store.StoreV3{
					//	Data:    map[int]float64{},
					//	DataBuf: map[int]float64{},
					//	TopN:    topN / 60,
					//}
	
					for i := 0; i < iterRounds; i++ {
						iters := int(float64(dataRange) * dataPortation)
						for j := 0; j < iters; j++ {
							key := int(gKey.Next(r))
							value := float64(gValue.Next(r) - 1) / 100
							baseCache.Add(key, value)
							v1Cache.Add(key, value)
							v2Cache.Add(key, value)
							v3Cache.Add(key, value)
							//v4Cache.Add(key, value)
						}
						baseCache.FinishAdd()
						v1Cache.FinishAdd()
						v2Cache.FinishAdd()
						v3Cache.FinishAdd()
						//v4Cache.FinishAdd()
					}

					// Calculate effective rate
					baseItems := baseCache.GetTopNItems(finalTopN)
					v1Items := v1Cache.GetTopNItems(finalTopN)
					v2Items := v2Cache.GetTopNItems(finalTopN)
					v3Items := v3Cache.GetTopNItems(finalTopN)
					//v4Items := v4Cache.GetTopNItems(finalTopN)

					//zap.L().Info("cache", zap.Any("baseItems", baseItems), zap.Any("v3", v3Cache), zap.Any("v3Items", v3Items))

					v1KRate := CalcTopNEffectiveRateKey(baseItems, v1Items)
					v1VRate := CalcTopNEffectiveRateValue(baseItems, v1Items)
					v2KRate := CalcTopNEffectiveRateKey(baseItems, v2Items)
					v2VRate := CalcTopNEffectiveRateValue(baseItems, v2Items)
					v3KRate := CalcTopNEffectiveRateKey(baseItems, v3Items)
					v3VRate := CalcTopNEffectiveRateValue(baseItems, v3Items)
					//v4KRate := CalcTopNEffectiveRateKey(baseItems, v4Items)
					//v4VRate := CalcTopNEffectiveRateValue(baseItems, v4Items)

					if v1KRate < 0.5 || v2KRate < 0.5 || v3KRate < 0.5 {
						fmt.Printf("%d\t%f\t%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s\n", dataRange, dataPortation,
							keyDist[:3], valueDist[:3],
							DecorateValue(v1KRate), DecorateValue(v1VRate),
							DecorateValue(v2KRate), DecorateValue(v2VRate),
							DecorateValue(v3KRate), DecorateValue(v3VRate),
						)
					}
				}
			}
		}
	}
}
