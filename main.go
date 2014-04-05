// Example that factors big numbers.
// Example that generates remainders of 10⁹⁹⁹⁹⁹⁹⁹⁹/p for each prime p.
// Results are stored in a local file, so the search can be resumed later.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/attilaolah/prcert/cache"
	"github.com/attilaolah/prcert/factor"
)

func main() {
	z, _ := cache.BaseExpShiftK(10, 100000000, -1, 0)
	m := factor.Modder(z)

	f, err := os.Open(".10_99999999.factors")
	if err != nil {
		panic(err.Error())
	}

	var line string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line = scanner.Text()
	}
	mark, err := strconv.ParseInt(strings.Split(strings.TrimSpace(line), " ")[0], 10, 64)
	if err != nil {
		panic(err.Error())
	}

	for i := int64(1); ; i++ {
		if i <= mark {
			m.Step()
			continue
		}
		p, r := m.Next()
		fmt.Println(i, p, r)
	}
}
