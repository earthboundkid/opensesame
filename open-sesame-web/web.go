package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/carlmjohnson/opensesame/pass"
)

func main() {
	const (
		upper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
		lower = "abcdefghijklmnopqrstuvwxyz"
		digit = "0123456789"
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// subscribe to SIGINT signals
	stopChan := make(chan os.Signal)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Serving %s %q", r.URL, r.UserAgent())

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		pass, err := pass.New(8, upper, lower, digit)
		if err != nil {
			log.Printf("Error %s %q %v", r.URL, r.UserAgent(), err)
			http.Error(w, "Something went wrong", 500)
			return
		}
		fmt.Fprintf(w, "Password: %s\n", pass)
	})

	srv := &http.Server{Addr: ":" + port, Handler: http.DefaultServeMux}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Printf("listen: %s\n", err)
		}
	}()

	<-stopChan // wait for SIGINT
	log.Println("Shutting down server...")

	// shut down gracefully, but wait no longer than 5 seconds before halting
	ctx, c := context.WithTimeout(context.Background(), 5*time.Second)
	defer c()
	srv.Shutdown(ctx)

	log.Println("Server gracefully stopped")
}
