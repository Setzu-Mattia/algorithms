package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math"
	"math/big"
	"math/rand"
)

var md5H = md5.New()

func main() {
	fmt.Println("rk('aaa', 'a')", rabinKarp("aaa", "a"))
}

func rabinKarp(str, pattern string) int {
	hString := rabinFingerprint(str, pattern)
	hPattern := rabinFingerprint(pattern, pattern)
	strLen := len(str)
	patternLen := len(pattern)

	for i := 0; i < strLen-patternLen; i++ {
		if hString == hPattern {
			if str[i:i+patternLen-1] == pattern[1:patternLen-1] {
				return i
			}
		}

		hString = rabinFingerprint(str[i+1 : i+patternLen], pattern)
	}

	return -1
}

func rabinFingerprint(str string, pattern string) *big.Int {
	fingerPrint := big.NewInt(0)
	n := len(str)
	m := len(pattern)
	l := int64(math.Log(float64((n^2)*m)))
	pol := big.NewInt(0)
	prime := big.NewInt(0)

	for i, z := 0, big.NewInt(0); i < n; i++ {
		z = z.Mul(hash(str[i:i+1]), (big.NewInt(int64(2 ^ i))))
		pol = z
	}

	prime = big.NewInt(pickPrime((int64(n ^ 2)) * l))


	fingerPrint = fingerPrint.Mod(pol, prime)
	return fingerPrint
}

func hash(str string) *big.Int {
	bi := big.NewInt(0)
	md5H.Write([]byte(str))
	hexstr := hex.EncodeToString(md5H.Sum(nil))
	bi.SetString(hexstr, 16)

	return bi
}

// Pick a prim > k
func pickPrime(k int64) int64 {
	randIndex := rand.Intn(1000)
	prime := int64(0)
	ch := make(chan int64)
	go generate(ch)

	for prime < k {
		prime = <-ch
		ch1 := make(chan int64)
		go filter(ch, ch1, prime)
		ch = ch1

		if prime > k {
			randIndex--
		}
	}

	return int64(prime)
}

func generate(ch chan<- int64) {
	for i := int64(2); ; i++ {
		ch <- i
	}
}

func filter(in <-chan int64, out chan<- int64, prime int64) {
	for {
		i := <-in
		if i%prime != 0 {
			out <- i
		}
	}
}
