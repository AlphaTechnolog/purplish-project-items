package usecase

import (
	"fmt"

	"github.com/alphatechnolog/purplish-items/internal/domain"
	"github.com/alphatechnolog/purplish-items/internal/repository"
	"github.com/google/uuid"
)

type ItemUsecase struct {
	sqldbRepo repository.SQLDBRepository
}

func NewItemUsecase(sqldbRepo repository.SQLDBRepository) *ItemUsecase {
	return &ItemUsecase{
		sqldbRepo,
	}
}

func (uc *ItemUsecase) GetItems() ([]domain.Item, error) {
	query := "SELECT id, name, description, price, status FROM items;"
	items := []domain.Item{}

	rows, err := uc.sqldbRepo.Query(query)
	if err != nil {
		return items, fmt.Errorf("unable to query items: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Status)
		if err != nil {
			return items, fmt.Errorf("cannot scan queryset: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (uc *ItemUsecase) GetItemsByWarehouse(warehouseID string) ([]domain.Item, error) {
	query := `
		SELECT
			i.id,
			i.name,
			i.description,
			i.price,
			i.status
		FROM
			items_warehouses iw
			INNER JOIN items i ON i.id = iw.item_id
		WHERE
			iw.warehouse_id = ?
			AND i.status = 1;
	`

	items := []domain.Item{}
	rows, err := uc.sqldbRepo.Query(query, warehouseID)
	if err != nil {
		return items, fmt.Errorf("unable to query items by warehouse: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var item domain.Item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Status)
		if err != nil {
			return items, fmt.Errorf("cannot scan queryset: %w", err)
		}
		items = append(items, item)
	}

	return items, nil
}

func (uc *ItemUsecase) GetItem(id string) (*domain.Item, error) {
	query := "SELECT id, name, description, price, status FROM items WHERE id = ?"
	row := uc.sqldbRepo.QueryRow(query, id)

	item := &domain.Item{}
	err := row.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Status)
	if err != nil {
		return nil, fmt.Errorf("failed to scan item: %w", err)
	}

	return item, nil
}

func (uc *ItemUsecase) AssignToWarehouse(p *domain.WarehouseAssignationPayload) error {
	query := "INSERT INTO items_warehouses (item_id, warehouse_id) VALUES (?, ?)"
	_, err := uc.sqldbRepo.Execute(query, p.ItemID, p.WarehouseID)
	if err != nil {
		return fmt.Errorf("failed to assign item to warehouse: %w", err)
	}

	return nil
}

func (uc *ItemUsecase) CreateItem(item *domain.Item) error {
	query := "INSERT INTO items (id, name, description, price, status) VALUES (?, ?, ?, ?, ?)"
	item.ID = uuid.New().String()
	item.Status = true
	_, err := uc.sqldbRepo.Execute(query, item.ID, item.Name, item.Description, item.Price, item.Status)
	if err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}

	return nil
}

func (uc *ItemUsecase) UpdateItem(item *domain.Item) error {
	query := "UPDATE items SET name = ?, description = ?, price = ?, status = ? WHERE id = ?"
	_, err := uc.sqldbRepo.Execute(query, item.Name, item.Description, item.Price, item.Status, item.ID)
	if err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}

	return nil
}

func (uc *ItemUsecase) DeleteItem(id string) error {
	if _, err := uc.sqldbRepo.Execute("DELETE FROM items WHERE id = ?", id); err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}
	if _, err := uc.sqldbRepo.Execute("DELETE FROM items_warehouses WHERE item_id = ?;", id); err != nil {
		return err
	}

	return nil
}
