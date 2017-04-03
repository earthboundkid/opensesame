package main

import (
	crand "crypto/rand"
	"fmt"
	"io"
	"math/rand"
)

func NewRand() *rand.Rand {
	var b [8]byte
	_, err := io.ReadFull(crand.Reader, b[:])
	if err != nil {
		panic(err)
	}
	seed := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
	return rand.New(rand.NewSource(seed))
}

func NewPassword(length int) string {
	const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	pass := make([]byte, 0, length)
	r := NewRand()
	for i := 0; i < cap(pass); i++ {
		char := alphabet[r.Intn(len(alphabet))]
		pass = append(pass, char)
	}
	return string(pass)
}

func main() {
	fmt.Println(NewPassword(8))
}
