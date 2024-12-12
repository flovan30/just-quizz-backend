package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
)

var WG sync.WaitGroup

func main() {
	// database.InitDB()

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

	WG.Wait()

	// close db connection
	// database.CloseDB()

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatal("Error while shutting down Server. Initiating force shutdown...")
	} else {
		fmt.Print("Server exiting")
	}
}

func setupHandler() *gin.Engine {
	router := gin.New()

	// default route
	router.GET("/ping", func(c *gin.Context) {
		WG.Add(1)
		defer WG.Done()

		c.JSON(http.StatusOK, gin.H{"message": "Pong !"})
	})

	return router
}
