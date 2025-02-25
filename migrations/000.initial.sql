CREATE TABLE IF NOT EXISTS items (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    description VARCHAR(255),
    price INTEGER NOT NULL,
    status INTEGER CHECK (status IN (0, 1))
);

CREATE TABLE IF NOT EXISTS items_warehouses (
    id VARCHAR(36) PRIMARY KEY,
    warehouse_id VARCHAR(36) NOT NULL,
    item_id VARCHAR(36) NOT NULL,
    CONSTRAINT unique_warehouse_item UNIQUE (warehouse_id, item_id)
);

CREATE INDEX idx_item_warehouses ON items_warehouses (warehouse_id, item_id);
