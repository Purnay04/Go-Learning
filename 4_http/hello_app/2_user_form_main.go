// run the file from main directory path 'hello_app'

package main

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	// render the html form
	http.HandleFunc("/hello", helloHandler)

	// render the o/p after reading the form
	http.HandleFunc("/welcome", welcomeHandler)

	log.Println("Server running at http://localhost:8081/hello")
	http.ListenAndServe(":8081", nil)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "form.html", nil)
}

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/hello", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Form error", http.StatusBadRequest)
		return
	}
	name := r.FormValue("name")
	renderTemplate(w, "welcome.html", struct{ Name string }{Name: name})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("templates", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "Template parsing error: "+err.Error(), http.StatusInternalServerError)
		return
	}

	t.Execute(w, data)
}
