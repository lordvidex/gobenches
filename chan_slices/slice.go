package chan_slices

import "math/rand"

// Generator ...
type Generator[T any] struct {
	total   int
	batch   int
	factory func(*rand.Rand) T
	seed    int64
}

// Batches returns a channel of batches of T. The size of each batch is specified by the Generator.batch field.
func (g *Generator[T]) Batches() <-chan []T {
	out := make(chan []T, 1)
	go func() {
		defer close(out)
		r := rand.New(rand.NewSource(g.seed))
		remaining := g.total

		for remaining > 0 {
			n := min(g.batch, remaining)

			batch := make([]T, n)
			for i := range batch {
				batch[i] = g.factory(r)
			}
			out <- batch
			remaining -= n
		}
	}()
	return out
}

// Stream returns a channel of T.
func (g *Generator[T]) Stream() <-chan T {
	out := make(chan T, 1)

	go func() {
		defer close(out)
		for x := range g.Batches() {
			for _, v := range x {
				out <- v
			}
		}
	}()

	return out
}
