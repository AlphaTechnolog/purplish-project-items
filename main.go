package main

import (
	"log"

	"github.com/alphatechnolog/purplish-materials/core"
	"github.com/alphatechnolog/purplish-materials/database"
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

	core.CreateMaterialsRoutes(db, r.Group("/materials"))
	core.CreateWarehousesRoutes(db, r.Group("/warehouses"))
}
