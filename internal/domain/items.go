package domain

import "github.com/google/uuid"

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	Price       int     `json:"price"`
	Status      bool    `json:"status"`
}

type WarehouseAssignationPayload struct {
	WarehouseID string `json:"warehouse_id"`
	ItemID      string `json:"item_id"`
}

func (wp *WarehouseAssignationPayload) ValidateUUIDs() error {
	if err := uuid.Validate(wp.WarehouseID); err != nil {
		return err
	}
	if err := uuid.Validate(wp.ItemID); err != nil {
		return err
	}
	return nil
}
