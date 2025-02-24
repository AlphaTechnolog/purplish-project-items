package core

import (
	"database/sql"
	"net/http"

	"github.com/alphatechnolog/purplish-items/database"
	"github.com/gin-gonic/gin"
)

func getItems(d *sql.DB, c *gin.Context) error {
	items, err := database.GetItems(d)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})

	return nil
}

func CreateItemsRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getItems)))
}
