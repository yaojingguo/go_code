package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "form: %v\n", r.Form)
		fmt.Fprintf(w, "name: %s, sex: %s\n", r.FormValue("name"), r.FormValue("sex"))
		fmt.Fprintf(w, "form: %v\n", r.Form)
		fmt.Fprintln(w, "hello, world")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
