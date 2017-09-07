package pass

import (
	"bytes"
	crand "crypto/rand"
	"errors"
	"io"
	"math/rand"
	"strings"
	"time"
)

func Rand() *rand.Rand {
	var b [8]byte
	_, err := io.ReadFull(crand.Reader, b[:])
	if err != nil {
		panic(err)
	}
	seed := int64(b[0]) | int64(b[1])<<8 | int64(b[2])<<16 | int64(b[3])<<24 |
		int64(b[4])<<32 | int64(b[5])<<40 | int64(b[6])<<48 | int64(b[7])<<56
	return rand.New(rand.NewSource(seed))
}

const deadline = 500 * time.Millisecond

var (
	ErrPassLength      = errors.New("Length of password must be greater than 1")
	ErrPassAlphaLength = errors.New("Length of password must be greater than or " +
		"equal to number of required characters")
	ErrAlphaLength = errors.New("Alphabet must not be empty")
	ErrTimeOut     = errors.New("Could not find a matching password in " + deadline.String())
)

func New(length int, alphabets ...string) (string, error) {
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
	r := Rand()

	// Loop until the generated password has required characteristics
	// or time runs out
	start := time.Now()
	for time.Since(start) < deadline {

		for i := 0; i < cap(pass); i++ {
			char := alphabet[r.Intn(len(alphabet))]
			pass = append(pass, char)
		}

		missingAlphabet := false
		for _, subalpha := range alphabets {
			if !bytes.ContainsAny(pass, subalpha) {
				pass = pass[:0]
				missingAlphabet = true
				break
			}
		}

		if !missingAlphabet {
			return string(pass), nil
		}
	}
	return "", ErrTimeOut
}
