package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleIndex)
	http.HandleFunc("/save", handleSave)

	http.Handle("/static", http.StripPrefix("/static",
		http.FileServer(http.Dir("assets"))))

	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("view.html"))
	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func handleSave(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		decorder := json.NewDecoder(r.Body)
		payload := struct {
			Name   string `json:"name"`
			Age    int    `json:"age"`
			Gender string `json:"gender"`
		}{}

		if err := decorder.Decode(&payload); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		message := fmt.Sprintf(
			"hello, my name is %s. I'm %d year old %s",
			payload.Name,
			payload.Age,
			payload.Gender,
		)
		w.Write([]byte(message))
		return
	}

	http.Error(w, "Only accept POST request", http.StatusBadRequest)
}
