package main

import (
	"log"
	"net/http"
	"os"

	"main.go/handlers"
)

func main() {

	l := log.New(os.Stdout, "[prodcut-api] ", log.Flags())
	hh := handlers.NewHello(l)
	gh := handlers.NewGoodbye(l)

	sm := http.NewServeMux()
	sm.Handle("/", hh)
	sm.Handle("/goodbye", gh)

	http.ListenAndServe(":8080", sm)
}
