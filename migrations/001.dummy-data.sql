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

INSERT INTO
    items_warehouses (id, warehouse_id, item_id)
VALUES
    -- pera -> warehouse: primario
    (
        'f67497e9-d61d-4d26-97a6-ec752a96b8c5',
        '7d090868-3df5-44e7-9280-3cad6204be59',
        'df4985ba-45ee-45dc-b31f-8fcbd677e9a2'
    ),
    -- agua -> warehouse: primario
    (
        'b24c5238-9867-444c-9daf-ccbe6de1bdb7',
        '7d090868-3df5-44e7-9280-3cad6204be59',
        '0f036704-033c-46ff-af90-8e7da10daf70'
    );
