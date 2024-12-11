package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	handler := setupHandler()

	server := &http.Server{
		Addr:    ":3000",
		Handler: handler,
	}

	done := make(chan os.Signal, 1)

	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("error: %s\n", err)
		}
	}()

	<-done
	log.Println("Shutting down server...")

	// close db conneection
	/*
	 code
	*/

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Error while shutting down Server. Initiating force shutdown...")
	} else {
		fmt.Print("Server exiting")
	}
}

func setupHandler() *gin.Engine {
	router := gin.New()

	return router
}
