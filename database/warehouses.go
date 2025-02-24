package database

import (
	"database/sql"
)

type Warehouse struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Status bool   `json:"status"`
}

// Represents a warehouse item with its materials
type FullWarehouse struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Status    bool       `json:"status"`
	Materials []Material `json:"materials"`
}

func GetWarehouses(d *sql.DB) ([]Warehouse, error) {
	var warehouses []Warehouse

	rows, err := d.Query("SELECT id, name, status FROM warehouses")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var warehouse Warehouse
		err = rows.Scan(&warehouse.ID, &warehouse.Name, &warehouse.Status)
		if err != nil {
			return nil, err
		}

		warehouses = append(warehouses, warehouse)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}

func GetFullWarehouses(d *sql.DB) ([]FullWarehouse, error) {
	var warehouses []FullWarehouse

	sql := `
		SELECT
			w.id,
			w.name,
			w.status,
			m.id AS material_id,
			m.name AS material_name,
			m.description AS material_description,
			m.status AS material_status
		FROM
			warehouses w
			INNER JOIN material_warehouses mw ON mw.warehouse_id = w.id
			INNER JOIN materials m ON m.id = mw.material_id;
	`

	rows, err := d.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var warehouse FullWarehouse
		var material Material

		err = rows.Scan(
			&warehouse.ID,
			&warehouse.Name,
			&warehouse.Status,
			&material.ID,
			&material.Name,
			&material.Description,
			&material.Status,
		)

		if err != nil {
			return nil, err
		}

		found := false

		for i, w := range warehouses {
			if w.ID != warehouse.ID {
				continue
			}

			warehouses[i].Materials = append(
				warehouses[i].Materials,
				material,
			)

			found = true
		}

		if !found {
			warehouse.Materials = append(warehouse.Materials, material)
			warehouses = append(
				warehouses,
				warehouse,
			)
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return warehouses, nil
}
