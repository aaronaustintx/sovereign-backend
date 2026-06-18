package main

import (
	"log"

	"github.com/aaronaustintx/sovereign-backend/internal/config"
	"github.com/aaronaustintx/sovereign-backend/internal/database"
	httpserver "github.com/aaronaustintx/sovereign-backend/internal/http"
)

func main() {
	cfg := config.Load()

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	router := httpserver.NewRouter(db, cfg)

	log.Println("API running on port", cfg.Port)
	router.Run(":" + cfg.Port)
}
