package main

import (
	"log"

	"github.com/alphatechnolog/purplish-items/core"
	"github.com/alphatechnolog/purplish-items/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.OpenDBConnection()
	if err != nil {
		log.Fatal("A fatal error occurred", err)
		return
	}
	defer db.Close()

	r := gin.Default()
	defer r.Run()

	core.CreateItemsRoutes(db, r.Group("/items"))
}
