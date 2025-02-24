CREATE TABLE IF NOT EXISTS materials (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    description TEXT,
    status INTEGER CHECK(status IN (0, 1))
);

CREATE TABLE IF NOT EXISTS warehouses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    status INTEGER CHECK(status IN (0, 1))
);

CREATE TABLE IF NOT EXISTS material_warehouses (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    material_id INTEGER,
    warehouse_id INTEGER,
    quantity INTEGER,
    FOREIGN KEY (material_id) REFERENCES materials(id) ON DELETE CASCADE,
    FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE,
    UNIQUE (material_id, warehouse_id)
);

INSERT INTO materials (name, status) VALUES
('Manzana', 1),
('Agua', 1);

INSERT INTO warehouses (name, status) VALUES ('Principal', 1);

INSERT INTO material_warehouses (material_id, warehouse_id, quantity) VALUES
(1, 1, 3),
(2, 1, 10);