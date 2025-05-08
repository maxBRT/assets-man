package main

import (
	"math/rand/v2"
	"net/http"
	"text/template"
)

func (cfg *apiConfig) addUser(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	const tmpl = `
		<p>Hello, {{ .Name }}! Today's lucky number is {{ .LuckyNumber }}.</p>
	`
	t := template.Must(template.New("response").Parse(tmpl))
	data := struct {
		Name        string
		LuckyNumber int
	}{
		Name:        email,
		LuckyNumber: rand.IntN(100),
	}

	// Render the template with data
	t.Execute(w, data)
}
