package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

var count int64 = 0

type response struct {
	Count int64 `json:"count"`
}

func newResponse(count int64)  *response{
	return &response{Count: count}
}

func main() {
	// Flag
	port := flag.String("p", "80", "Port")
	flag.Parse()
	startHttpServer(*port)
}

// handles common request
func visitHandler(w http.ResponseWriter, r *http.Request) {
	c := atomic.AddInt64(&count, 1)
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(newResponse(c))

	w.Write(b)
}

// starts http server
func startHttpServer(port string) {
	srv := createHttpServer(port)
	ch := make(chan struct{})
	go func() {
		defer close(ch)

		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-ch
}

// creates HTTP server
func createHttpServer(port string) *http.Server {
	//r := mux.NewRouter()
	r := http.NewServeMux()
	r.HandleFunc("/", visitHandler)
	srv := &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Printf("server listening on %s\n", srv.Addr)
	return srv
}
