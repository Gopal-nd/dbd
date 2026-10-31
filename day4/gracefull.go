package main

import (
	"context"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

func gracefullShutdown() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	defer stop()

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(time.Duration(8 * time.Second))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, World!"))
	})

	server := http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       time.Duration(10 * time.Second),
		ReadHeaderTimeout: time.Duration(100 * time.Millisecond),
	}

	errCh := make(chan error, 1)

	go func() {
		fmt.Println("server PORT", server.Addr)
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			errCh <- err
		}
	}()

	select {
	case e := <-errCh:
		fmt.Println("error: ", e)
	case <-ctx.Done():
		fmt.Println("context done")
		// server.Shutdown(ctx)
	}
	stop() // reach here when signal is received, now it will talk to the main() function to stop

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

	defer shutdownCancel()

	server.Shutdown(shutdownCtx)
	fmt.Println("server shutdown")
}
