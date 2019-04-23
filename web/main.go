package main

import "fmt"
import "net/http"

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	var message = "Welcome "
	w.Write([]byte(message))
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	var message = "Hello world!"
	w.Write([]byte(message))
}

func main() {
	http.Handle("/static",
		http.StripPrefix("/static",
			http.FileServer(http.Dir("assets"))))

	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/index", handlerIndex)
	http.HandleFunc("/hello", handlerHello)

	var address = "localhost:9000"
	fmt.Printf("serve started at %s\n", address)

	err := http.ListenAndServe(address, nil)

	if err != nil {
		fmt.Println(err.Error())
	}
}
