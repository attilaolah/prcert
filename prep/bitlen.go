// Package main.
//
// The purpose of this little script is to find the smallest primorial p_n#
// that has a higher bit length than 10⁹⁹⁹⁹⁹⁹⁹⁹. That number has at least one
// hundred million digits, and is the product of many small primes.
package main

import (
	"fmt"

	big "github.com/ncw/gmp"

	"github.com/attilaolah/prcert/cache"
	"github.com/attilaolah/prcert/sieve"
)

func main() {
	i := 1
	n := big.NewInt(1)
	q := big.NewInt(0)
	p, err := cache.BaseExp(10, 99999999)
	if err != nil {
		panic(err)
	}
	l := p.BitLen()
	fmt.Printf("LEN: %10d\n", l)
	for q = range sieve.BigSieve() {
		n.Mul(n, q)
		if i < 1000000 {
			// Skip ahead. For small numbers, console output is the bottleneck.
			i++
			continue
		}
		c := n.BitLen()
		fmt.Printf("\rLEN: %10d %d", c, i)
		if c > l {
			// It looks like we've found our number!
			fmt.Println()
			break
		}
		i++
	}
}
