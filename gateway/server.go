package gateway

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"rgate/config"
	"time"

	"github.com/gorilla/mux"
)

func Serve(r *mux.Router) {
	srv := &http.Server{
		Addr:         "0.0.0.0:" + config.Port(),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		log.Println("http: Starting server")
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	log.Println("http: Closing server")
	err := srv.Shutdown(ctx)
	if err != nil {
		log.Printf("error: Shutting down server: %s", err)
	}
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	os.Exit(0)
}
