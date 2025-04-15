package main

import (
	"log"
	"net/http"

	"github.com/FabioRg06/Authentify/internal/bootstrap"
	"github.com/FabioRg06/Authentify/internal/config"
)

func main() {
	config.Init()
	container, err := bootstrap.NewContainer()
	if err != nil {
		log.Fatalf("❌ Failed to initialize container: %v", err)
	}
	defer container.DB.Close()

	http.HandleFunc("/register", container.UserHandler.Register)

	log.Println("🚀 Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("❌ Server failed: %v", err)
	}
}
