package main

import (
	"bytes"
	crand "crypto/rand"
	"errors"
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

var (
	ErrPassLength      = errors.New("Length of password must be greater than 1")
	ErrPassAlphaLength = errors.New("Length of password must be greater than or " +
		"equal to number of required characters")
	ErrAlphaLength = errors.New("Alphabet must not be empty")
)

func NewPassword(length int, alphabets ...string) (string, error) {
	alphabet := strings.Join(alphabets, "")

	if length < 1 {

		return "", ErrPassLength
	}

	if length < len(alphabets) {
		return "", ErrPassAlphaLength
	}

	if alphabet == "" {
		return "", ErrAlphaLength
	}

	pass := make([]byte, 0, length)
	r := NewRand()

	// Loop until the generated password has required characteristics
loop:
	for i := 0; i < cap(pass); i++ {
		char := alphabet[r.Intn(len(alphabet))]
		pass = append(pass, char)
	}

	for _, alphabet := range alphabets {
		if !bytes.ContainsAny(pass, alphabet) {
			pass = pass[:0]
			goto loop
		}
	}

	return string(pass), nil
}

func main() {
	alpha := "ABCDEFGHIJKLMNOPQRSTUVWXYZ abcdefghijklmnopqrstuvwxyz 0123456789"
	length := flag.Int("length", 8, "length of password to generate")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s [opts] [alphabet]:

	Creates a password by randomly selecting characters from its alphabet.

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

	if pass, err := NewPassword(*length, alphas...); err == nil {
		fmt.Println(pass)
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
