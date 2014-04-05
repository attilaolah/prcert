// Package factor implements prime factorisation functions.
package factor

import (
	big "github.com/ncw/gmp"

	"github.com/attilaolah/prcert/sieve"
)

type modder struct {
	z, q, r, t *big.Int
	sieve      chan *big.Int
}

// Modder creates a generator that yields remainders of z for each prime.
func Modder(z *big.Int) *modder {
	return &modder{
		z:     z,
		q:     big.NewInt(1),
		r:     big.NewInt(0),
		t:     big.NewInt(0),
		sieve: sieve.BigSieve(),
	}
}

// Step skips one ahead.
func (m *modder) Step() {
    <-m.sieve
}

// Next yields the next prime divisor and the remainder.
func (m *modder) Next() (p *big.Int, t *big.Int) {
    p = <-m.sieve
    m.z.QuoRem(m.z, p, m.t)
    m.t.Mul(m.t, m.q)
    m.t.Add(m.t, m.r)
    m.q.Mul(m.q, p)
    m.r.Set(m.t)
    t = m.t.Rem(m.t, p)
	return
}
