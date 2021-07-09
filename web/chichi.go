package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("index page\n")
		w.Write([]byte("index page"))
	})
	r.Get("/profile", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("profile page\n")
		w.Write([]byte("profile page"))
	})
	r.Post("/login", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("login: %s = %s\n", key, value)
		}
		w.Write([]byte("login page"))
	})
	r.Post("/logout", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		for key, value := range r.Form {
			fmt.Printf("logout: %s = %s\n", key, value)
		}
		w.Write([]byte("logout page"))
	})
	http.ListenAndServe(":3000", r)
}
