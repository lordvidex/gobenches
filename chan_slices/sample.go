package chan_slices

import "math/rand"

// Small is a compact struct (~24 bytes).
type Small struct {
	A int64
	B int64
	C int64
}

// NewSmall ...
func NewSmall(r *rand.Rand) Small {
	return Small{A: r.Int63(), B: r.Int63(), C: r.Int63()}
}

func (s Small) work() uint64 {
	return uint64((s.A ^ s.B) + (s.C & 0xFFFF))
}

// Big is a fat struct (~768 bytes): 12 * 64B = 768B.
type Big struct {
	A [256]byte
	B [256]byte
	C [256]byte
	D [256]byte
	E [256]byte
	F [256]byte
	G [256]byte
	H [256]byte
	I [256]byte
	J [256]byte
	K [256]byte
	L [256]byte
}

// NewBig ...
func NewBig(r *rand.Rand) Big {
	var b Big
	r.Read(b.A[:])
	r.Read(b.B[:])
	r.Read(b.C[:])
	r.Read(b.D[:])
	r.Read(b.E[:])
	r.Read(b.F[:])
	r.Read(b.G[:])
	r.Read(b.H[:])
	r.Read(b.I[:])
	r.Read(b.J[:])
	r.Read(b.K[:])
	r.Read(b.L[:])
	return b
}

func (b Big) work() uint64 {
	var sum uint64
	for _, x := range b.A {
		sum += uint64(x)
	}
	for _, x := range b.B {
		sum += uint64(x)
	}
	for _, x := range b.C {
		sum += uint64(x)
	}
	return sum
}
