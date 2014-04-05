// Package cache is used to generate large `*big.Int`s.
// It uses the filesystem to store `GobEncode`d caches of numbers.
package cache

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"path/filepath"
)

const cacheDir = ".cache"

// BaseExp returns b^e.
func BaseExp(b, e int64) (z *big.Int, err error) {
	z = big.NewInt(b)
	filename := filepath.Join(cacheDir, fmt.Sprintf("%d_%d.gob", b, e))
	if err = os.Mkdir(cacheDir, 0755); err != nil && !os.IsExist(err) {
		return
	}
	file, err := os.Open(filename)
	if err == nil {
		var b []byte
		if b, err = ioutil.ReadAll(file); err == nil {
			err = z.GobDecode(b)
		}
		return
	}
	if !os.IsNotExist(err) {
		return
	}
	z.Exp(z, big.NewInt(e), nil)
	err = cache(z, filename)
	return
}

// BaseExpK returns b^e+k.
func BaseExpK(b, e, k int64) (z *big.Int, err error) {
	if z, err = BaseExp(b, e); err == nil {
		z.Add(z, big.NewInt(k))
	}
	return
}

// BaseExpShiftK returns b^(e+s)+k.
func BaseExpShiftK(b, e, s, k int64) (z *big.Int, err error) {
	if z, err = BaseExp(b, e); err == nil {
		sh := big.NewInt(b)
		sh.Exp(sh, big.NewInt(abs(s)), nil)
		if s > 0 {
			z.Mul(z, sh)
		} else if s != 0 {
			z.Div(z, sh)
		}
		z.Add(z, big.NewInt(k))
	}
	return
}

func cache(z *big.Int, filename string) (err error) {
	var b []byte
	if b, err = z.GobEncode(); err == nil {
		err = ioutil.WriteFile(filename, b, 0644)
	}
	return
}

func abs(i int64) int64 {
	if i < 0 {
		i = -i
	}
	return i
}
