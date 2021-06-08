# TopN Store

## Stores

- StoreV1: Evict by the sum of value
- StoreV2: Evict by the average sum of value
- StoreV3: Evict by the value of each batch

## Results

```
=================================
Store v1(top200):
Avg key efficiency: 0.937604, Avg value efficiency: 0.996485, Avg peak keys: 148.679167, Max peak keys: 200
Bad cases:
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   400	keyDist=    hotspot	valueDist=   constant
kEff=0.200000	vEff=0.994126	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist=    hotspot
kEff=0.250000	vEff=0.973099	kPeak=   200	range=  1000	keys=   500	keyDist= sequential	valueDist=    uniform
kEff=0.250000	vEff=1.000000	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist=    uniform
kEff=0.300000	vEff=0.978695	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    zipfian
kEff=0.300000	vEff=1.000000	kPeak=   200	range=  1000	keys=  2000	keyDist= sequential	valueDist=    uniform
kEff=0.350000	vEff=0.993730	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist=   constant
kEff=0.350000	vEff=0.986892	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist=    zipfian
kEff=0.350000	vEff=0.994629	kPeak=   200	range=  1000	keys=  2000	keyDist= sequential	valueDist=    hotspot
kEff=0.400000	vEff=0.984879	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    uniform
kEff=0.400000	vEff=1.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist= sequential
kEff=0.400000	vEff=1.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist=   constant
kEff=0.400000	vEff=0.996304	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist= sequential
kEff=0.450000	vEff=0.950413	kPeak=   200	range=  1000	keys=   100	keyDist=    uniform	valueDist=   constant
kEff=0.450000	vEff=0.993963	kPeak=   200	range=  1000	keys=   100	keyDist= sequential	valueDist=    uniform
kEff=0.450000	vEff=0.991690	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist=    uniform
kEff=0.450000	vEff=0.984572	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist=   constant
kEff=0.450000	vEff=1.000000	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist= sequential

=================================
Store v1(top2000):
Avg key efficiency: 0.993854, Avg value efficiency: 0.997917, Avg peak keys: 301.716667, Max peak keys: 1000
Bad cases:
kEff=0.000000	vEff=0.000000	kPeak=   400	range=   400	keys=   400	keyDist=    hotspot	valueDist=   constant

=================================
Store v2(top200):
Avg key efficiency: 0.737813, Avg value efficiency: 0.776813, Avg peak keys: 148.679167, Max peak keys: 200
Bad cases:
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   200	keyDist=    zipfian	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   200	keyDist= sequential	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   400	keyDist=    hotspot	valueDist=   constant
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   400	keyDist=    hotspot	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   800	keyDist=    zipfian	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   800	keyDist=    zipfian	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=   400	keys=   800	keyDist= sequential	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   100	keyDist=    uniform	valueDist=   constant
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   100	keyDist= sequential	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    zipfian	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    zipfian	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    zipfian	valueDist=    hotspot
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    zipfian	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    hotspot	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    hotspot	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    hotspot	valueDist=    hotspot
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist=    hotspot	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=   500	keyDist= sequential	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    zipfian	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    zipfian	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    zipfian	valueDist=    hotspot
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    zipfian	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist=    hotspot	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    zipfian	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    zipfian	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    zipfian	valueDist=    hotspot
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    zipfian	valueDist= sequential
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=    uniform
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=    zipfian
kEff=0.000000	vEff=0.000000	kPeak=   200	range=  1000	keys=  2000	keyDist= sequential	valueDist= sequential
kEff=0.050000	vEff=1.000000	kPeak=   200	range=   201	keys=   201	keyDist=    uniform	valueDist=   constant
kEff=0.050000	vEff=0.583333	kPeak=   200	range=   400	keys=    40	keyDist=    uniform	valueDist=   constant
kEff=0.050000	vEff=0.111416	kPeak=   200	range=   400	keys=   400	keyDist=    zipfian	valueDist= sequential
kEff=0.050000	vEff=0.059310	kPeak=   200	range=   400	keys=   400	keyDist=    hotspot	valueDist= sequential
kEff=0.050000	vEff=0.225704	kPeak=   200	range=   400	keys=   800	keyDist=    uniform	valueDist= sequential
kEff=0.050000	vEff=0.072589	kPeak=   200	range=   400	keys=   800	keyDist=    zipfian	valueDist=    uniform
kEff=0.050000	vEff=0.264990	kPeak=   200	range=   400	keys=   800	keyDist=    hotspot	valueDist=   constant
kEff=0.050000	vEff=0.077194	kPeak=   200	range=   400	keys=   800	keyDist=    hotspot	valueDist=    uniform
kEff=0.050000	vEff=0.044921	kPeak=   200	range=   400	keys=   800	keyDist=    hotspot	valueDist= sequential
kEff=0.050000	vEff=0.426087	kPeak=   200	range=   400	keys=   800	keyDist= sequential	valueDist=    uniform
kEff=0.050000	vEff=0.718944	kPeak=   200	range=  1000	keys=   100	keyDist=    uniform	valueDist=    uniform
kEff=0.050000	vEff=0.694779	kPeak=   200	range=  1000	keys=   100	keyDist=    uniform	valueDist= sequential
kEff=0.050000	vEff=0.171843	kPeak=   200	range=  1000	keys=   100	keyDist=    zipfian	valueDist=    uniform
kEff=0.050000	vEff=0.234266	kPeak=   200	range=  1000	keys=   100	keyDist=    zipfian	valueDist=    zipfian
kEff=0.050000	vEff=0.112684	kPeak=   200	range=  1000	keys=   100	keyDist=    zipfian	valueDist= sequential
kEff=0.050000	vEff=0.168428	kPeak=   200	range=  1000	keys=   100	keyDist=    hotspot	valueDist=    uniform
kEff=0.050000	vEff=0.066667	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist=   constant
kEff=0.050000	vEff=0.254530	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist=    uniform
kEff=0.050000	vEff=0.088235	kPeak=   200	range=  1000	keys=   500	keyDist=    hotspot	valueDist=   constant
kEff=0.050000	vEff=0.185445	kPeak=   200	range=  1000	keys=   500	keyDist= sequential	valueDist=    uniform
kEff=0.050000	vEff=0.038462	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist=   constant
kEff=0.050000	vEff=0.191925	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    hotspot
kEff=0.050000	vEff=0.045853	kPeak=   200	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=    hotspot
kEff=0.050000	vEff=0.127632	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist=    hotspot
kEff=0.050000	vEff=0.081081	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist=   constant
kEff=0.050000	vEff=0.149967	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    uniform
kEff=0.050000	vEff=0.101270	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    hotspot
kEff=0.050000	vEff=0.023166	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=   constant
kEff=0.050000	vEff=0.031977	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=    hotspot
kEff=0.050000	vEff=0.037666	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist= sequential
kEff=0.050000	vEff=0.131252	kPeak=   200	range=  1000	keys=  2000	keyDist= sequential	valueDist=    uniform
kEff=0.050000	vEff=0.156375	kPeak=   200	range=  1000	keys=  2000	keyDist= sequential	valueDist=    hotspot
kEff=0.100000	vEff=0.345679	kPeak=   200	range=   400	keys=   200	keyDist=    uniform	valueDist=   constant
kEff=0.100000	vEff=0.412095	kPeak=   200	range=   400	keys=   200	keyDist=    zipfian	valueDist=    hotspot
kEff=0.100000	vEff=0.217724	kPeak=   200	range=   400	keys=   200	keyDist=    hotspot	valueDist=    uniform
kEff=0.100000	vEff=0.369248	kPeak=   200	range=   400	keys=   400	keyDist=    uniform	valueDist= sequential
kEff=0.100000	vEff=0.102228	kPeak=   200	range=   400	keys=   400	keyDist=    zipfian	valueDist=    uniform
kEff=0.100000	vEff=0.184978	kPeak=   200	range=   400	keys=   400	keyDist=    zipfian	valueDist=    zipfian
kEff=0.100000	vEff=0.286326	kPeak=   200	range=   400	keys=   400	keyDist= sequential	valueDist= sequential
kEff=0.100000	vEff=0.155235	kPeak=   200	range=   400	keys=   800	keyDist=    uniform	valueDist=   constant
kEff=0.100000	vEff=0.497461	kPeak=   200	range=   400	keys=   800	keyDist=    uniform	valueDist=    uniform
kEff=0.100000	vEff=0.038888	kPeak=   200	range=   400	keys=   800	keyDist=    zipfian	valueDist=    hotspot
kEff=0.100000	vEff=0.095131	kPeak=   200	range=   400	keys=   800	keyDist=    hotspot	valueDist=    zipfian
kEff=0.100000	vEff=0.616533	kPeak=   200	range=  1000	keys=   100	keyDist=    hotspot	valueDist=    zipfian
kEff=0.100000	vEff=0.322270	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist=    zipfian
kEff=0.100000	vEff=0.041748	kPeak=   200	range=  1000	keys=   500	keyDist=    zipfian	valueDist=   constant
kEff=0.100000	vEff=0.394694	kPeak=   200	range=  1000	keys=   500	keyDist= sequential	valueDist=    hotspot
kEff=0.100000	vEff=0.485929	kPeak=   200	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=   constant
kEff=0.100000	vEff=0.136938	kPeak=   200	range=  1000	keys=  1000	keyDist= sequential	valueDist=    zipfian
kEff=0.100000	vEff=0.071032	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist= sequential
kEff=0.100000	vEff=0.186894	kPeak=   200	range=  1000	keys=  2000	keyDist= sequential	valueDist=    zipfian
kEff=0.150000	vEff=0.638761	kPeak=   200	range=   400	keys=   200	keyDist=    uniform	valueDist= sequential
kEff=0.150000	vEff=0.079143	kPeak=   200	range=   400	keys=   200	keyDist=    zipfian	valueDist= sequential
kEff=0.150000	vEff=0.089301	kPeak=   200	range=   400	keys=   200	keyDist=    hotspot	valueDist= sequential
kEff=0.150000	vEff=0.036075	kPeak=   200	range=   400	keys=   400	keyDist=    zipfian	valueDist=    hotspot
kEff=0.150000	vEff=0.173189	kPeak=   200	range=   400	keys=   400	keyDist=    hotspot	valueDist=    zipfian
kEff=0.150000	vEff=0.216656	kPeak=   200	range=   400	keys=   400	keyDist= sequential	valueDist=    uniform
kEff=0.150000	vEff=0.174728	kPeak=   200	range=  1000	keys=   100	keyDist=    hotspot	valueDist= sequential
kEff=0.150000	vEff=0.316652	kPeak=   200	range=  1000	keys=   500	keyDist=    uniform	valueDist=    hotspot
kEff=0.150000	vEff=0.203996	kPeak=   200	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    zipfian
kEff=0.200000	vEff=0.113208	kPeak=   200	range=   400	keys=   200	keyDist=    hotspot	valueDist=   constant
kEff=0.200000	vEff=0.577185	kPeak=   200	range=   400	keys=   400	keyDist=    uniform	valueDist=    uniform
kEff=0.200000	vEff=0.484167	kPeak=   200	range=   400	keys=   800	keyDist=    uniform	valueDist=    zipfian
kEff=0.200000	vEff=0.510552	kPeak=   200	range=   400	keys=   800	keyDist=    uniform	valueDist=    hotspot
kEff=0.250000	vEff=0.721473	kPeak=   200	range=   201	keys=   201	keyDist=    zipfian	valueDist=   constant
kEff=0.250000	vEff=0.593323	kPeak=   200	range=   400	keys=   200	keyDist=    uniform	valueDist=    uniform
kEff=0.250000	vEff=0.062349	kPeak=   200	range=   400	keys=   200	keyDist=    zipfian	valueDist=   constant
kEff=0.250000	vEff=0.185172	kPeak=   200	range=   400	keys=   200	keyDist=    zipfian	valueDist=    zipfian
kEff=0.250000	vEff=0.166227	kPeak=   200	range=   400	keys=   400	keyDist=    uniform	valueDist=   constant
kEff=0.250000	vEff=0.712312	kPeak=   200	range=   400	keys=   400	keyDist=    uniform	valueDist=    zipfian
kEff=0.250000	vEff=0.029004	kPeak=   200	range=   400	keys=   800	keyDist=    zipfian	valueDist=   constant
kEff=0.250000	vEff=0.284395	kPeak=   200	range=  1000	keys=   100	keyDist=    zipfian	valueDist=    hotspot
kEff=0.250000	vEff=0.171779	kPeak=   200	range=  1000	keys=   100	keyDist=    hotspot	valueDist=   constant
kEff=0.250000	vEff=0.584071	kPeak=   200	range=  1000	keys=   100	keyDist=    hotspot	valueDist=    hotspot
kEff=0.250000	vEff=0.187006	kPeak=   200	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    zipfian
kEff=0.300000	vEff=0.358668	kPeak=   200	range=   201	keys=   201	keyDist=    hotspot	valueDist=   constant
kEff=0.300000	vEff=0.503647	kPeak=   200	range=   201	keys=   402	keyDist=    hotspot	valueDist=   constant
kEff=0.300000	vEff=0.263817	kPeak=   200	range=   400	keys=   400	keyDist=    hotspot	valueDist=    hotspot
kEff=0.300000	vEff=0.748810	kPeak=   200	range=   400	keys=   400	keyDist= sequential	valueDist=    zipfian
kEff=0.300000	vEff=0.598304	kPeak=   200	range=   400	keys=   400	keyDist= sequential	valueDist=    hotspot
kEff=0.300000	vEff=0.136964	kPeak=   200	range=   400	keys=   800	keyDist=    hotspot	valueDist=    hotspot
kEff=0.300000	vEff=0.405484	kPeak=   200	range=   400	keys=   800	keyDist= sequential	valueDist=    zipfian
kEff=0.300000	vEff=0.762834	kPeak=   200	range=  1000	keys=   100	keyDist=    uniform	valueDist=    zipfian
kEff=0.300000	vEff=0.784295	kPeak=   200	range=  1000	keys=   100	keyDist=    uniform	valueDist=    hotspot
kEff=0.300000	vEff=0.335690	kPeak=   200	range=  1000	keys=   500	keyDist= sequential	valueDist=    zipfian
kEff=0.350000	vEff=1.000000	kPeak=   200	range=   201	keys=   100	keyDist=    hotspot	valueDist=   constant
kEff=0.350000	vEff=0.159167	kPeak=   200	range=   201	keys=   402	keyDist=    zipfian	valueDist=   constant
kEff=0.350000	vEff=0.919900	kPeak=   200	range=   400	keys=    40	keyDist=    uniform	valueDist= sequential
kEff=0.350000	vEff=0.872637	kPeak=   200	range=   400	keys=   200	keyDist=    uniform	valueDist=    zipfian
kEff=0.350000	vEff=0.511027	kPeak=   200	range=   400	keys=   800	keyDist= sequential	valueDist=    hotspot
kEff=0.400000	vEff=0.940239	kPeak=   200	range=   400	keys=    40	keyDist=    hotspot	valueDist=   constant
kEff=0.400000	vEff=0.574748	kPeak=   200	range=   400	keys=   200	keyDist= sequential	valueDist=    uniform
kEff=0.400000	vEff=0.748881	kPeak=   200	range=  1000	keys=   100	keyDist= sequential	valueDist=    uniform
kEff=0.450000	vEff=0.877375	kPeak=   200	range=   400	keys=    40	keyDist=    uniform	valueDist=    uniform
kEff=0.450000	vEff=0.344519	kPeak=   200	range=   400	keys=   200	keyDist=    hotspot	valueDist=    zipfian
kEff=0.450000	vEff=0.032364	kPeak=   200	range=   400	keys=   400	keyDist=    zipfian	valueDist=   constant
kEff=0.450000	vEff=0.907319	kPeak=   200	range=  1000	keys=   100	keyDist= sequential	valueDist=    hotspot
kEff=0.600000	vEff=0.157219	kPeak=   200	range=   400	keys=    40	keyDist=    zipfian	valueDist=   constant
kEff=0.600000	vEff=0.315463	kPeak=   200	range=   400	keys=   200	keyDist=    hotspot	valueDist=    hotspot
kEff=0.600000	vEff=0.057496	kPeak=   200	range=  1000	keys=  2000	keyDist=    zipfian	valueDist=   constant
kEff=0.750000	vEff=0.420935	kPeak=   200	range=  1000	keys=   100	keyDist=    zipfian	valueDist=   constant
kEff=0.750000	vEff=0.123973	kPeak=   200	range=  1000	keys=  1000	keyDist=    zipfian	valueDist=   constant

=================================
Store v3(top20):
Avg key efficiency: 0.793333, Avg value efficiency: 0.774041, Avg peak keys: 184.781250, Max peak keys: 733
Bad cases:
kEff=0.000000	vEff=0.000000	kPeak=   351	range=   400	keys=   400	keyDist=    hotspot	valueDist=   constant
kEff=0.000000	vEff=0.000000	kPeak=   641	range=  1000	keys=  2000	keyDist= sequential	valueDist= sequential
kEff=0.050000	vEff=0.199798	kPeak=   714	range=  1000	keys=   500	keyDist= sequential	valueDist=    uniform
kEff=0.100000	vEff=0.305668	kPeak=   711	range=  1000	keys=  1000	keyDist= sequential	valueDist=    hotspot
kEff=0.100000	vEff=0.094286	kPeak=   695	range=  1000	keys=  1000	keyDist= sequential	valueDist= sequential
kEff=0.100000	vEff=0.235951	kPeak=   702	range=  1000	keys=  2000	keyDist= sequential	valueDist=    zipfian
kEff=0.150000	vEff=0.369880	kPeak=   384	range=   400	keys=   200	keyDist= sequential	valueDist=    uniform
kEff=0.150000	vEff=0.202278	kPeak=   383	range=   400	keys=   400	keyDist= sequential	valueDist=    uniform
kEff=0.150000	vEff=0.123222	kPeak=   400	range=   400	keys=   800	keyDist= sequential	valueDist= sequential
kEff=0.150000	vEff=0.248980	kPeak=   695	range=  1000	keys=  1000	keyDist=    uniform	valueDist=   constant
kEff=0.150000	vEff=0.132515	kPeak=   708	range=  1000	keys=  1000	keyDist= sequential	valueDist=    uniform
kEff=0.150000	vEff=0.129455	kPeak=   698	range=  1000	keys=  2000	keyDist= sequential	valueDist=    uniform
kEff=0.150000	vEff=0.203089	kPeak=   695	range=  1000	keys=  2000	keyDist= sequential	valueDist=    hotspot
kEff=0.200000	vEff=0.447384	kPeak=   155	range=   200	keys=   100	keyDist= sequential	valueDist= sequential
kEff=0.200000	vEff=0.302582	kPeak=   200	range=   200	keys=   200	keyDist= sequential	valueDist= sequential
kEff=0.200000	vEff=0.244453	kPeak=   200	range=   200	keys=   400	keyDist= sequential	valueDist= sequential
kEff=0.200000	vEff=0.234497	kPeak=   137	range=   201	keys=   201	keyDist= sequential	valueDist= sequential
kEff=0.200000	vEff=0.466667	kPeak=   699	range=  1000	keys=   100	keyDist= sequential	valueDist= sequential
kEff=0.200000	vEff=0.293194	kPeak=   701	range=  1000	keys=   500	keyDist=    uniform	valueDist=   constant
kEff=0.200000	vEff=0.436215	kPeak=   725	range=  1000	keys=   500	keyDist= sequential	valueDist=    zipfian
kEff=0.200000	vEff=0.228040	kPeak=   703	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    uniform
kEff=0.250000	vEff=0.301657	kPeak=   201	range=   201	keys=   201	keyDist= sequential	valueDist=    uniform
kEff=0.250000	vEff=0.290278	kPeak=   379	range=   400	keys=   800	keyDist=    uniform	valueDist=   constant
kEff=0.250000	vEff=0.182968	kPeak=   382	range=   400	keys=   800	keyDist= sequential	valueDist=    uniform
kEff=0.250000	vEff=0.364516	kPeak=   707	range=  1000	keys=   500	keyDist=    uniform	valueDist= sequential
kEff=0.250000	vEff=0.183537	kPeak=   598	range=  1000	keys=   500	keyDist= sequential	valueDist= sequential
kEff=0.250000	vEff=0.285188	kPeak=   700	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    uniform
kEff=0.250000	vEff=0.418846	kPeak=   718	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    hotspot
kEff=0.300000	vEff=0.291112	kPeak=   201	range=   201	keys=   402	keyDist= sequential	valueDist= sequential
kEff=0.300000	vEff=0.331395	kPeak=   400	range=   400	keys=   200	keyDist= sequential	valueDist= sequential
kEff=0.300000	vEff=0.293849	kPeak=   382	range=   400	keys=   800	keyDist=    uniform	valueDist=    uniform
kEff=0.300000	vEff=0.363330	kPeak=   703	range=  1000	keys=   500	keyDist=    uniform	valueDist=    uniform
kEff=0.300000	vEff=0.485338	kPeak=   703	range=  1000	keys=   500	keyDist=    uniform	valueDist=    zipfian
kEff=0.300000	vEff=0.535824	kPeak=   703	range=  1000	keys=   500	keyDist= sequential	valueDist=    hotspot
kEff=0.300000	vEff=0.300642	kPeak=   685	range=  1000	keys=  1000	keyDist=    uniform	valueDist= sequential
kEff=0.300000	vEff=0.293702	kPeak=   709	range=  1000	keys=  1000	keyDist= sequential	valueDist=    zipfian
kEff=0.300000	vEff=0.167035	kPeak=   700	range=  1000	keys=  2000	keyDist=    uniform	valueDist=   constant
kEff=0.300000	vEff=0.310325	kPeak=   718	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    zipfian
kEff=0.300000	vEff=0.212935	kPeak=   693	range=  1000	keys=  2000	keyDist=    uniform	valueDist= sequential
kEff=0.350000	vEff=0.392471	kPeak=   150	range=   150	keys=   150	keyDist= sequential	valueDist= sequential
kEff=0.350000	vEff=0.225180	kPeak=   150	range=   150	keys=   300	keyDist= sequential	valueDist= sequential
kEff=0.350000	vEff=0.414622	kPeak=   200	range=   200	keys=   400	keyDist=    uniform	valueDist=    uniform
kEff=0.350000	vEff=0.280271	kPeak=   200	range=   200	keys=   400	keyDist= sequential	valueDist=    uniform
kEff=0.350000	vEff=0.740806	kPeak=    98	range=   201	keys=   100	keyDist= sequential	valueDist= sequential
kEff=0.350000	vEff=0.273638	kPeak=   201	range=   201	keys=   402	keyDist= sequential	valueDist=    uniform
kEff=0.350000	vEff=0.464407	kPeak=   385	range=   400	keys=   200	keyDist=    uniform	valueDist=   constant
kEff=0.350000	vEff=0.138924	kPeak=   400	range=   400	keys=   400	keyDist= sequential	valueDist= sequential
kEff=0.350000	vEff=0.407933	kPeak=   381	range=   400	keys=   800	keyDist=    uniform	valueDist=    zipfian
kEff=0.350000	vEff=0.369613	kPeak=   377	range=   400	keys=   800	keyDist=    uniform	valueDist=    hotspot
kEff=0.350000	vEff=0.305003	kPeak=   383	range=   400	keys=   800	keyDist= sequential	valueDist=    hotspot
kEff=0.350000	vEff=0.291951	kPeak=   199	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=    uniform
kEff=0.350000	vEff=0.320031	kPeak=   199	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=    hotspot
kEff=0.350000	vEff=0.281124	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist= sequential
kEff=0.400000	vEff=0.591623	kPeak=   150	range=   150	keys=    75	keyDist= sequential	valueDist=    uniform
kEff=0.400000	vEff=0.376506	kPeak=   201	range=   201	keys=   402	keyDist=    uniform	valueDist=    uniform
kEff=0.400000	vEff=0.482516	kPeak=   379	range=   400	keys=   200	keyDist=    uniform	valueDist= sequential
kEff=0.400000	vEff=0.353719	kPeak=   382	range=   400	keys=   400	keyDist=    uniform	valueDist=   constant
kEff=0.400000	vEff=0.498799	kPeak=   390	range=   400	keys=   400	keyDist=    uniform	valueDist=    hotspot
kEff=0.400000	vEff=0.374276	kPeak=   384	range=   400	keys=   400	keyDist=    uniform	valueDist= sequential
kEff=0.400000	vEff=0.435010	kPeak=   375	range=   400	keys=   400	keyDist= sequential	valueDist=    zipfian
kEff=0.400000	vEff=0.350625	kPeak=   381	range=   400	keys=   800	keyDist= sequential	valueDist=    zipfian
kEff=0.400000	vEff=0.725519	kPeak=   714	range=  1000	keys=   100	keyDist=    uniform	valueDist=    uniform
kEff=0.400000	vEff=0.329235	kPeak=   199	range=  1000	keys=  1000	keyDist=    hotspot	valueDist= sequential
kEff=0.400000	vEff=0.277056	kPeak=   706	range=  1000	keys=  2000	keyDist=    uniform	valueDist=    hotspot
kEff=0.450000	vEff=0.705110	kPeak=   150	range=   150	keys=    75	keyDist= sequential	valueDist= sequential
kEff=0.450000	vEff=0.350837	kPeak=   150	range=   150	keys=   150	keyDist= sequential	valueDist=    uniform
kEff=0.450000	vEff=0.290745	kPeak=   200	range=   200	keys=   200	keyDist= sequential	valueDist=    uniform
kEff=0.450000	vEff=0.368885	kPeak=   198	range=   200	keys=   400	keyDist=    uniform	valueDist=   constant
kEff=0.450000	vEff=0.424439	kPeak=   200	range=   200	keys=   400	keyDist= sequential	valueDist=    zipfian
kEff=0.450000	vEff=0.502934	kPeak=   199	range=   201	keys=   201	keyDist=    uniform	valueDist= sequential
kEff=0.450000	vEff=0.405721	kPeak=   201	range=   201	keys=   402	keyDist=    uniform	valueDist= sequential
kEff=0.450000	vEff=0.449621	kPeak=   201	range=   201	keys=   402	keyDist= sequential	valueDist=    zipfian
kEff=0.450000	vEff=1.000000	kPeak=   279	range=   400	keys=    40	keyDist= sequential	valueDist= sequential
kEff=0.450000	vEff=0.472608	kPeak=   381	range=   400	keys=   200	keyDist=    uniform	valueDist=    uniform
kEff=0.450000	vEff=0.606472	kPeak=   382	range=   400	keys=   200	keyDist= sequential	valueDist=    hotspot
kEff=0.450000	vEff=0.375035	kPeak=   378	range=   400	keys=   400	keyDist=    uniform	valueDist=    uniform
kEff=0.450000	vEff=0.413645	kPeak=   384	range=   400	keys=   400	keyDist= sequential	valueDist=    hotspot
kEff=0.450000	vEff=0.296609	kPeak=   382	range=   400	keys=   800	keyDist=    uniform	valueDist= sequential
kEff=0.450000	vEff=0.475806	kPeak=   710	range=  1000	keys=   100	keyDist=    uniform	valueDist=   constant
kEff=0.450000	vEff=0.665877	kPeak=   714	range=  1000	keys=   100	keyDist=    uniform	valueDist= sequential
kEff=0.450000	vEff=0.756776	kPeak=   733	range=  1000	keys=   100	keyDist= sequential	valueDist=    uniform
kEff=0.450000	vEff=0.498698	kPeak=   703	range=  1000	keys=   500	keyDist=    uniform	valueDist=    hotspot
kEff=0.450000	vEff=0.393844	kPeak=   201	range=  1000	keys=   500	keyDist=    hotspot	valueDist=    uniform
kEff=0.500000	vEff=0.336543	kPeak=   150	range=   150	keys=   300	keyDist= sequential	valueDist=    uniform
kEff=0.500000	vEff=0.422667	kPeak=   200	range=   200	keys=   200	keyDist=    uniform	valueDist=   constant
kEff=0.500000	vEff=0.498771	kPeak=   200	range=   201	keys=   100	keyDist= sequential	valueDist=    uniform
kEff=0.500000	vEff=0.431034	kPeak=   201	range=   201	keys=   201	keyDist=    uniform	valueDist=   constant
kEff=0.500000	vEff=0.498706	kPeak=   201	range=   201	keys=   201	keyDist=    uniform	valueDist=    uniform
kEff=0.500000	vEff=0.334055	kPeak=   201	range=   201	keys=   402	keyDist=    uniform	valueDist=   constant
kEff=0.500000	vEff=0.371439	kPeak=   200	range=  1000	keys=   500	keyDist=    hotspot	valueDist= sequential
kEff=0.500000	vEff=0.285769	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=   constant
kEff=0.500000	vEff=0.358317	kPeak=   200	range=  1000	keys=  2000	keyDist=    hotspot	valueDist=    zipfian
kEff=0.550000	vEff=0.404339	kPeak=   150	range=   150	keys=   300	keyDist=    uniform	valueDist=   constant
kEff=0.550000	vEff=0.472583	kPeak=   150	range=   150	keys=   300	keyDist=    uniform	valueDist= sequential
kEff=0.550000	vEff=0.447923	kPeak=   200	range=   200	keys=   400	keyDist= sequential	valueDist=    hotspot
kEff=0.550000	vEff=0.484526	kPeak=   377	range=   400	keys=   400	keyDist=    uniform	valueDist=    zipfian
kEff=0.550000	vEff=0.456791	kPeak=    80	range=   400	keys=   800	keyDist=    hotspot	valueDist= sequential
kEff=0.550000	vEff=0.334416	kPeak=   201	range=  1000	keys=   500	keyDist=    hotspot	valueDist=   constant
kEff=0.550000	vEff=0.480051	kPeak=   256	range=  1000	keys=   500	keyDist=    hotspot	valueDist=    zipfian
kEff=0.550000	vEff=0.443359	kPeak=   713	range=  1000	keys=  1000	keyDist=    uniform	valueDist=    zipfian
kEff=0.550000	vEff=0.318152	kPeak=   199	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=   constant
kEff=0.600000	vEff=0.483835	kPeak=   150	range=   150	keys=   150	keyDist=    uniform	valueDist=   constant
kEff=0.600000	vEff=0.420376	kPeak=   199	range=   200	keys=   400	keyDist=    uniform	valueDist= sequential
kEff=0.600000	vEff=0.497032	kPeak=   201	range=   201	keys=   402	keyDist=    uniform	valueDist=    zipfian
kEff=0.600000	vEff=0.486919	kPeak=   201	range=   201	keys=   402	keyDist=    uniform	valueDist=    hotspot
kEff=0.600000	vEff=0.458751	kPeak=   230	range=  1000	keys=   500	keyDist=    hotspot	valueDist=    hotspot
kEff=0.600000	vEff=0.340838	kPeak=   200	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=    uniform
kEff=0.600000	vEff=0.394172	kPeak=   203	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=    hotspot
kEff=0.650000	vEff=0.462416	kPeak=   150	range=   150	keys=   300	keyDist=    uniform	valueDist=    uniform
kEff=0.650000	vEff=0.483024	kPeak=   150	range=   150	keys=   300	keyDist= sequential	valueDist=    hotspot
kEff=0.650000	vEff=0.497570	kPeak=   200	range=   200	keys=   200	keyDist=    uniform	valueDist= sequential
kEff=0.650000	vEff=0.410307	kPeak=   201	range=   201	keys=   402	keyDist= sequential	valueDist=    hotspot
kEff=0.700000	vEff=0.481252	kPeak=   200	range=   200	keys=   400	keyDist=    uniform	valueDist=    hotspot
kEff=0.700000	vEff=0.409335	kPeak=   206	range=  1000	keys=  1000	keyDist=    hotspot	valueDist=    zipfian
kEff=0.750000	vEff=0.466284	kPeak=    80	range=   400	keys=   800	keyDist=    hotspot	valueDist=    uniform
kEff=0.800000	vEff=0.478572	kPeak=    80	range=   400	keys=   800	keyDist=    hotspot	valueDist=   constant

=================================
Store v3(top200):
Avg key efficiency: 0.961979, Avg value efficiency: 0.976873, Avg peak keys: 297.760417, Max peak keys: 1000
Bad cases:
kEff=0.000000	vEff=0.000000	kPeak=  1000	range=  1000	keys=  2000	keyDist= sequential	valueDist= sequential
kEff=0.200000	vEff=0.000000	kPeak=   400	range=   400	keys=   400	keyDist=    hotspot	valueDist=   constant
kEff=0.200000	vEff=0.550576	kPeak=  1000	range=  1000	keys=  1000	keyDist= sequential	valueDist=    uniform
kEff=0.300000	vEff=0.641137	kPeak=  1000	range=  1000	keys=   500	keyDist= sequential	valueDist= sequential
kEff=0.300000	vEff=0.377778	kPeak=  1000	range=  1000	keys=  1000	keyDist= sequential	valueDist= sequential
kEff=0.350000	vEff=0.750716	kPeak=   400	range=   400	keys=   800	keyDist= sequential	valueDist= sequential
kEff=0.350000	vEff=0.570600	kPeak=  1000	range=  1000	keys=  2000	keyDist=    uniform	valueDist=   constant
kEff=0.450000	vEff=0.467265	kPeak=  1000	range=  1000	keys=  2000	keyDist= sequential	valueDist=    uniform
```