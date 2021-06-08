package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"time"

	"topn-stores/generator"
	"topn-stores/store"

	"go.uber.org/zap"
)

const iterRounds = 60
const finalTopN = 20
const storeIdBase = "base"

var r = rand.New(rand.NewSource(time.Now().Unix()))

var newStoreFns = map[string]NewStoreFn{
	storeIdBase: func() store.Store {
		return store.NewNaiveStore()
	},
	"v1(top200)": func() store.Store {
		return store.NewStoreV1(200)
	},
	"v1(top2000)": func() store.Store {
		return store.NewStoreV1(2000)
	},
	"v2(top200)": func() store.Store {
		return store.NewStoreV2(200)
	},
	"v3(top200)": func() store.Store {
		return store.NewStoreV3(200)
	},
	"v3(top20)": func() store.Store {
		return store.NewStoreV3(20)
	},
}

func CalcKeyEfficiency(base map[int]float64, target map[int]float64) float64 {
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

func CalcValueEfficiency(base map[int]float64, target map[int]float64) float64 {
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

type NewStoreFn func() store.Store

type Suite struct {
	DataRange         int
	NumberOfKeys      float64
	KeyDistribution   string
	ValueDistribution string
}

func (s *Suite) String() string {
	return fmt.Sprintf("range=%6d\tkeys=%6d\tkeyDist=%11s\tvalueDist=%11s",
		s.DataRange,
		int(float64(s.DataRange)*s.NumberOfKeys),
		s.KeyDistribution,
		s.ValueDistribution)
}

func (s *Suite) Run() map[string]StoreResult {
	result := map[string]StoreResult{}

	gKey := generator.GetDistribution(s.DataRange, s.KeyDistribution, r)
	gValue := generator.GetDistribution(101, s.ValueDistribution, r)

	stores := map[string]store.Store{}
	for key, fn := range newStoreFns {
		stores[key] = fn()
	}

	for i := 0; i < iterRounds; i++ {
		iters := int(float64(s.DataRange) * s.NumberOfKeys)
		for j := 0; j < iters; j++ {
			key := int(gKey.Next(r))
			value := float64(gValue.Next(r)-1) / 100
			for _, s := range stores {
				s.Add(key, value)
			}
		}
		for _, s := range stores {
			s.FinishAdd()
		}
	}

	// Calculate effective rate
	baseResult := stores[storeIdBase].GetTopNItems(finalTopN)
	for key, s := range stores {
		if key == storeIdBase {
			continue
		}
		items := s.GetTopNItems(finalTopN)
		keyRate := CalcKeyEfficiency(baseResult, items)
		if math.IsNaN(keyRate) {
			keyRate = 0
		}
		valueRate := CalcValueEfficiency(baseResult, items)
		if math.IsNaN(valueRate) {
			valueRate = 0
		}
		result[key] = StoreResult{
			KeyRate:   keyRate,
			ValueRate: valueRate,
			PeakKeys:  s.GetPeakKeys(),
		}
	}

	return result
}

type StoreResult struct {
	KeyRate   float64
	ValueRate float64
	PeakKeys  int
}

func main() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)

	suites := make([]*Suite, 0)
	for _, dataRange := range []int{150, 200, 201, 400, 1000} {
		for _, numberOfKeys := range []float64{0.1, 0.5, 1, 2} {
			for _, keyDist := range generator.AllDistributions {
				for _, valueDist := range generator.AllDistributions {
					if keyDist == "sequential" && valueDist == "constant" {
						// This doesn't make sense when comparing top N, since every item is the same
						continue
					}
					suites = append(suites, &Suite{
						DataRange:         dataRange,
						NumberOfKeys:      numberOfKeys,
						KeyDistribution:   keyDist,
						ValueDistribution: valueDist,
					})
				}
			}
		}
	}

	type suiteResult struct {
		Suite  *Suite
		Result StoreResult
	}
	var resultsByStore = map[string][]suiteResult{}
	for key := range newStoreFns {
		resultsByStore[key] = make([]suiteResult, 0)
	}

	allSuites := len(suites)
	for idx, suite := range suites {
		zap.L().Info("Running suite", zap.Any("suite", suite), zap.Int("idx", idx), zap.Int("len", allSuites))
		r := suite.Run()
		// Aggregate results of each store
		for key, result := range r {
			resultsByStore[key] = append(resultsByStore[key], suiteResult{
				Suite:  suite,
				Result: result,
			})
		}
	}

	// Finally, output results for each store
	keys := make([]string, 0)
	for k := range resultsByStore {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		fmt.Println("=================================")
		fmt.Printf("Store %s:\n", k)
		sort.SliceStable(resultsByStore[k], func(i, j int) bool {
			return resultsByStore[k][i].Result.KeyRate < resultsByStore[k][j].Result.KeyRate
		})

		totalKeyRate := 0.0
		totalValueRate := 0.0
		totalPeakKeys := 0
		maxPeakKeys := 0
		for _, result := range resultsByStore[k] {
			totalKeyRate += result.Result.KeyRate
			totalValueRate += result.Result.ValueRate
			totalPeakKeys += result.Result.PeakKeys
			if result.Result.PeakKeys > maxPeakKeys {
				maxPeakKeys = result.Result.PeakKeys
			}
		}

		fmt.Printf("Avg key efficiency: %f, Avg value efficiency: %f, Avg peak keys: %f, Max peak keys: %d\n",
			totalKeyRate/float64(len(resultsByStore[k])),
			totalValueRate/float64(len(resultsByStore[k])),
			float64(totalPeakKeys)/float64(len(resultsByStore[k])),
			maxPeakKeys)
		fmt.Println("Bad cases:")

		for _, result := range resultsByStore[k] {
			if result.Result.ValueRate < 0.5 || result.Result.KeyRate < 0.5 {
				fmt.Printf("kEff=%f\tvEff=%f\tkPeak=%6d\t%s\n",
					result.Result.KeyRate,
					result.Result.ValueRate,
					result.Result.PeakKeys,
					result.Suite.String())
			}
		}

		fmt.Println()
	}
}
