package http

import (
	"net/http"

	"github.com/alphatechnolog/purplish-items/internal/domain"
	"github.com/alphatechnolog/purplish-items/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ItemHandler struct {
	itemUsecase *usecase.ItemUsecase
}

func NewItemHandler(itemUsecase *usecase.ItemUsecase) *ItemHandler {
	return &ItemHandler{itemUsecase}
}

func (h *ItemHandler) GetItems(c *gin.Context) {
	items, err := h.itemUsecase.GetItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *ItemHandler) GetItemsByWarehouse(c *gin.Context) {
	warehouseID := c.Param("id")
	if warehouseID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Warehouse ID is required"})
		return
	}

	if err := uuid.Validate(warehouseID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Warehouse ID"})
		return
	}

	items, err := h.itemUsecase.GetItemsByWarehouse(warehouseID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *ItemHandler) GetItem(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item ID is required"})
		return
	}

	item, err := h.itemUsecase.GetItem(itemID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

func (h *ItemHandler) AssignToWarehouse(c *gin.Context) {
	var payload domain.WarehouseAssignationPayload
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := payload.ValidateUUIDs(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if err := h.itemUsecase.AssignToWarehouse(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"ok": true})
}

func (h *ItemHandler) CreateItem(c *gin.Context) {
	var item domain.Item
	if err := c.ShouldBind(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := h.itemUsecase.CreateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"item": item})
}

func (h *ItemHandler) UpdateItem(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item ID is required"})
		return
	}

	var item domain.Item
	if err := c.ShouldBindJSON(&item); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	item.ID = itemID

	if err := h.itemUsecase.UpdateItem(&item); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"item": item})
}

func (h *ItemHandler) DeleteItem(c *gin.Context) {
	itemID := c.Param("id")
	if itemID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Item ID is required"})
		return
	}

	if err := h.itemUsecase.DeleteItem(itemID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": true})
}
