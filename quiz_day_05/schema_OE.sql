
-- TASK 1
CREATE TABLE oe.categories (
    category_id SMALLINT PRIMARY KEY,
    category_name VARCHAR(15),
    description TEXT,
    picture BYTEA
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);

CREATE TABLE oe.suppliers (
    supplier_id SMALLINT PRIMARY KEY,
    company_name VARCHAR(40),
    contact_name VARCHAR(30),
    contact_title VARCHAR(30),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    phone VARCHAR(24),
    fax VARCHAR(24),
    home_page TEXT
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);

CREATE TABLE oe.employees(
    employee_id SMALLINT PRIMARY KEY,
    last_name VARCHAR(20),
    first_name VARCHAR(10),
    title VARCHAR(30),
    title_of_courtesy VARCHAR(25),
    birth_date DATE,
    hire_date DATE,
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    home_phone VARCHAR(24),
    extension VARCHAR(4),
    photo BYTEA,
    notes TEXT,
    reports_to SMALLINT,
    photo_path VARCHAR(255)
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);

CREATE TABLE oe.customers(
    customer_id VARCHAR(5) PRIMARY KEY,
    company_name VARCHAR(40),
    contact_name VARCHAR(30),
    contact_title VARCHAR(30),
    address VARCHAR(60),
    city VARCHAR(15),
    region VARCHAR(15),
    postal_code VARCHAR(10),
    country VARCHAR(15),
    phone VARCHAR(24),
    fax VARCHAR(24)
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);


CREATE TABLE oe.shippers(
    shipper_id SMALLINT PRIMARY KEY,
    company_name VARCHAR(40),
    phone VARCHAR(24)
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);   

CREATE TABLE oe.products(
    product_id SMALLINT PRIMARY KEY,
    product_name VARCHAR(40),
    quantity_per_unit VARCHAR(20),
    unit_price REAL,
    units_in_stock SMALLINT,
    units_on_order SMALLINT,
    reorder_level SMALLINT,
    discontinued INT,
    CONSTRAINT fk_products_supplier 
        FOREIGN KEY (supplier_id) REFERENCES oe.suppliers(supplier_id),
    CONSTRAINT fk_products_category 
        FOREIGN KEY (category_id) REFERENCES oe.categories(category_id)
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);

CREATE TABLE oe.orders(
    order_id SMALLINT PRIMARY KEY,
    order_date DATE,
    required_date DATE,
    shipped_date DATE,
    freight REAL,
    ship_name VARCHAR(40),
    ship_address VARCHAR(60),
    ship_city VARCHAR(15),
    ship_region VARCHAR(15),
    ship_postal_code VARCHAR(10),
    ship_country VARCHAR(15),
    CONSTRAINT fk_orders_customer 
        FOREIGN KEY (customer_id) REFERENCES oe.customers(customer_id),
    CONSTRAINT fk_orders_employee 
        FOREIGN KEY (employee_id) REFERENCES oe.employees(employee_id),
    CONSTRAINT fk_orders_shipper 
        FOREIGN KEY (ship_via) REFERENCES oe.shippers(shipper_id)
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);

CREATE TABLE oe.order_details(
    CONSTRAINT fk_order_detail_order 
        FOREIGN KEY (order_id) REFERENCES oe.orders(order_id) ON DELETE CASCADE,
    CONSTRAINT fk_order_detail_product 
        FOREIGN KEY (product_id) REFERENCES oe.products(product_id),
    unit_price REAL,
    quantity SMALLINT,
    discount REAL,
    PRIMARY KEY (order_id, product_id)
    -- created_at TIMESTAMP,
    -- updated_at TIMESTAMP
);



-- TASK 2

-- nomor 1
SELECT 
    c.category_id,
    c.category_name,
    COUNT(p.product_id) AS total_products
FROM oe.categories AS c
LEFT JOIN oe.products AS p 
    ON c.category_id = p.category_id
GROUP BY 
    c.category_id,
    c.category_name
ORDER BY total_products DESC;

-- nomor 2
SELECT 
    s.supplier_id,
    s.company_name,
    COUNT(p.product_id) AS total_products
FROM oe.suppliers AS s
LEFT JOIN oe.products AS p 
    ON s.supplier_id = p.supplier_id
GROUP BY 
    s.supplier_id,
    s.company_name
ORDER BY total_products DESC;

-- nomor 3
SELECT 
    s.supplier_id,
    s.company_name,
    COUNT(p.product_id) AS total_products,
    TO_CHAR(AVG(p.unit_price), 'FM999,999,990.00') AS avg_unit_price
FROM oe.suppliers AS s
LEFT JOIN oe.products AS p 
    ON s.supplier_id = p.supplier_id
GROUP BY 
    s.supplier_id,
    s.company_name
ORDER BY avg_unit_price DESC NULLS LAST;

-- nomor 4
SELECT 
    p.product_id,
    p.product_name,
    s.supplier_id,
    s.company_name,
    p.unit_price,
    p.units_in_stock,
    p.units_on_order,
    p.reorder_level
FROM oe.products AS p
INNER JOIN oe.suppliers AS s 
    ON p.supplier_id = s.supplier_id
WHERE p.units_in_stock <= p.reorder_level
    AND p.discontinued = 0
ORDER BY p.product_name;

-- nomor 5
SELECT 
    c.customer_id,
    c.company_name,
    COUNT(o.order_id) AS total_orders
FROM oe.customers AS c
LEFT JOIN oe.orders AS o 
    ON c.customer_id = o.customer_id
GROUP BY 
    c.customer_id,
    c.company_name
ORDER BY total_orders DESC;

-- nomor 6
SELECT 
    o.order_id,
    o.customer_id,
    o.order_date,
    o.required_date,
    o.shipped_date,
    (o.shipped_date - o.order_date) AS delivery_time
FROM oe.orders AS o
WHERE o.shipped_date IS NOT NULL
    AND (o.shipped_date - o.required_date) > 7
ORDER BY delivery_time DESC;

-- nomor 7
SELECT 
    p.product_id,
    p.product_name,
    SUM(od.quantity) AS total_qty
FROM oe.products AS p
INNER JOIN oe.order_details AS od 
    ON p.product_id = od.product_id
GROUP BY 
    p.product_id,
    p.product_name
ORDER BY total_qty DESC;

-- nomor 8
SELECT 
    c.category_id,
    c.category_name,
    SUM(od.quantity) AS total_qty_ordered
FROM oe.categories AS c
INNER JOIN oe.products AS p 
    ON c.category_id = p.category_id
INNER JOIN oe.order_details AS od 
    ON p.product_id = od.product_id
GROUP BY 
    c.category_id,
    c.category_name
ORDER BY total_qty_ordered DESC;

-- nomor 9
WITH category_orders AS (
    SELECT 
        c.category_id,
        c.category_name,
        SUM(od.quantity) AS total_qty_ordered
    FROM oe.categories AS c
    INNER JOIN oe.products AS p 
        ON c.category_id = p.category_id
    INNER JOIN oe.order_details AS od 
        ON p.product_id = od.product_id
    GROUP BY 
        c.category_id,
        c.category_name
)
SELECT 
    category_id,
    category_name,
    total_qty_ordered
FROM category_orders
WHERE total_qty_ordered = (SELECT MIN(total_qty_ordered) FROM category_orders)
    OR total_qty_ordered = (SELECT MAX(total_qty_ordered) FROM category_orders)
ORDER BY total_qty_ordered DESC;

-- nomor 10
SELECT 
    s.shipper_id,
    s.company_name,
    p.product_id,
    p.product_name,
    SUM(od.quantity) AS total_qty_ordered
FROM oe.shippers AS s
INNER JOIN oe.orders AS o 
    ON s.shipper_id = o.ship_via
INNER JOIN oe.order_details AS od 
    ON o.order_id = od.order_id
INNER JOIN oe.products AS p 
    ON od.product_id = p.product_id
WHERE o.shipped_date IS NOT NULL
GROUP BY 
    s.shipper_id,
    s.company_name,
    p.product_id,
    p.product_name
ORDER BY s.shipper_id, p.product_name;

-- nomor 11
WITH shipper_product_stats AS (
    SELECT 
        s.shipper_id,
        s.company_name,
        p.product_id,
        p.product_name,
        SUM(od.quantity) AS total_qty_ordered
    FROM oe.shippers AS s
    INNER JOIN oe.orders AS o 
        ON s.shipper_id = o.ship_via
    INNER JOIN oe.order_details AS od 
        ON o.order_id = od.order_id
    INNER JOIN oe.products AS p 
        ON od.product_id = p.product_id
    WHERE o.shipped_date IS NOT NULL
    GROUP BY 
        s.shipper_id,
        s.company_name,
        p.product_id,
        p.product_name
),
ranked_products AS (
    SELECT 
        shipper_id,
        company_name,
        product_id,
        product_name,
        total_qty_ordered,
        ROW_NUMBER() OVER (
            PARTITION BY shipper_id 
            ORDER BY total_qty_ordered DESC
        ) AS rank_most,
        ROW_NUMBER() OVER (
            PARTITION BY shipper_id 
            ORDER BY total_qty_ordered ASC
        ) AS rank_least
    FROM shipper_product_stats
)
SELECT 
    shipper_id,
    company_name,
    product_id,
    product_name,
    total_qty_ordered
FROM ranked_products
WHERE rank_most = 1 OR rank_least = 1
ORDER BY shipper_id, product_name;
