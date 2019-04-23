package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"web.b.22/config"
)

type CostumMux struct {
	http.ServeMux
}

func (c CostumMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if config.Configuration().Log.Verbose {
		log.Println("Incoming request from", r.Host, "accessing", r.URL.String())
	}
	c.ServeMux.ServeHTTP(w, r)
}

func main() {
	router := new(CostumMux)
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})

	router.HandleFunc("/howareyou", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("How arer you?"))
	})

	server := new(http.Server)
	server.Handler = router
	server.ReadTimeout = config.Configuration().Server.ReadTimeout * time.Second
	server.WriteTimeout = config.Configuration().Server.WriteTimeout * time.Second
	server.Addr = fmt.Sprintf(":%d", config.Configuration().Server.Port)

	if config.Configuration().Log.Verbose {
		log.Printf("Starting server at %s \n", server.Addr)
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
