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

const rounds = 60
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

func CalcKeyAccuracy(base map[int]float64, target map[int]float64) float64 {
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

func CalcValueAccuracy(base map[int]float64, target map[int]float64) float64 {
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
	roundKeyRateSumByStore := map[string]float64{}
	for key, fn := range newStoreFns {
		stores[key] = fn()
	}

	for roundIdx := 0; roundIdx < rounds; roundIdx++ {
		keys := int(float64(s.DataRange) * s.NumberOfKeys)
		for keyIdx := 0; keyIdx < keys; keyIdx++ {
			key := int(gKey.Next(r))
			value := float64(gValue.Next(r)-1) / 100
			for _, s := range stores {
				s.Add(key, value)
			}
		}

		for _, s := range stores {
			s.FinishAdd()
		}

		// Record per-round key accuracy
		baseResult := stores[storeIdBase].GetRoundTopNItems(finalTopN)
		for key, s := range stores {
			if key == storeIdBase {
				continue
			}
			items := s.GetRoundTopNItems(finalTopN)
			keyRate := CalcKeyAccuracy(baseResult, items)
			if math.IsNaN(keyRate) {
				keyRate = 0
			}
			roundKeyRateSumByStore[key] += keyRate
		}
	}

	// Record final accuracy
	baseResult := stores[storeIdBase].GetTopNItems(finalTopN)
	for key, s := range stores {
		if key == storeIdBase {
			continue
		}
		items := s.GetTopNItems(finalTopN)
		keyRate := CalcKeyAccuracy(baseResult, items)
		if math.IsNaN(keyRate) {
			keyRate = 0
		}
		valueRate := CalcValueAccuracy(baseResult, items)
		if math.IsNaN(valueRate) {
			valueRate = 0
		}
		result[key] = StoreResult{
			FinalKeyRate:    keyRate,
			FinalValueRate:  valueRate,
			AvgRoundKeyRate: roundKeyRateSumByStore[key] / rounds,
			PeakKeys:        s.GetPeakKeys(),
		}
	}

	return result
}

type StoreResult struct {
	FinalKeyRate    float64
	FinalValueRate  float64
	AvgRoundKeyRate float64
	PeakKeys        int
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
			return resultsByStore[k][i].Result.FinalKeyRate < resultsByStore[k][j].Result.FinalKeyRate
		})

		sumFinalKeyRate := 0.0
		sumFinalValueRate := 0.0
		sumPerRoundKeyRate := 0.0
		sumPeakKeys := 0
		maxPeakKeys := 0
		for _, result := range resultsByStore[k] {
			sumFinalKeyRate += result.Result.FinalKeyRate
			sumFinalValueRate += result.Result.FinalValueRate
			sumPerRoundKeyRate += result.Result.AvgRoundKeyRate
			sumPeakKeys += result.Result.PeakKeys
			if result.Result.PeakKeys > maxPeakKeys {
				maxPeakKeys = result.Result.PeakKeys
			}
		}

		fmt.Printf("Avg key accuracy: %f, Avg value accuracy: %f, Avg key accuracy per round: %f, Avg peak keys: %f, Max peak keys: %d\n",
			sumFinalKeyRate/float64(len(resultsByStore[k])),
			sumFinalValueRate/float64(len(resultsByStore[k])),
			sumPerRoundKeyRate/float64(len(resultsByStore[k])),
			float64(sumPeakKeys)/float64(len(resultsByStore[k])),
			maxPeakKeys)
		fmt.Println("Bad cases:")

		for _, result := range resultsByStore[k] {
			if result.Result.FinalValueRate < 0.5 || result.Result.FinalKeyRate < 0.5 || result.Result.AvgRoundKeyRate < 0.5 {
				fmt.Printf("kAcc=%f\tvAcc=%f\tkAccPerRound=%f\tkPeak=%6d\t%s\n",
					result.Result.FinalKeyRate,
					result.Result.FinalValueRate,
					result.Result.AvgRoundKeyRate,
					result.Result.PeakKeys,
					result.Suite.String())
			}
		}

		fmt.Println()
	}
}
