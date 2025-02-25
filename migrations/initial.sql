CREATE TABLE IF NOT EXISTS items (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    description VARCHAR(255),
    price INTEGER NOT NULL,
    status INTEGER CHECK (status IN (0, 1))
);

INSERT INTO
    items (id, name, price, status)
VALUES
    (
        'df4985ba-45ee-45dc-b31f-8fcbd677e9a2',
        'Pera',
        10,
        1
    ),
    (
        '0f036704-033c-46ff-af90-8e7da10daf70',
        'Agua',
        5,
        1
    );

CREATE TABLE IF NOT EXISTS items_warehouses (
    id VARCHAR(36) PRIMARY KEY,
    warehouse_id VARCHAR(36) NOT NULL,
    item_id VARCHAR(36) NOT NULL,
    quantity INTEGER NOT NULL
);

INSERT INTO
    items_warehouses (id, warehouse_id, item_id, quantity)
VALUES
    -- 10 pera -> warehouse: primario
    (
        'f67497e9-d61d-4d26-97a6-ec752a96b8c5',
        '7d090868-3df5-44e7-9280-3cad6204be59',
        'df4985ba-45ee-45dc-b31f-8fcbd677e9a2',
        10
    ),
    -- 5 agua -> warehouse: primario
    (
        'b24c5238-9867-444c-9daf-ccbe6de1bdb7',
        '7d090868-3df5-44e7-9280-3cad6204be59',
        '0f036704-033c-46ff-af90-8e7da10daf70',
        5
    );
