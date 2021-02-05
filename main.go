package main

import (
	"net/http"
)

type helloHandler struct{}

func (m *helloHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

type aboutHandler struct{}

func (m *aboutHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("About!"))
}

func welcome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("welcome!"))
}

func main() {
	hh := helloHandler{}
	ah := aboutHandler{}
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello world"))
	// })
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: nil,
	}
	http.Handle("/hello", &hh)
	http.Handle("/about", &ah)
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Home!"))
	})
	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	http.ServeFile(w, r, "wwwroot"+r.URL.Path)
	// })
	http.ListenAndServe(":8080", http.FileServer(http.Dir("wwwroot")))
	http.Handle("/welcome", http.HandlerFunc(welcome))
	server.ListenAndServe()
	// http.ListenAndServe("localhost:8080", nil)//DefaultServeMux
}
