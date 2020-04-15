package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	ListenAddress string `envconfig:"LISTEN_ADDRESS" default:":8080"`
	Version       string `envconfig:"VERSION" required:"true"`
}

func main() {
	var cfg config
	envconfig.MustProcess("", &cfg)

	var healthy int32
	log.SetOutput(os.Stdout)
	log.SetPrefix("sample: ")
	log.SetFlags(log.LstdFlags)

	mux := http.NewServeMux()
	mux.HandleFunc("/version", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(w, cfg.Version)
	})
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, req *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusOK)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
	srv := &http.Server{
		Addr:         cfg.ListenAddress,
		Handler:      mux,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan struct{})
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		log.Println("server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			log.Fatalf("could not gracefully shutdown the server: %v", err)
		}
		close(done)
	}()

	log.Println("server is ready to handle requests at", cfg.ListenAddress)
	atomic.StoreInt32(&healthy, 1)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("could not listen on %s: %v", cfg.ListenAddress, err)
	}
	<-done
	log.Println("server stopped")
}
