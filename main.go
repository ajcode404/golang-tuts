package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"main.go/handlers"
)

func main() {

	l := log.New(os.Stdout, "[prodcut-api] ", log.Flags())

	// create the handlers
	ph := handlers.NewProducts(l)

	// create a new serve mux and register handlers
	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// create a new server
	s := http.Server{
		Addr:         ":8080",           // configure the bind address
		Handler:      sm,                // set the default handler
		ErrorLog:     l,                 // set the longer  for the server
		ReadTimeout:  1 * time.Second,   // max time to read request from the client
		WriteTimeout: 1 * time.Second,   // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections ysing TCP Keep-Alive
	}
	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)
	context, _ := context.WithTimeout(context.Background(), 3*time.Second)
	s.Shutdown(context)
}
