package factor

import (
	"errors"
	"math/big"
	"time"

	"github.com/attilaolah/prcert/sieve"
)

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
			if q.QuoRem(z, p, r); r.Sign() == 0 {
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
