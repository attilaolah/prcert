package main

import (
	"fmt"

	"github.com/attilaolah/prcert/sieve"
)

const k = 99999999

func main() {
	fmt.Println(1, 2, 0)
	fmt.Println(2, 3, 2)
	fmt.Println(3, 5, 0)

	j := 4
	var m, p, i uint64
	for p = range sieve.Sieve7() {
		m, i = 10, 1
		for ; m != 1; i++ {
			m = (10 * m) % p
		}
		for i = k % i; i != 0; i-- {
			m = (10 * m) % p
		}
		fmt.Println(j, p, p-m)
		j++
	}
}
