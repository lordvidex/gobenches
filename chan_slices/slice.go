package chan_slices

import "math/rand"

// Generator ...
type Generator[T any] struct {
	total   int
	batch   int
	factory func(*rand.Rand) T
	seed    int64

	// buf is very large.. buf preallocates all total items to prevent allocations during bench runs
	buf []T
}

// init preallocates buf
func (g *Generator[T]) init() {
	g.buf = make([]T, g.total)
	r := rand.New(rand.NewSource(g.seed))
	for i := range g.buf {
		g.buf[i] = g.factory(r)
	}
}

// Batches returns a channel of batches of T. The size of each batch is specified by the Generator.batch field.
func (g *Generator[T]) Batches() <-chan []T {
	out := make(chan []T, 1)
	go func() {
		defer close(out)
		for i := 0; i < g.total; i += g.batch {
			out <- g.buf[i:min(i+g.batch, g.total)]
		}
	}()
	return out
}

// Stream returns a channel of T.
func (g *Generator[T]) Stream() <-chan T {
	out := make(chan T, 1)

	go func() {
		defer close(out)
		for _, x := range g.buf {
			out <- x
		}
	}()

	return out
}
