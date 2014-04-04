// This example shows how to keep trying to factor a number z, but fall back to
// factoring z-1 if we can't find a prime factor for z in reasonable time.
package main

import (
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/attilaolah/prcert/factor"
)

func main() {
	z, ok := big.NewInt(0).SetString(os.Args[1], 10)
	if !ok {
		fmt.Println("ERR!")
		return
	}

	var p *big.Int
	var err error
	for z.BitLen() > 1 {
		if p, z, err = factor.SplitOrQuit(big.NewInt(0).Set(z), 2*time.Second); err != nil {
			fmt.Println("ERR", err)
			fmt.Println("z--")
			z.Add(z, big.NewInt(-1))
			continue
		}
		fmt.Println(p)
	}

	fmt.Printf("\n")
}
