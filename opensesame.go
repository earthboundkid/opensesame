package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/carlmjohnson/opensesame/pass"
)

func main() {
	const (
		upper           = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower           = "abcdefghijklmnopqrstuvwxyz"
		digit           = "0123456789"
		defaultAlphabet = "upper lower digit"
	)

	length := flag.Int("length", 8, "length of password to generate")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage of %s [opts] [alphabet]:

	Creates a password by randomly selecting characters from its alphabet.

	Alphabet is a space separated list of character classes to use.
	At least one character in each class will be output.
	Character classes are either literal sets (like "abc" and "123") or the
	special names "upper", "lower", "digit", and "default".

	Default alphabet is %q.

`, os.Args[0], defaultAlphabet)
		flag.PrintDefaults()
	}
	flag.Parse()

	alpha := flag.Arg(0)
	if alpha == "" {
		alpha = defaultAlphabet
	}
	alpha = strings.Replace(alpha, "default", defaultAlphabet, -1)
	alpha = strings.Replace(alpha, "upper", upper, -1)
	alpha = strings.Replace(alpha, "lower", lower, -1)
	alpha = strings.Replace(alpha, "digits", digit, -1)
	alpha = strings.Replace(alpha, "digit", digit, -1)

	alphas := strings.Split(alpha, " ")

	if pass, err := pass.New(*length, alphas...); err == nil {
		fmt.Println(pass)
	} else {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}
}
