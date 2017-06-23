package main

import (
	"bytes"
	crand "crypto/rand"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
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

func NewPassword(length int, alphabets ...string) string {
	alphabet := strings.Join(alphabets, "")
	pass := make([]byte, 0, length)
	r := NewRand()

	for {
		for i := 0; i < cap(pass); i++ {
			char := alphabet[r.Intn(len(alphabet))]
			pass = append(pass, char)
		}
		done := true
		for _, alphabet := range alphabets {
			if !bytes.ContainsAny(pass, alphabet) {
				done = false
				pass = pass[:0]
				break
			}
		}
		if done {
			return string(pass)
		}
	}
}

func main() {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz 0123456789"
	length := flag.Int("length", 8, "length of password to generate")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s [opts] [alphabet]:

	Alphabet is a space separated list of character classes to use.
	At least one character in each class will be output.
	Default alphabet is one upper, one lower, one digit (%q).

`, os.Args[0], alpha)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.Arg(0) != "" {
		alpha = flag.Arg(0)
	}
	alphas := strings.Split(alpha, " ")

	fmt.Println(NewPassword(*length, alphas...))
}
