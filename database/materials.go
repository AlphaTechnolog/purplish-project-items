package database

import "database/sql"

type Material struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Status      bool    `json:"status"`
}

func GetMaterials(d *sql.DB) ([]Material, error) {
	var materials []Material

	rows, err := d.Query("SELECT id, name, description, status FROM materials")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var material Material
		err = rows.Scan(&material.ID, &material.Name, &material.Description, &material.Status)
		if err != nil {
			return nil, err
		}

		materials = append(materials, material)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return materials, nil
}
