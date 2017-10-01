package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/carlmjohnson/opensesame/pass"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

func pageHandler(w http.ResponseWriter, r *http.Request) {
	const (
		upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower = "abcdefghijklmnopqrstuvwxyz"
		digit = "0123456789"
	)

	// Validate request
	_ = r.ParseForm()
	alphabets := r.Form["alpha"]
	// Filter empty strings
	c := 0
	for _, s := range alphabets {
		if s != "" {
			alphabets[c] = s
			c++
		}
	}
	alphabets = alphabets[:c]

	// Figure out which extra boxes were checked
	for _, n := range r.Form["checkboxes"] {
		if v := r.Form.Get("alpha-" + n); v != "" {
			alphabets = append(alphabets, v)
		}
	}

	if len(alphabets) == 0 {
		alphabets = []string{upper, lower, digit}
	}

	lengthStr := r.Form.Get("length")
	length, _ := strconv.Atoi(lengthStr)
	if length < 1 || length > 256 {
		length = 8
	}

	// Get template values
	pass, err := pass.New(length, alphabets...)
	if err != nil {
		log.Printf("Error %s %q %v", r.URL, r.UserAgent(), err)
		http.Error(w, "Something went wrong", 500)
		return
	}

	// Filter out known alphabets
	var hasUpper, hasLower, hasDigit bool
	c = 0
	for _, s := range alphabets {
		switch {
		case !hasUpper && s == upper:
			hasUpper = true
		case !hasLower && s == lower:
			hasLower = true
		case !hasDigit && s == digit:
			hasDigit = true
		default:
			alphabets[c] = s
			c++
		}
	}
	alphabets = alphabets[:c]

	// Respond
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	err = tmpl.ExecuteTemplate(w, "index.html", struct {
		Length                                   int
		ExtraAlphabets                           []string
		Password                                 string
		CheckedUpper, CheckedLower, CheckedDigit bool
	}{
		Length:         length,
		ExtraAlphabets: alphabets,
		Password:       pass,
		CheckedUpper:   hasUpper,
		CheckedLower:   hasLower,
		CheckedDigit:   hasDigit,
	})
	if err != nil {
		log.Printf("Template error: %v", err)
	}
}
