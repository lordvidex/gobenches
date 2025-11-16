# chan_slices

## Preconditions
- This is batch processing
- There exists a feed returning N items (in batch)
## Questions
- Is it more efficient memory/cpu to pass a slice of structs over a channel or passing each single struct over the channel?

## Cases
- channel with slice of structs (Big, Small)
- channel with slice of pointers to structs (Big, Small)
- channel with pointers 
- channel with structs (Big, Small)

## Results
```bash
❯ go test -bench=. ./chan_slices -benchmem
goos: darwin
goarch: arm64
pkg: github.com/lordvidex/gobenches/chan_slices
cpu: Apple M2 Pro
BenchmarkLargeStruct/batch=1-10                                8         141202682 ns/op             236 B/op          3 allocs/op
BenchmarkLargeStruct/batch=1-stream-10                         6         198144702 ns/op            3520 B/op          3 allocs/op
BenchmarkLargeStruct/batch=8-10                                9         118449773 ns/op             470 B/op          4 allocs/op
BenchmarkLargeStruct/batch=8-stream-10                         6         197813632 ns/op            3242 B/op          2 allocs/op
BenchmarkLargeStruct/batch=1024-10                            13          88625064 ns/op             667 B/op          4 allocs/op
BenchmarkLargeStruct/batch=1024-stream-10                      6         197523722 ns/op            3232 B/op          2 allocs/op
BenchmarkLargeStructPtr/batch=1-10                             7         145424833 ns/op             245 B/op          3 allocs/op
BenchmarkLargeStructPtr/batch=1-stream-10                      7         144574202 ns/op             152 B/op          3 allocs/op
BenchmarkLargeStructPtr/batch=8-10                             9         122698144 ns/op             217 B/op          3 allocs/op
BenchmarkLargeStructPtr/batch=8-stream-10                      8         141622490 ns/op             152 B/op          3 allocs/op
BenchmarkLargeStructPtr/batch=1024-10                         13          88601971 ns/op             502 B/op          4 allocs/op
BenchmarkLargeStructPtr/batch=1024-stream-10                   7         144386393 ns/op             152 B/op          3 allocs/op
BenchmarkSmallStruct/batch=1-10                               51          21594809 ns/op             198 B/op          3 allocs/op
BenchmarkSmallStruct/batch=1-stream-10                        56          21277289 ns/op             160 B/op          2 allocs/op
BenchmarkSmallStruct/batch=8-10                              405           2932624 ns/op             171 B/op          3 allocs/op
BenchmarkSmallStruct/batch=8-stream-10                        56          21380825 ns/op             171 B/op          2 allocs/op
BenchmarkSmallStruct/batch=1024-10                          4774            263438 ns/op             172 B/op          3 allocs/op
BenchmarkSmallStruct/batch=1024-stream-10                     60          21258601 ns/op             160 B/op          2 allocs/op
BenchmarkSmallStructPtr/batch=1-10                            56          21597937 ns/op             168 B/op          3 allocs/op
BenchmarkSmallStructPtr/batch=1-stream-10                     56          21359793 ns/op             152 B/op          3 allocs/op
BenchmarkSmallStructPtr/batch=8-10                           402           2996583 ns/op             168 B/op          3 allocs/op
BenchmarkSmallStructPtr/batch=8-stream-10                     57          21254529 ns/op             152 B/op          3 allocs/op
BenchmarkSmallStructPtr/batch=1024-10                       4095            295485 ns/op             170 B/op          3 allocs/op
BenchmarkSmallStructPtr/batch=1024-stream-10                  58          21473529 ns/op             152 B/op          3 allocs/op
PASS
ok      github.com/lordvidex/gobenches/chan_slices      32.267s
```

## Conclusions
- It is more efficient to pass a slice of structs over a channel than passing each single struct over the channel.
```bash
# small
BenchmarkSmallStructPtr/batch=1024-10                       4095            295485 ns/op             170 B/op          3 allocs/op
BenchmarkSmallStructPtr/batch=1024-stream-10                  58          21473529 ns/op             152 B/op          3 allocs/op
# large - although the performance gains gets smaller the larger the struct size.
BenchmarkLargeStructPtr/batch=1024-10                         13          88601971 ns/op             502 B/op          4 allocs/op
BenchmarkLargeStructPtr/batch=1024-stream-10                   7         144386393 ns/op             152 B/op          3 allocs/op
```
- The larger the batch size, the better the gains
```bash
BenchmarkSmallStruct/batch=1-10                               51          21594809 ns/op             198 B/op          3 allocs/op
BenchmarkSmallStruct/batch=8-10                              405           2932624 ns/op             171 B/op          3 allocs/op
BenchmarkSmallStruct/batch=1024-10                          4774            263438 ns/op             172 B/op          3 allocs/op
```
- There is negligible performance difference between passing a slice of structs over passing a slice of pointer to structs over the channel.
```bash
# large 
BenchmarkLargeStruct/batch=1024-10                            13          88625064 ns/op             667 B/op          4 allocs/op
BenchmarkLargeStructPtr/batch=1024-10                         13          88601971 ns/op             502 B/op          4 allocs/op
```
- The memory allocation is typically the same in all scenarios. (In the scenario that the structs have been created already)

## Reviews
- Open a PR for any additional case.
