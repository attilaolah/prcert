// Example that factors big numbers.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/attilaolah/prcert/cache"
	"github.com/attilaolah/prcert/factor"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Printf("usage: %s base exp k\n", os.Args[0])
		return
	}
	var err error
	var base, exp, k int64
	if base, err = strconv.ParseInt(os.Args[1], 10, 64); err != nil {
		fmt.Println(err)
		return
	}
	if exp, err = strconv.ParseInt(os.Args[2], 10, 64); err != nil {
		fmt.Println(err)
		return
	}
	if k, err = strconv.ParseInt(os.Args[3], 10, 64); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(base, exp, k)
	z, err := cache.BaseExpK(base, exp, k)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(z.BitLen())
	p, q := factor.Split(z)
	fmt.Println(p, q.BitLen())
}
