package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CostumMux struct {
	http.ServeMux
	middlewares []func(next http.Handler) http.Handler
}

func (c *CostumMux) RegisterMiddleware(next func(next http.Handler) http.Handler) {
	c.middlewares = append(c.middlewares, next)
}

func (c *CostumMux) ServerHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &c.ServeMux

	for _, next := range c.middlewares {
		current = next(current)
	}

	current.ServeHTTP(w, r)
}

func main() {
	mux := new(CostumMux)

	mux.HandleFunc("/student", actionStudent)

	mux.RegisterMiddleware(MiddlewareAuth)
	mux.RegisterMiddleware(MiddlewareAllowOnlyGet)

	server := new(http.Server)
	server.Addr = ":9000"
	server.Handler = mux

	fmt.Println("server started at localhost : 9000")
	server.ListenAndServe()
}

func actionStudent(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {
		OutputJSON(w, SellectStudent(id))
		return
	}
	OutputJSON(w, GetStudent())
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)

	if err != nil {
		w.Write([]byte(err.Error()))
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}
