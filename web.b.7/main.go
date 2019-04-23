package main

import "net/http"
import "fmt"
import "html/template"

type SuperHero struct {
	Name    string
	Alias   string
	Friends []string
}

func (s SuperHero) SayHello(from string, message string) string {
	return fmt.Sprintf("%s said: \"%s\"", from, message)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var person = SuperHero{
			Name:    "Bruce Wayne",
			Alias:   "Batman",
			Friends: []string{"Superman", "Flash", "Green Lentern"},
		}

		var tmpl = template.Must(template.ParseFiles("view.html"))

		if err := tmpl.Execute(w, person); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}
