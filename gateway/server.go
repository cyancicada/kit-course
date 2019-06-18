package gateway

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	kitlog "github.com/go-kit/kit/log"
)

var Config = map[string]string{
	"access_log_path" : "gateway/access.log",
	"jwt_auth_secret" : "secret",
	"pid_path" : "gateway/pid",
	"server_port" : ":9001",
}
func RunServer(logger kitlog.Logger, httpAddr string, router *Router) {
	srv := &http.Server{
		Addr:    httpAddr,
		Handler: router.r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
