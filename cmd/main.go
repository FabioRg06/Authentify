package main

import (
	"log"
	"net/http"

	"github.com/FabioRg06/Authentify/internal/bootstrap"
	"github.com/FabioRg06/Authentify/internal/config"
	"github.com/FabioRg06/Authentify/internal/infrastructure/middleware"
)

func main() {
	config.Init()
	container, err := bootstrap.NewContainer()
	if err != nil {
		log.Fatalf("❌ Failed to initialize container: %v", err)
	}
	defer container.DB.Close()

	http.Handle("/register", middleware.LogRequests(http.HandlerFunc(container.UserHandler.Register)))
	http.Handle("/users", middleware.LogRequests(http.HandlerFunc(container.UserHandler.Get)))
	log.Println("🚀 Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
