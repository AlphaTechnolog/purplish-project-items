CREATE TABLE IF NOT EXISTS items (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(60) NOT NULL,
    description VARCHAR(255),
    status INTEGER CHECK (status IN (0, 1))
);

INSERT INTO
    items (id, name, status)
VALUES
    (
        'df4985ba-45ee-45dc-b31f-8fcbd677e9a2',
        'Manzana',
        1
    ),
    ('0f036704-033c-46ff-af90-8e7da10daf70', 'Agua', 1);
