# zzrand

Maybe a faster pseudorandom number generator.

# Inspired by

* [fastrand](https://github.com/valyala/fastrand)
* [wyhash](https://github.com/wangyi-fudan/wyhash)

# Work

Based on fastrand, replaced xorshift with wyhash. Sometimes faster than `fastrand`

Add more functions, more convenient


# Benchmark

A little faster than `fastrand` in parallel call (`SetParallelism(10)`).

```
Benchmark_ZzrandUint32P
Benchmark_ZzrandUint32P-8     	516694711	         2.112 ns/op
Benchmark_FastrandUint32P
Benchmark_FastrandUint32P-8   	531871572	         2.190 ns/op
Benchmark_StdrandUint32P
Benchmark_StdrandUint32P-8    	 9634488	       122.9 ns/op
Benchmark_ZzrandIntP
Benchmark_ZzrandIntP-8        	497771090	         2.097 ns/op
Benchmark_FastrandIntP
Benchmark_FastrandIntP-8      	618760429	         2.157 ns/op
Benchmark_StdrandIntP
Benchmark_StdrandIntP-8       	 9904023	       125.1 ns/op
Benchmark_ZzrandIntnP
Benchmark_ZzrandIntnP-8       	545704658	         2.152 ns/op
Benchmark_FastrandIntnP
Benchmark_FastrandIntnP-8     	604431744	         2.126 ns/op
Benchmark_StdrandIntnP
Benchmark_StdrandIntnP-8      	11444502	       105.5 ns/op
Benchmark_ZzrandUint32
Benchmark_ZzrandUint32-8      	141396771	         8.469 ns/op
Benchmark_FastrandUint32
Benchmark_FastrandUint32-8    	146265015	         8.199 ns/op
Benchmark_ZzrandInt
Benchmark_ZzrandInt-8         	141230719	         8.461 ns/op
Benchmark_FastrandInt
Benchmark_FastrandInt-8       	143432836	         8.177 ns/op
Benchmark_ZzrandIntn
Benchmark_ZzrandIntn-8        	141461223	         8.469 ns/op
Benchmark_FastrandIntn
Benchmark_FastrandIntn-8      	146308654	         8.171 ns/op
```