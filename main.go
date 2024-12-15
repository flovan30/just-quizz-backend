package main

import (
	"context"
	"errors"
	"fmt"
	"just-quizz-server/database"
	"just-quizz-server/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var WG sync.WaitGroup

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.InitDB()

	handler := setupHandler()

	server := &http.Server{
		Addr:    os.Getenv("API_PORT"),
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
	database.CloseDB()

	if err := server.Shutdown(context.TODO()); err != nil {
		log.Fatal("Error while shutting down Server. Initiating force shutdown...")
	} else {
		fmt.Print("Server is now shutdown.")
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

	// register routes
	routes.RegisterThemeGroup(router, &WG)

	return router
}
