package main

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/carlmjohnson/opensesame/pass"
)

const html = `<html>
    <head>
        <title>Open Sesame Web</title>
    </head>
    <body>
        <p>Your password:</p>
        <h1>{{ .Password }}</h1>
        <p>Create a new password:</p>
        <form action="" method="get">
            <fieldset id="alphabets">
            <button id="add-button" type="button">+</button>
            </fieldset>

            <label for="length">Length</label>
            <input
                type="number"
                id="length"
                name="length"
                autcomplete="off"
                value="{{ .Length }}"
                max="256"
                min="1"
            >
            <button type="submit">Generate Password!</button>
        </form>
        <script type="text/javascript">
document.addEventListener("DOMContentLoaded", () => {
  let defaultAlphabets = {{ .Alphabets }};
  let container = document.getElementById("alphabets");
  let addBtn = document.getElementById("add-button");
  let createAlpha = alpha => {
    let ta = document.createElement("textarea");
    ta.value = alpha;
    ta.name = "alpha";

    let button = document.createElement("button");
    button.attributes.type = "button";
    button.textContent = "-";
    button.addEventListener("click", () => {
      ta.remove();
      button.remove();
    });

    container.insertBefore(ta, addBtn);
    container.insertBefore(button, addBtn);
  };

  defaultAlphabets.forEach(createAlpha);

  addBtn.addEventListener("click", () => createAlpha(""));
});
        </script>
    </body>
</html>
`

var tmpl = template.Must(template.New("").Parse(html))

func pageHandler(w http.ResponseWriter, r *http.Request) {
	const (
		upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower = "abcdefghijklmnopqrstuvwxyz"
		digit = "0123456789"
	)

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

	if len(alphabets) == 0 {
		alphabets = []string{upper, lower, digit}
	}
	lengthStr := r.Form.Get("length")
	length, _ := strconv.Atoi(lengthStr)
	if length < 1 || length > 256 {
		length = 8
	}
	pass, err := pass.New(length, alphabets...)
	if err != nil {
		log.Printf("Error %s %q %v", r.URL, r.UserAgent(), err)
		http.Error(w, "Something went wrong", 500)
		return
	}
	err = tmpl.Execute(w, struct {
		Length    int
		Alphabets []string
		Password  string
	}{
		Length:    length,
		Alphabets: alphabets,
		Password:  pass,
	})
	if err != nil {
		log.Printf("Template error: %v", err)
	}
}
