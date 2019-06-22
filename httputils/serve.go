package httputils

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Serve ...
func Serve(h http.Handler) {

	srv := http.Server{
		Addr:              ":8080",
		Handler:           h,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	go func() {
		srv := http.Server{
			Addr:              ":9200",
			Handler:           promhttp.Handler(),
			ReadHeaderTimeout: 5 * time.Second,
		}
		srv.ListenAndServe()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, os.Interrupt)
	sig := <-quit // wait Signal
	fmt.Printf("SIGNAL %+v recieved, then sutting down...\n", sig)

	// Shutdown server
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		fmt.Fprintln(os.Stderr, "failed to gracefully shutdown:", err)
	} else {
		fmt.Println("succeed to gracefully shutdown")
	}
}
