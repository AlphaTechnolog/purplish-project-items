package core

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/alphatechnolog/purplish-items/database"
	"github.com/alphatechnolog/purplish-items/middlewares"
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

	fmt.Println("body contents:", string(bodyContents))

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

func assignToWarehouse(d *sql.DB, c *gin.Context) error {
	bodyContents, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return err
	}

	var payload struct {
		ItemID      string `json:"item_id"`
		WarehouseID string `json:"warehouse_id"`
	}

	if err := json.Unmarshal(bodyContents, &payload); err != nil {
		return err
	}

	if err = database.AssignItemToWarehouse(d, payload.ItemID, payload.WarehouseID); err != nil {
		return err
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})

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
	r.GET("/", middlewares.APIGatewayScopeCheck("r:items"), WrapError(WithDB(d, getItems)))
	r.GET("/for-warehouse/:ID", middlewares.APIGatewayScopeCheck("r:items"), WrapError(WithDB(d, getItemsByWarehouse)))
	r.POST("/", middlewares.APIGatewayScopeCheck("c:items"), WrapError(WithDB(d, createItem)))
	r.POST("/assign-to-warehouse/", middlewares.APIGatewayScopeCheck("c:items"), WrapError(WithDB(d, assignToWarehouse)))
	r.DELETE("/:ID", middlewares.APIGatewayScopeCheck("d:items"), WrapError(WithDB(d, deleteItem)))
}
