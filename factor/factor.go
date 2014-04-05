// Package factor implements prime factorisation functions.
package factor

import (
	"errors"
	"math/big"
	"time"

	"github.com/attilaolah/prcert/sieve"
)

// ErrTimeout means the timeout was reached during a function call.
var ErrTimeout = errors.New("timeout while trying to find a prime factor")

// Factor returns a channel and sends all prime factors of z on that channel.
func Factor(z *big.Int) (ch chan *big.Int) {
	ch = make(chan *big.Int, 1024)
	go func() {
		defer close(ch)
		p, q := Split(z)
		ch <- p
		for q.BitLen() > 1 {
			p, q = Split(q)
			ch <- p
		}
	}()
	return
}

// Split tries to find the first prime factor of z.
// If found, it returns the factor and the remainder.
// If no prime factor is found, the first return argument will be set to z.
func Split(z *big.Int) (p, q *big.Int) {
	p, q, _ = splitOrQuit(z, make(chan time.Time))
	return
}

// SplitOrQuit is just like Split, but it returns an error if it times out.
func SplitOrQuit(z *big.Int, t time.Duration) (p, q *big.Int, err error) {
	return splitOrQuit(z, time.After(t))
}

// Mods sends primes and remainders after deviding z by the prime.
func Mods(z *big.Int) (ch chan *big.Int) {
	return ModsAfter(z, big.NewInt(0))
}

// ModsAfter is like Mods, but skips primes smaller than m.
// Note that it re-uses both z and the remainder that it returns.
// All returned values are only valid during one iteration.
// TODO: convert this to a generator and drop the channel.
func ModsAfter(z, m *big.Int) (ch chan *big.Int) {
	q := big.NewInt(1)
	r := big.NewInt(0)
	t := big.NewInt(0)
	ch = make(chan *big.Int)
	go func() {
		// Endless channel, no need to close.
		for p := range sieve.BigSieve() {
			if p.Cmp(m) < 0 {
				continue
			}
			ch <- p
			z.QuoRem(z, p, t)
			t.Mul(t, q)
			t.Add(t, r)
			q.Mul(q, p)
			r.Set(t)
			ch <- t.Rem(t, p)
		}
	}()
	return
}

// Just like Split, but return an error when receiving a kill signal from t.
func splitOrQuit(z *big.Int, quit <-chan time.Time) (p, q *big.Int, err error) {
	q, r := big.NewInt(0), big.NewInt(0)
	if z.Sign() == 0 {
		return
	}
	max := roughSqrt(z)
	primes := sieve.BigSieve()
	for {
		select {
		case <- quit:
			err = ErrTimeout
			return
		case p = <-primes:
			if q.DivMod(z, p, r); r.Sign() == 0 {
				return
			}
			if max.Cmp(p) == -1 {
				q.SetInt64(1)
				p.Set(z)
				return
			}
		}
	}
	return
}

// Rought square root of z.
func roughSqrt(z *big.Int) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64((z.BitLen()+1)/2)), nil)
}
