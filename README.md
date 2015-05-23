Popcount
--------
The results:

	BenchmarkHamming32	300000000	         4.87 ns/op
	BenchmarkSlow32	cd go/100000000	        15.7 ns/op
	BenchmarkFast32	src/gith300000000	         4.84 ns/op
	BenchmarkPopCnt32	500000000	         3.14 ns/op
	BenchmarkFast64	300000000	         5.56 ns/op
	BenchmarkByteTable32	200000000	         6.61 ns/op
	BenchmarkByteTable64	100000000	        11.5 ns/op
	BenchmarkHamming64	300000000	         5.66 ns/op
	BenchmarkSlow64	50000000	        37.1 ns/op
