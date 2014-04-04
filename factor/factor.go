package factor

import (
	"errors"
	"math/big"
	"time"

	"github.com/attilaolah/prcert/sieve"
)

var ErrTimeout = errors.New("timeout while trying to find a prime factor")

// Timeout1 tries to call Split, but returns after t has elapsed.
// TODO: implement a stop channel to signal Split to die, now it zombies off.
func Timeout1(z *big.Int, t time.Duration) (p, q *big.Int, err error) {
	ch := asyncSplit(z)
	select {
	case p = <-ch:
		q = <-ch
	case <-time.After(t):
		err = ErrTimeout
	}
	return
}

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
	q, r := big.NewInt(0), big.NewInt(0)
	if z.Sign() == 0 {
		return
	}
	max := roughSqrt(z)
	for p = range sieve.BigSieve() {
		if q.QuoRem(z, p, r); r.Sign() == 0 {
			break
		}
		if max.Cmp(p) == -1 {
			q.SetInt64(1)
			p.Set(z)
			break
		}
	}
	return
}

// Route Split to a channel and return immediately.
func asyncSplit(z *big.Int) (ch chan *big.Int) {
	ch = make(chan *big.Int, 2)
	go func() {
		defer close(ch)
		p, q := Split(z)
		ch <- p
		ch <- q
	}()
	return
}

// Rought square root of z.
func roughSqrt(z *big.Int) *big.Int {
	return big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64((z.BitLen()+1)/2)), nil)
}