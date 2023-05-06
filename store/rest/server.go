package rest

import (
	"context"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var server *http.Server

func Start() {
	router := gin.Default()

	addBookRoutes(router.Group("/book"))
	addReceiptRoutes(router.Group("/book-receipt"))
	addOrderRoutes(router.Group("/order"))

	server = &http.Server{Addr: ":8080", Handler: router}

	log.Println("server listening on :8080")

	if err := server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("failed to start server: %s", err)
	}
}

func Stop() {
	if server == nil {
		return
	}

	log.Println("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("failed to shutdown server:", err)
	}

	log.Println("server exited")
	server = nil
}
