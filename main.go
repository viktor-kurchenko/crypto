package main

import (
	"crypto/hmac"
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

var hashFns = []hashFn{
	{name: "SHA-224", fn: sha256.New224},
	{name: "SHA-256", fn: sha256.New},
	{name: "SHA3-224", fn: sha3.New224},
	{name: "SHA3-256", fn: sha3.New256},
	{name: "SHA3-384", fn: sha3.New384},
	{name: "Whirlpool", fn: whirlpool.New},
}

func main() {
	phrase := []byte("hello")
	key := []byte("cryptography")

	printHashes(phrase)
	printHMACs(key, phrase)
}

func printHashes(phrase []byte) {
	for i := range hashFns {
		h := hashFns[i].fn()
		h.Write(phrase)
		fmt.Printf("%s: %x\n", hashFns[i].name, h.Sum(nil))
	}
}

func printHMACs(key, phrase []byte) {
	for i := range hashFns {
		h := hmac.New(hashFns[i].fn, key)
		h.Write(phrase)
		fmt.Printf("HMAC-%s: %x\n", hashFns[i].name, h.Sum(nil))
	}
}
