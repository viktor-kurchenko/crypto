package main

import (
	"crypto/sha256"
	"fmt"
	"hash"

	"github.com/jzelinskie/whirlpool"
	"golang.org/x/crypto/sha3"
)

type hashFn struct {
	name string
	fn   func() hash.Hash
}

func main() {
	phrase := []byte("hello")

	printHashes(phrase)
}

func printHashes(phrase []byte) {
	hashFns := []hashFn{
		{name: "SHA-224", fn: sha256.New224},
		{name: "SHA-256", fn: sha256.New},
		{name: "SHA3-224", fn: sha3.New224},
		{name: "SHA3-256", fn: sha3.New256},
		{name: "SHA3-384", fn: sha3.New384},
		{name: "Whirlpool", fn: whirlpool.New},
	}
	for i := range hashFns {
		h := hashFns[i].fn()
		h.Write(phrase)
		fmt.Printf("%s: %x\n", hashFns[i].name, h.Sum(nil))
	}
}
