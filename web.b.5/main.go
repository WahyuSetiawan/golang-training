package main

import (
	"html/template"
	"net/http"
	"path"
)

type M map[string]interface{}

func main() {
	// var tmpl, err = template.ParseGlob("views/*")

	// if err != nil {
	// 	panic(err.Error())
	// 	return
	// }

	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "Batman"}
		// err = tmpl.ExecuteTemplate(w, "index", data)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		var tmpl = template.Must(template.ParseFiles(
			path.Join("views", "index.html"),
			path.Join("views", "_header.html"),
			path.Join("views", "_message.html"),
		))

		var err = tmpl.ExecuteTemplate(w, "index", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		var data = M{"name": "K"}
		// err = tmpl.ExecuteTemplate(w, "about", data)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// }

		var tmpl = template.Must(template.ParseFiles(
			path.Join("views", "index.html"),
			path.Join("views", "_header.html"),
			path.Join("views", "_message.html"),
		))

		var err = tmpl.ExecuteTemplate(w, "about", data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	http.ListenAndServe(":9000", nil)
}
