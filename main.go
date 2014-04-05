// Example that factors big numbers.
package main

import (
	"fmt"
	"math/big"
	"os"
	"strconv"

	"github.com/attilaolah/prcert/cache"
	"github.com/attilaolah/prcert/factor"
)

func main() {

	if len(os.Args) != 5 {
		fmt.Printf("usage: %s b e s k (z=b^(e+s)+k)\n", os.Args[0])
		return
	}
	var err error
	var base, exp, s, k int64
	if base, err = strconv.ParseInt(os.Args[1], 10, 64); err != nil {
		fmt.Println(err)
		return
	}
	if exp, err = strconv.ParseInt(os.Args[2], 10, 64); err != nil {
		fmt.Println(err)
		return
	}
	if s, err = strconv.ParseInt(os.Args[3], 10, 64); err != nil {
		fmt.Println(err)
		return
	}
	if k, err = strconv.ParseInt(os.Args[4], 10, 64); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("IN: %d^(%d%+d)%+d\n", base, exp, s, k)
	z, err := cache.BaseExpShiftK(base, exp, s, k)
	if err != nil {
		fmt.Println("ERR:", err)
		return
	}
	mods := factor.ModsAfter(z, big.NewInt(1487))
	for p := range mods {
		fmt.Println(p, <-mods)
	}
}
