package chan_slices

import (
	"math/rand"
	"testing"
)

const (
	items      = 200_000
	batchOne   = 1
	batchSmall = 8
	batchBig   = 1024
	seed       = 1
)

var sink uint64

func BenchmarkLargeStruct(b *testing.B) {
	tests := []struct {
		name string
		gen  Generator[Big]
	}{
		{
			name: "batch=1",
			gen: Generator[Big]{
				total: items,
				batch: batchOne,
				factory: func(r *rand.Rand) Big {
					return NewBig(r)
				},
				seed: seed,
			},
		},
		{
			name: "batch=8",
			gen: Generator[Big]{
				total: items,
				batch: batchSmall,
				factory: func(r *rand.Rand) Big {
					return NewBig(r)
				},
				seed: seed,
			},
		},
		{
			name: "batch=1024",
			gen: Generator[Big]{
				total: items,
				batch: batchBig,
				factory: func(r *rand.Rand) Big {
					return NewBig(r)
				},
				seed: seed,
			},
		},
	}

	for _, tt := range tests {
		tt.gen.init()
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Batches() {
					for _, v := range x {
						acc += v.work()
					}
				}
				sink = acc
			}
		})
		b.Run(tt.name+"-stream", func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Stream() {
					acc += x.work()
				}
				sink = acc
			}
		})
	}
}

func BenchmarkLargeStructPtr(b *testing.B) {
	fac := func(r *rand.Rand) *Big {
		b := NewBig(r)
		return &b
	}
	tests := []struct {
		name string
		gen  Generator[*Big]
	}{
		{
			name: "batch=1",
			gen: Generator[*Big]{
				total:   items,
				batch:   batchOne,
				factory: fac,
				seed:    seed,
			},
		},
		{
			name: "batch=8",
			gen: Generator[*Big]{
				total:   items,
				batch:   batchSmall,
				factory: fac,
				seed:    seed,
			},
		},
		{
			name: "batch=1024",
			gen: Generator[*Big]{
				total:   items,
				batch:   batchBig,
				factory: fac,
				seed:    seed,
			},
		},
	}

	for _, tt := range tests {
		tt.gen.init()
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Batches() {
					for _, v := range x {
						acc += v.work()
					}
				}
				sink = acc
			}
		})
		b.Run(tt.name+"-stream", func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Stream() {
					acc += x.work()
				}
				sink = acc
			}
		})
	}
}

func BenchmarkSmallStruct(b *testing.B) {
	fac := func(r *rand.Rand) Small {
		return NewSmall(r)
	}
	tests := []struct {
		name string
		gen  Generator[Small]
	}{
		{
			name: "batch=1",
			gen: Generator[Small]{
				total:   items,
				batch:   batchOne,
				factory: fac,
				seed:    seed,
			},
		},
		{
			name: "batch=8",
			gen: Generator[Small]{
				total:   items,
				batch:   batchSmall,
				factory: fac,
				seed:    seed,
			},
		},
		{
			name: "batch=1024",
			gen: Generator[Small]{
				total:   items,
				batch:   batchBig,
				factory: fac,
				seed:    seed,
			},
		},
	}

	for _, tt := range tests {
		tt.gen.init()
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Batches() {
					for _, v := range x {
						acc += v.work()
					}
				}
				sink = acc
			}
		})
		b.Run(tt.name+"-stream", func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Stream() {
					acc += x.work()
				}
				sink = acc
			}
		})
	}
}

func BenchmarkSmallStructPtr(b *testing.B) {
	fac := func(r *rand.Rand) *Small {
		b := NewSmall(r)
		return &b
	}
	tests := []struct {
		name string
		gen  Generator[*Small]
	}{
		{
			name: "batch=1",
			gen: Generator[*Small]{
				total:   items,
				batch:   batchOne,
				factory: fac,
				seed:    seed,
			},
		},
		{
			name: "batch=8",
			gen: Generator[*Small]{
				total:   items,
				batch:   batchSmall,
				factory: fac,
				seed:    seed,
			},
		},
		{
			name: "batch=1024",
			gen: Generator[*Small]{
				total:   items,
				batch:   batchBig,
				factory: fac,
				seed:    seed,
			},
		},
	}

	for _, tt := range tests {
		tt.gen.init()
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Batches() {
					for _, v := range x {
						acc += v.work()
					}
				}
				sink = acc
			}
		})
		b.Run(tt.name+"-stream", func(b *testing.B) {
			b.ReportAllocs()
			for b.Loop() {
				var acc uint64
				for x := range tt.gen.Stream() {
					acc += x.work()
				}
				sink = acc
			}
		})
	}
}
