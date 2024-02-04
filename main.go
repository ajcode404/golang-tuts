package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}

		fmt.Fprintf(rw, "Hello %s\n", data)
		log.Printf("We hit %s", data)
	})

	http.ListenAndServe(":8080", nil)
}
