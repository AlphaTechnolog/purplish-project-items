package main

import (
	"log"

	"github.com/alphatechnolog/purplish-items/infrastructure/database"
	"github.com/alphatechnolog/purplish-items/internal/config"
	"github.com/alphatechnolog/purplish-items/internal/di"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

const ENV_FILE = ".env"

func main() {
	cfg, err := config.LoadConfig(ENV_FILE)
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
		panic(err)
	}

	db := database.MustOpenDB("sqlite3", cfg.DatabaseURL)
	defer db.Close()

	router := gin.Default()
	defer router.Run(":" + cfg.ServerPort)

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	itemGroup := router.Group("/items")

	itemInjector := di.NewItemInjector(db)
	itemInjector.Inject(itemGroup)
}
