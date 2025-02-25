package core

import (
	"database/sql"
	"encoding/json"
	"io"
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

func getItemsByWarehouse(d *sql.DB, c *gin.Context) error {
	warehouseID := c.Param("ID")
	items, err := database.GetItemsByWarehouse(d, warehouseID)
	if err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})

	return nil
}

func createItem(d *sql.DB, c *gin.Context) error {
	bodyContents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	var createPayload database.CreateItemPayload
	if err := json.Unmarshal(bodyContents, &createPayload); err != nil {
		return err
	}

	if err = database.CreateItem(d, createPayload); err != nil {
		return err
	}

	c.JSON(http.StatusCreated, gin.H{"ok": true})

	return nil
}

func deleteItem(d *sql.DB, c *gin.Context) error {
	itemID := c.Param("ID")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return nil
	}

	if err := database.DeleteItem(d, itemID); err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

	return nil
}

func CreateItemsRoutes(d *sql.DB, r *gin.RouterGroup) {
	r.GET("/", WrapError(WithDB(d, getItems)))
	r.GET("/for-warehouse/:ID", WrapError(WithDB(d, getItemsByWarehouse)))
	r.POST("/", WrapError(WithDB(d, createItem)))
	r.DELETE("/:ID", WrapError(WithDB(d, deleteItem)))
}
