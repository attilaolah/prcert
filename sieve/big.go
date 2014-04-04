// Package sieve is:
// "[an] efficient Eratosthenesque prime sieve using channels".
package sieve

import (
	"math/big"
)

// BigSieve generates prime numbers as *big.Int.
func BigSieve() (out chan *big.Int) {
	out = make(chan *big.Int, 1024)
	go func() {
		for p := range Sieve() {
			out <- big.NewInt(int64(p))
		}
	}()
	return
}
