// Package sieve is:
// "[an] efficient Eratosthenesque prime sieve using channels".
package sieve

import (
	"math/big"
)

// BigSieve generates prime numbers as *big.Int.
func BigSieve() (ch chan *big.Int) {
	ch = make(chan *big.Int)
	go func() {
		for p := range Sieve() {
			ch <- big.NewInt(int64(p))
		}
	}()
	return
}
