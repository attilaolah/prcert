package sieve

import (
	"math/big"
)

func BigSieve() (out chan *big.Int) {
	out = make(chan *big.Int, 1024)
	go func() {
		for p := range Sieve() {
			out <- big.NewInt(int64(p))
		}
	}()
	return
}
