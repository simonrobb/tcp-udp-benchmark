# TCP/UDP comparison benchmark

## Run benchmark

Ensure you have Golang (`brew install golang`) and netcat (`nc`) installed (`brew install nc`).

Run the following to start UDP and TCP servers with netcat, and run the benchmarks:

```
make -j 3 start-udp-server start-tcp-server benchmark
```

## Results

UDP is about 41% faster than TCP on my 2017 Macbook.

```
goos: darwin
goarch: amd64
BenchmarkSpamTCP-8   	1000000000	         0.489 ns/op	       0 B/op	       0 allocs/op
BenchmarkSpamUDP-8   	1000000000	         0.289 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	_/Users/simonrobb/code/clay/tcp-udp-benchmark	16.328s
```