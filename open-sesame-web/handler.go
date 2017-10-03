package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/carlmjohnson/opensesame/pass"
)

var tmpl = template.Must(template.ParseGlob("templates/*"))

const (
	upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	lower = "abcdefghijklmnopqrstuvwxyz"
	digit = "0123456789"
)

var labels = map[string]string{
	upper: "Uppercase",
	lower: "Lowercase",
	digit: "Digits",
}

func pageHandler(w http.ResponseWriter, r *http.Request) {

	// Validate request
	_ = r.ParseForm()

	// Figure out which extra boxes were checked
	// and add corresponding input values to alphabets
	alphabets := make([]string, 0, len(r.Form["checkboxes"]))
	for _, cbVal := range r.Form["checkboxes"] {
		if inpVal := r.Form.Get(cbVal); inpVal != "" {
			alphabets = append(alphabets, inpVal)
		}
	}

	// Filter empty alphabet strings
	c := 0
	for _, s := range alphabets {
		if s != "" {
			alphabets[c] = s
			c++
		}
	}
	alphabets = alphabets[:c]

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

	type templateAlpha struct {
		Label, Value string
	}

	templateAlphas := make([]templateAlpha, 0, len(alphabets))
	for _, alpha := range alphabets {
		label := labels[alpha]
		if label == "" {
			label = "Requirement"
		}
		templateAlphas = append(templateAlphas, templateAlpha{label, alpha})
	}

	// Respond
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	err = tmpl.ExecuteTemplate(w, "index.html", struct {
		Length    int
		Password  string
		Alphabets []templateAlpha
	}{
		Length:    length,
		Password:  pass,
		Alphabets: templateAlphas,
	})
	if err != nil {
		log.Printf("Template error: %v", err)
	}
}
