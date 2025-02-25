package database

import "database/sql"

type Item struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Price       int     `json:"price"`
	Status      bool    `json:"status"`
}

func GetItems(d *sql.DB) ([]Item, error) {
	var items []Item = []Item{}

	rows, err := d.Query("SELECT id, name, description, price, status FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Status)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}

func GetItemsByWarehouse(d *sql.DB, warehouseID string) ([]Item, error) {
	var items []Item = []Item{}

	sql := `
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

	rows, err := d.Query(sql, warehouseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var item Item
		err = rows.Scan(&item.ID, &item.Name, &item.Description, &item.Price, &item.Status)
		if err != nil {
			return nil, err
		}

		items = append(items, item)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return items, nil
}
