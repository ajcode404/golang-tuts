package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "oops", http.StatusBadRequest)
			return
		}
		fmt.Fprintf(w, "Hello %s\n", d)
		log.Printf("Data %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Good bye")
	})

	http.ListenAndServe(":8080", nil)
}
