
CREATE TYPE discount_type_enum AS ENUM ('percentage', 'fixed_amount');
CREATE TYPE usage_type_enum AS ENUM ('single_use', 'multi_use');

CREATE TABLE categories (
    category_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL UNIQUE
);

CREATE TABLE coupons (
    coupon_id SERIAL PRIMARY KEY,
    coupon_code VARCHAR(50) NOT NULL UNIQUE,
    discount_type discount_type_enum NOT NULL,
    discount_value DECIMAL(10, 2) NOT NULL,
    terms_and_conditions TEXT,
    usage_type usage_type_enum NOT NULL,
    valid_from TIMESTAMP NOT NULL,
    valid_until TIMESTAMP NOT NULL,
    minimum_purchase_amount DECIMAL(10, 2)
);

CREATE TABLE medicienes (
    medicine_id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category_id INT REFERENCES categories(category_id)
);

CREATE TABLE coupon_medicines (
    coupon_id INT REFERENCES coupons(coupon_id),
    medicine_id INT REFERENCES medicienes(medicine_id),
    PRIMARY KEY (coupon_id, medicine_id)
);

CREATE TABLE coupon_categories (
    coupon_id INT REFERENCES coupons(coupon_id),
    category_id INT REFERENCES categories(category_id),
    PRIMARY KEY (coupon_id, category_id)
);
