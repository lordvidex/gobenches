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
❯ go test -bench=. ./chan_slices -benchmem -benchtime=1s
goos: darwin
goarch: arm64
pkg: github.com/lordvidex/gobenches/chan_slices
cpu: Apple M2 Pro
BenchmarkLargeStruct/batch=1-10                                4         257362042 ns/op        153610428 B/op    200029 allocs/op
BenchmarkLargeStruct/batch=1-stream-10                         4         301234448 ns/op        153608392 B/op    200023 allocs/op
BenchmarkLargeStruct/batch=8-10                                5         217946892 ns/op        153607480 B/op     25023 allocs/op
BenchmarkLargeStruct/batch=8-stream-10                         5         217961650 ns/op        153608670 B/op     25029 allocs/op
BenchmarkLargeStruct/batch=1024-10                             6         167973125 ns/op        153606216 B/op       207 allocs/op
BenchmarkLargeStruct/batch=1024-stream-10                      6         171586618 ns/op        153607160 B/op       207 allocs/op
BenchmarkLargeStructPtr/batch=1-10                             4         256205812 ns/op        155206288 B/op    400012 allocs/op
BenchmarkLargeStructPtr/batch=1-stream-10                      3         385123597 ns/op        155206138 B/op    400009 allocs/op
BenchmarkLargeStructPtr/batch=8-10                             5         214235033 ns/op        155207800 B/op    225028 allocs/op
BenchmarkLargeStructPtr/batch=8-stream-10                      5         214364133 ns/op        155207939 B/op    225029 allocs/op
BenchmarkLargeStructPtr/batch=1024-10                          7         161950869 ns/op        155455662 B/op    200204 allocs/op
BenchmarkLargeStructPtr/batch=1024-stream-10                   6         174463542 ns/op        155457418 B/op    200212 allocs/op
BenchmarkSmallStruct/batch=1-10                               43          27645390 ns/op         4805621 B/op     200005 allocs/op
BenchmarkSmallStruct/batch=1-stream-10                        18          63736565 ns/op         4805831 B/op     200007 allocs/op
BenchmarkSmallStruct/batch=8-10                              187           6322351 ns/op         4805608 B/op      25005 allocs/op
BenchmarkSmallStruct/batch=8-stream-10                        36          30503064 ns/op         4805752 B/op      25007 allocs/op
BenchmarkSmallStruct/batch=1024-10                           508           2346407 ns/op         4806164 B/op        201 allocs/op
BenchmarkSmallStruct/batch=1024-stream-10                     55          22152431 ns/op         4806350 B/op        203 allocs/op
BenchmarkSmallStructPtr/batch=1-10                            36          29759909 ns/op         6405605 B/op     400005 allocs/op
BenchmarkSmallStructPtr/batch=1-stream-10                     16          65297263 ns/op         6405756 B/op     400008 allocs/op
BenchmarkSmallStructPtr/batch=8-10                           138           8677544 ns/op         6405635 B/op     225005 allocs/op
BenchmarkSmallStructPtr/batch=8-stream-10                     34          33578371 ns/op         6405813 B/op     225008 allocs/op
BenchmarkSmallStructPtr/batch=1024-10                        256           4658119 ns/op         6655356 B/op     200201 allocs/op
BenchmarkSmallStructPtr/batch=1024-stream-10                  54          21927065 ns/op         6655563 B/op     200204 allocs/op
PASS
ok      github.com/lordvidex/gobenches/chan_slices      27.307s
```
